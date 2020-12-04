use std::fs;
use std::path::Path;
use std::collections::HashMap;
use regex::Regex;

fn solution_1(passports: &[HashMap<String, String>]) -> Result<usize, &'static str> {
    let mut valid = 0;
    for passport in passports.iter() {
        if has_required_fields(passport) {
            valid += 1;
        }
    }
    Ok(valid)
}

fn solution_2(passports: &[HashMap<String, String>]) -> Result<usize, &'static str> {
    let mut valid = 0;
    'outer: for passport in passports.iter() {
        if !has_required_fields(passport) {
           continue; 
        }

        for (k, v) in passport.iter() {
            if let Err(_) = is_valid_field(k, v) {
                continue 'outer;
            }
        }

        valid += 1;
    }
    Ok(valid)
}

fn has_required_fields(passport: &HashMap<String, String>) -> bool {
    for &k in ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"].iter() {
        if !passport.contains_key(k) {
            return false
        }
    }
    true
}

fn is_valid_field<'a>(key: &str, value: &str) -> Result<(), &'a str> {
    let value_chars: Vec<char> = value.chars().collect();
    
    match key {
        "byr" | "iyr" | "eyr" => {
            let year = match value.parse::<usize>() {
                Ok(y) => y,
                _ => return Err("unable to parse year")
            };
            if key == "byr" && (year < 1920 || year > 2002) {
                return Err("byr not within range")
            } 
            if key == "iyr" && (year < 2010 || year > 2020) {
                return Err("iyr not within range")
            }
            if key == "eyr" && (year < 2020 || year > 2030) {
                return Err("eyr not within range")
            }
        },
        "hgt" => {
            let unit_start_index = value_chars.len() as i32 - 2;
            if unit_start_index - 1 < 0 {
                return Err("hgt value too short")
            }

            let magnitude_chars = &value_chars[..unit_start_index as usize];
            let magnitude_string: String = magnitude_chars.iter().collect();
            let magnitude = match magnitude_string.parse::<usize>() {
                Ok(m) => m,
                _ => return Err("unable to parse hgt magnitude")
            };

            let unit_chars = &value_chars[unit_start_index as usize..];
            let unit: String = unit_chars.iter().collect();
            match unit.as_str() {
                "cm" => {
                    if magnitude < 150 || magnitude > 193 {
                        return Err("hgt not in range")
                    }
                },
                "in" => {
                    if magnitude < 59 || magnitude > 76 {
                        return Err("hgt not in range")
                    }
                }
                _ => return Err("invalid hgt unit")
            }
        },
        "hcl" => {
            let mut value_chars_iter = value_chars.iter();
            if value_chars_iter.next().ok_or("hcl value too short")? != &"#".chars().next().unwrap() {
                return Err("invalid hcl prefix")
            }

            let information: String = value_chars_iter.collect();
            if information.chars().count() != 6 {
                return Err("incorrect hcl length")
            }

            if !Regex::new(r"[0-9a-f]{6}").unwrap().is_match(information.as_str()) {
                return Err("at least 1 illegal hcl character detected")
            }
        },
        "ecl" => {
            if value_chars.len() != 3 {
                return Err("invalid ecl length")
            }

            if !Regex::new(r"amb|blu|brn|gry|grn|hzl|oth").unwrap().is_match(value) {
                return Err("illegal ecl value")
            }
        },
        "pid" => {
            if value_chars.len() != 9 {
                return Err("invalid pid length")
            }

            if !Regex::new(r"[0-9]{9}").unwrap().is_match(value) {
                return Err("illegal pid value")
            }
        },
        "cid" => {
            // noop
        },
        _ => return Err("unknown field")
    }

    Ok(())
}

fn parse_passport<'a>(data: String) -> Result<HashMap<String, String>, &'a str> {
    let mut passport: HashMap<String, String> = HashMap::new();
    for field in data.split(" ") {
        let mut key_value = field.split(":");
        passport.insert(key_value.next().ok_or("couldn't get field key")?.to_string(), key_value.next().ok_or("couldn't get field value")?.to_string());
    }
    Ok(passport)
}

fn parse_input<'a, P: AsRef<Path>>(filename: P) -> Result<Vec<HashMap<String, String>>, &'a str> {
    let mut passports: Vec<HashMap<String, String>> = vec![];
    let input = fs::read_to_string(filename).unwrap();
    let raw_passports_iter = input.split("\n\n");
    for raw_passport in raw_passports_iter {
        passports.push(parse_passport(raw_passport.replace("\n", " "))?);
    }
    Ok(passports)
}

fn main() {
    let passports = parse_input("input.txt").unwrap();
    println!("the solution to part 1 is: {}", solution_1(&passports).unwrap());
    println!("the solution to part 2 is: {}", solution_2(&passports).unwrap());
}