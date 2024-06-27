use std::fs::File;
use std::io::{self, BufRead, BufReader};

fn main() -> io::Result<()> {
    let file = File::open("../../../data/2019/day1.txt")?;
    let reader = BufReader::new(file);

    let mut first = 0;
    let mut second = 0;

    for line in reader.lines() {
        let num: i32 = line?.trim().parse().expect("Expected a number");
        let fuel = (num as f32 / 3.0).floor() as i32 - 2;

        first += fuel;
        second += fuel_of_fuel(fuel);
    }

    println!("First: {}", first);
    println!("Second: {}", second);

    Ok(())
}

fn fuel_of_fuel(fuel: i32) -> i32 {
    if fuel <= 0 {
        return 0;
    }
    return fuel + fuel_of_fuel((fuel as f32 / 3.0).floor() as i32 - 2);
}
