use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::collections::HashSet;

fn solution_1(input: &[i32], sum: i32) -> Result<i32, &str> {
    let mut compliments: HashSet<i32> = HashSet::new();
    for &val in input.iter() {
        let compliment = sum - val;
        if compliments.contains(&compliment) {
            return Ok(compliment * val);
        }
        compliments.insert(val);
    }

    Err("couldn't find a solution")
}

fn solution_2(input: &[i32], sum: i32) -> Result<i32, &str> {
    for (i, &vali) in input.iter().enumerate() {
        let mut compliments: HashSet<i32> = HashSet::new();
        for (j, &valj) in input.iter().enumerate() {
            if i == j {
                continue;
            }
            let compliment = sum - vali - valj;
            if compliments.contains(&compliment) {
                return Ok(compliment * vali * valj);
            }
            compliments.insert(valj);
        }
    }
    
    Err("couldn't find a solution")
}

fn read_lines<P: AsRef<Path>>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>> {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn main() {
    let input: Vec<i32> = read_lines("./input.txt").unwrap().map(|line| match line {
        Ok(text) => text.parse::<i32>().unwrap(),
        _ => panic!("failed to read line")
    }).collect();

    println!("the solution to part 1 is {}", solution_1(&input, 2020).unwrap());
    println!("the solution to part 2 is {}", solution_2(&input, 2020).unwrap());
}