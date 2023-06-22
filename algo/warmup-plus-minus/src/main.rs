use std::io::{self, BufRead};

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

fn plus_minus(arr: &[i32]) {
    let len: usize = arr.len();
    let mut count: [f32; 3] = [0.0, 0.0, 0.0];

    for i in 0..len as usize {
        if arr[i] > 0 {
            count[0] += 1.0;
        } else if arr[i] < 0 {
            count[1] += 1.0;
        } else {
            count[2] += 1.0;
        }
    }

    let flen: f32 = len as f32;
    let proportions: [f32; 3] = [count[0] / flen, count[1] / flen, count[2] / flen];

    println!("{:.6}", proportions[0]);
    println!("{:.6}", proportions[1]);
    println!("{:.6}", proportions[2]);
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let _ = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim()
        .parse::<i32>()
        .unwrap();

    let arr: Vec<i32> = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim_end()
        .split(' ')
        .map(|s| s.to_string().parse::<i32>().unwrap())
        .collect();

    plus_minus(&arr);
}
