use std::io::BufRead;

// Try bubble sort with local counters.
// space: O(n), time: O(n^2) with amortized average performance
// due to the maximum swaps constraint.

fn new_year_chaos(v: &mut Vec<u64>) -> u64 {
    let len: usize = v.len();

    if len < 2 {
        return 0;
    }

    let mut cnt: Vec<u8> = vec![0; len + 1];
    let mut chg: bool;
    let mut tot: u64 = 0;

    loop {
        chg = false;

        for i in 0..(len - 1) as usize {
            if v[i] > v[i + 1] {
                cnt[v[i] as usize] += 1;
                if cnt[v[i] as usize] > 2 {
                    return std::u64::MAX; // Too chaotic
                }
                tot += 1;
                v.swap(i, i + 1);
                chg = true;
            }
        }

        if !chg {
            break;
        }
    }

    return tot;
}

fn main() {
    let stdin = std::io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let _ = stdin_iterator.next();
    let mut c: u8 = 0;

    for line in stdin_iterator {
        if c == 1 {
            let mut v: Vec<u64> = line
                .unwrap()
                .trim()
                .split(' ')
                .map(|s| s.to_string().parse::<u64>().unwrap())
                .collect();
            let n: u64 = new_year_chaos(&mut v);
            if n != std::u64::MAX {
                println!("{}", n);
            } else {
                println!("Too chaotic");
            }
        }
        c = (c + 1) % 2;
    }
}
