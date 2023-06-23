use std::io::{self, BufRead};

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

fn mini_max_sum(arr: &[u64]) {
    let len: usize = arr.len();
    let mut sum: u64 = arr[0];
    let mut min: u64 = arr[0];
    let mut max: u64 = arr[0];

    for i in 1..len as usize {
        sum += arr[i];
        if arr[i] < min {
            min = arr[i];
        }
        if arr[i] > max {
            max = arr[i];
        }
    }
    println!("{} {}", sum - max, sum - min);
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let arr: Vec<u64> = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim_end()
        .split(' ')
        .map(|s| s.to_string().parse::<u64>().unwrap())
        .collect();

    mini_max_sum(&arr);
}
