use std::collections::HashMap;

fn solution_1(input: &[i32], sum: i32) -> Option<i32> {
    let mut compliments: HashMap<i32, i32> = HashMap::new();
    for val in input.iter() {
        let compliment = sum - val;
        match compliments.get(&compliment) {
            Some(_) => return Some(compliment * val),
            _ => (),
        }
        compliments.insert(*val, *val);
    }
    None
}

fn solution_2(input: &[i32], sum: i32) -> Option<i32> {
    for (i, vali) in input.iter().enumerate() {
        let mut compliments: HashMap<i32, i32> = HashMap::new();
        for (j, valj) in input.iter().enumerate() {
            if i == j {
                continue;
            }
            let compliment = sum - vali - valj;
            match compliments.get(&compliment) {
                Some(_) => return Some(compliment * vali * valj),
                _ => (),
            }
            compliments.insert(*valj, *valj);
        }
    }

    None
}

fn main() {
    let input = [1, 2, 3, 4, 5, 6];
    solution_1(&input, 2020);
    solution_2(&input, 2020);
}