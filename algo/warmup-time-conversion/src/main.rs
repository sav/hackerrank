use std::io::{self, BufRead};

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

fn time_conversion(s: &str) {
    let mut s: String = String::from(s);
    let am: bool = s.ends_with("AM");

    s = s.replace("AM", "");
    s = s.replace("PM", "");

    let mut time_iterator = s
        .trim_end()
        .split(':')
        .map(|s| s.to_string().parse::<u8>().unwrap());

    let mut h: u8 = time_iterator.next().unwrap();
    let m: u8 = time_iterator.next().unwrap();
    let s: u8 = time_iterator.next().unwrap();

    assert_ne!(h, 0);

    if am {
        if h == 12 {
            h = 0;
        }
    } else {
        if h < 12 {
            h += 12;
        }
    }

    println!("{:0>2}:{:0>2}:{:0>2}", h, m, s);
}

fn main() {
    let stdin = io::stdin();

    for line in stdin.lock().lines() {
        time_conversion(&line.unwrap());
    }
}
