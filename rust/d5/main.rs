use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn solution_1(input: &[String]) -> usize {
    let mut highest = 0;
    for bp in input.iter() {
        let id = compute_id(bp);
        if id > highest {
            highest = id;
        }
    }

    highest
}

fn solution_2(input: &[String]) -> usize {
    let mut ids: Vec<usize> = vec![];
    for bp in input.iter() {
        ids.push(compute_id(bp));
    }
    ids.sort();

    let mut my_id = 0;
    for (i, id) in ids.iter().enumerate() {
        if i + 2 > ids.len() {
            break;
        }

        if ids.get(i+1).unwrap() - id == 2 {
            my_id = id + 1;
            break;
        }
    }

    my_id
}

fn compute_id(bp: &str) -> usize {
    let bp_chars: Vec<char> = bp.chars().collect();
    let row = locate(&bp_chars[..7], 128, 'F');
    let col = locate(&bp_chars[7..], 8, 'L');

    row*8 + col
}

fn locate(input: &[char], mut space: usize, lower: char) -> usize {
    let mut location = 0;
    for &c in input.iter() {
        if c == lower {
			space = location + (space-location)/2
		} else {
			location = location + (space-location)/2
		}
    }

    location
}

fn read_lines<P: AsRef<Path>>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>> {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn main() {
    let input: Vec<String> = read_lines("input.txt").unwrap().map(|line| match line {
        Ok(text) => text,
        _ => panic!("error parsing input file")
    }).collect();

    println!("the answer to part 1 is: {}", solution_1(&input));
    println!("the answer to part 2 is: {}", solution_2(&input));
}