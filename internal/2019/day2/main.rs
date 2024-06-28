use std::fs::File;
use std::io::{self, BufRead, BufReader};

fn main() -> io::Result<()> {
    let file = File::open("../../../data/2019/day2.txt")?;
    let reader = BufReader::new(file);

    let mut list: Vec<i32> = vec![];
    for line_result in reader.lines() {
        let line = line_result?;
        list = line.split(',')
            .map(|s| s.trim().parse::<i32>().unwrap())
            .collect(); 
    }

    let mut res = run_intcode(12, 2, list.clone());
    println!("First: {}", res);

    let mut noun = 0;
    let mut verb = 0;
    while res != 19690720 {
        // trying until 100 to start
        if verb == 100 {
            noun += 1;
            verb = 0;
        } else {
            verb += 1;
        }
        res = run_intcode(noun, verb, list.clone());
    }

    println!("Second: {}", noun * 100 + verb); 
    Ok(())
}

fn run_intcode(noun: i32, verb: i32, list: Vec<i32>) -> i32 {
    let mut ops: Vec<i32> = list.clone();
    ops[1] = noun;
    ops[2] = verb;

    let mut i = 0;
    while i < ops.len()-4 { 
        let a = ops[i+1] as usize;
        let b = ops[i+2] as usize;
        let c = ops[i+3] as usize;

        if ops[i] == 1 {
            ops[c] = ops[a] + ops[b];
        } else if ops[i] == 2 {
            ops[c] = ops[a] * ops[b];
        } else if ops[i] == 99 {
            break;
        } else {
            panic!("Number is not a valid operation {}", ops[i])
        }

        i += 4;
    }
    return ops[0];
}
