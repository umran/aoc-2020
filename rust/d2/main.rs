use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn parse_line(input: &str) -> Result<(usize, usize, char, Vec<char>), &str> {
    let rule_rest: Vec<&str> = input.split(":").collect();
    let space_pw: Vec<&str> = rule_rest.get(1).ok_or("invalid index")?.split(" ").collect();
    let pw: Vec<char> = space_pw.get(1).ok_or("invalid index")?.chars().collect();
    let rnge_char: Vec<&str> = rule_rest.get(0).ok_or("invalid index")?.split(" ").collect();
    let rnge: Vec<&str> = rnge_char.get(0).ok_or("invalid index")?.split("-").collect();
    let character_chars: Vec<char> = rnge_char.get(1).ok_or("invalid index")?.chars().collect();

    let result = (
        rnge.get(0).ok_or("invalid index")?.parse::<usize>().unwrap(),
        rnge.get(1).ok_or("invalid index")?.parse::<usize>().unwrap(),
        *character_chars.get(0).ok_or("invalid index")?,
        pw
    );
    Ok(result)
}

fn solution_1(input: &[String]) -> Result<i32, &str> {
    let mut valid = 0;
    for line in input.iter() {
        let (min, max, character, password) = parse_line(line)?;
        let mut count = 0;
        for &c in password.iter() {
            if character == c {
                count += 1;
            }
        }

        if count >= min && count <= max {
            valid += 1;
        }
    }
    Ok(valid)
}

fn solution_2(input: &[String]) -> Result<i32, &str> {
    let mut valid = 0;
    for line in input.iter() {
        let (min, max, character, password) = parse_line(line)?;
        let &first_char = password.get(min-1).ok_or("invalid index")?;
        let &second_char = password.get(max-1).ok_or("invalid index")?;

        if first_char == character && second_char == character {
            continue;
        }

        if first_char == character || second_char == character {
            valid += 1;
        }
    }
    Ok(valid)
}

fn read_lines<P: AsRef<Path>>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>> {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn main() {
    let input: Vec<String> = read_lines("./input.txt").unwrap().map(|line| match line {
        Ok(text) => text.parse::<String>().unwrap(),
        _ => panic!("failed to read line")
    }).collect();

    println!("the solution to part 1 is {}", solution_1(&input).unwrap());
    println!("the solution to part 2 is {}", solution_2(&input).unwrap());
}