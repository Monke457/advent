use std::fs::File;
use std::io::{self, BufRead, BufReader};

fn main() -> io::Result<()> {
    let Ok((first, second)) = read_lines("../../../data/2019/day3.txt") 
        else { 
            panic!("Error") 
        };

    for element in first {
        println!("{}", element);
    }
    for element in second {
        println!("{}", element);
    }

    Ok(())
}

fn read_lines(filename: &str) -> io::Result<(Vec<String>, Vec<String>)> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);

    let mut lines = vec![];
    for line_result in reader.lines() {
        let line = line_result?;
        lines.push(line);
    }

    if lines.len() != 2 {
        panic!("Expected 2 lines, have {}", lines.len());
    }

    return Ok((parse_line(&lines[0]), parse_line(&lines[1])));
}

fn parse_line(line: &String) -> Vec<String> {
    let mut els: Vec<String> = vec![];
    for element in line.split(',') {
        els.push(element.to_string());
    }
    return els.clone();
}

