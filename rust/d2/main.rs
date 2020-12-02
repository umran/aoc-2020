use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn solution_1(_: &[String]) -> Result<i32, &str> {
    Err("blah")
}

fn solution_2(_: &[String]) -> Result<i32, &str> {
    Err("blah")
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