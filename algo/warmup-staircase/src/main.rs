use std::io::{self, BufRead};

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

fn staircase(n: u32) {
    for i in 0..n as u32 {
        let mut s: String = String::from("");
        for j in 0..n as u32 {
            if i >= n - j - 1 {
                s.push('#');
            } else {
                s.push(' ');
            }
        }
        println!("{}", s);
    }
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let n = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim()
        .parse::<u32>()
        .unwrap();

    staircase(n);
}
