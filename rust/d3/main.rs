use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn solution_1(input: &[String], slope: (usize, usize)) -> Result<usize, &str> {
    count_trees(input, slope)
}

fn solution_2<'a>(input: &'a [String], slopes: &'a [(usize, usize)]) -> Result<usize, &'a str> {
    let mut product = 1;
    for &slope in slopes.iter() {
        product = product * count_trees(input, slope)?;
    }

    Ok(product)
}

fn count_trees(input: &[String], slope: (usize, usize)) -> Result<usize, &str> {
    let mut trees: usize = 0;
    let (mut x, mut y) = (0, 0);
    let (dx, dy) = slope;
    
    while y < input.len() {
        let row: Vec<char> = input.get(y)
            .ok_or("invalid y index")?
            .chars()
            .collect();

        if *row.get(x).ok_or("invalid x index")? == "#".chars().next().ok_or("error converting string literal to char")? {
            trees += 1;
        }

        x = (x + dx) % row.len();
        y += dy;
    }
    
    Ok(trees)
}

fn read_lines<P: AsRef<Path>>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>> {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn main() {
    let input: Vec<String> = read_lines("./input.txt").unwrap().map(|line| match line {
        Ok(text) => text,
        _ => panic!("couldn't read line")
    }).collect();
    println!("the solution to part 1 is {}", solution_1(&input, (3, 1)).unwrap());
    println!("the solution to part 2 is {}", solution_2(&input, &[
        (1, 1),
        (3, 1),
        (5, 1),
        (7, 1),
        (1, 2),
    ]).unwrap());
}