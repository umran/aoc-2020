use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::collections::HashMap;

fn solution_1(directory: &HashMap<String, HashMap<String, usize>>, target: &str) -> usize {
    let mut total = 0;
    for (bag, _) in directory.iter() {
        if leads_to_target(directory, bag, target) {
            total += 1;
        }
    }

    total
}

fn solution_2(directory: &HashMap<String, HashMap<String, usize>>, bag: &str) -> usize {
    return count_inner_bags(directory, bag)
}

fn leads_to_target(directory: &HashMap<String, HashMap<String, usize>>, bag: &str, target: &str) -> bool {
    for (candidate, _) in directory.get(bag).unwrap().iter() {
        if candidate == target {
            return true
        }
        if leads_to_target(directory, candidate, target) {
            return true
        }
    }

    false
}

fn count_inner_bags(directory: &HashMap<String, HashMap<String, usize>>, bag: &str) -> usize {
    let mut total = 0;
    for (inner_bag, count) in directory.get(bag).unwrap().iter() {
        total += count + count*count_inner_bags(directory, inner_bag);
    }
    total
}

fn parse_bags(input: &[String]) -> HashMap<String, HashMap<String, usize>> {
    let mut directory: HashMap<String, HashMap<String, usize>> = HashMap::new();
    for line_with_trailing_period in input.iter() {
        let mut inner_map: HashMap<String, usize> = HashMap::new();
        let line = line_with_trailing_period.strip_suffix(".").unwrap();
        let mut key_contents = line.split("contain");
        let key = key_contents.next().unwrap().strip_suffix(" ").unwrap();
        let contents = key_contents.next().unwrap().split(",");
        for content_with_space in contents {
            let content = content_with_space.strip_prefix(" ").unwrap();
            if content == "no other bags" {
                break;
            }
            let mut content_parts = content.split(" ");
            let amount = content_parts.next().unwrap().parse::<usize>().unwrap();
            let rest: Vec<&str> = content_parts.collect();
            let mut inner_key = rest.join(" ");
            if amount == 1 {
                inner_key = [inner_key, "s".to_string()].join("");
            }
            inner_map.insert(inner_key, amount);
        }
        directory.insert(key.to_string(), inner_map);
    }

    directory
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

    let directory = parse_bags(&input);


    println!("the solution to part 1 is: {}", solution_1(&directory, "shiny gold bags"));
    println!("the solution to part 2 is: {}", solution_2(&directory, "shiny gold bags"));
}