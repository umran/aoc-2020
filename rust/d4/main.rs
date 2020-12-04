use std::fs;
use std::path::Path;
use std::collections::HashMap;

fn solution_1(passports: &[HashMap<String, String>]) -> Result<usize, &'static str> {
    let mut valid = 0;
    for passport in passports.iter() {
        if has_required_fields(passport) {
            valid += 1;
        }
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

// fn is_valid_field(key: &str, value: &str) -> bool {
//     let _: Vec<char> = value.chars().collect();
    
//     match key {
//         "byr" | "iyr" | "eyr" => {

//         },
//         "hgt" => {

//         },
//         "hcl" => {

//         },
//         "ecl" => {

//         },
//         "pid" => {

//         },
//         "cid" => {

//         },
//         _ => return false
//     }

//     true
// }

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
}