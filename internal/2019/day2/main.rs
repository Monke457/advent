use std::fs::File;
use std::io::{self, BufRead, BufReader};

fn main() -> io::Result<()> {
    let file = File::open("../../../data/2019/day2.txt")?;
    let reader = BufReader::new(file);

    for line_result in reader.lines() {
        let line = line_result?;
        let mut ops: Vec<i32> = line
            .split(',')
            .map(|s| s.trim().parse::<i32>().unwrap())
            .collect(); 

        ops[1] = 12;
        ops[2] = 2;
        let mut i = 0;
        while i < ops.len()-4 && ops[i] != 99 { 
            let a = ops[i] as usize;
            let b = ops[i+1] as usize;
            let c = ops[i+2] as usize;

            if ops[i] == 1 {
                ops[c] = ops[a] + ops[b];
            }

            if ops[i] == 2 {
                ops[c] = ops[a] * ops[b];
            }
            i += 4;
        }
        for op in ops {
            println!("{}", op);
        }
    }

    Ok(())
}
