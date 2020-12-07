use std::fs;
use std::collections::HashSet;
use std::collections::HashMap;

fn solution_1(groups: &[&str]) -> usize {
    let mut total = 0;
    for group in groups.iter() {
        let mut seen_questions: HashSet<char> = HashSet::new();
        let individuals: Vec<&str> = group.split("\n").collect();
        for &individual in individuals.iter() {
            for q in individual.chars() {
                if seen_questions.contains(&q) {
                    continue;
                }
                seen_questions.insert(q);
                total += 1;
            }
        }
    }
    total
}

fn solution_2(groups: &[&str]) -> usize {
    let mut total = 0;
    for group in groups {
        let mut question_counts: HashMap<char, usize> = HashMap::new();
        let individuals: Vec<&str> = group.split("\n").collect();
        for &individual in individuals.iter() {
            for q in individual.chars() {
                if let Some(count) = question_counts.get_mut(&q) {
                    *count += 1;
                    if count == &individuals.len() {
                        total += 1;
                    }
                    continue;
                }
                question_counts.insert(q, 1);
                if individuals.len() == 1 {
                    total += 1;
                }
            }
        }
    }
    total
}

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    let groups: Vec<&str> = input.split("\n\n").collect();
    
    println!("the solution to part 1 is {}", solution_1(&groups));
    println!("the solution to part 2 is {}", solution_2(&groups));
}