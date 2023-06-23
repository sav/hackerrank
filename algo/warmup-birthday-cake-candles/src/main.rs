use std::io::{self, BufRead};

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

fn birthday_cake_candles(candles: &[u64]) -> u64 {
    let len: usize = candles.len();
    let (mut max, mut max_cnt): (u64, u64) = (0, 0);
    for i in 0..len as usize {
        if candles[i] == max {
            max_cnt += 1;
        } else if candles[i] > max {
            max = candles[i];
            max_cnt = 1;
        }
    }
    return max_cnt;
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let _candles_count = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim()
        .parse::<u64>()
        .unwrap();

    let candles: Vec<u64> = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim_end()
        .split(' ')
        .map(|s| s.to_string().parse::<u64>().unwrap())
        .collect();

    let result = birthday_cake_candles(&candles);

    println!("{}", result);
}
