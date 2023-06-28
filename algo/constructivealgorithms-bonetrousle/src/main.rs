use std::io::BufRead;

fn bonetrousle(k: u128, b: u128, n: u128) {
    let min: u128 = (1..=b).sum::<u128>();
    let max: u128 = (k - b + 1..=k).sum::<u128>();

    if n < min || n > max {
        println!("-1");
        return;
    }

    let mut sum: u128 = min;
    let mut res: Vec<u128> = (1..b + 1).collect();
    let mut cur: u128 = b - 1; // current block

    loop {
        let c = (n - sum) / (k - b + 1);
        if c > 0 && sum < n {
            sum -= res[cur as usize];
            sum += k - (b - cur - 1);
            res[cur as usize] = k - (b - cur - 1);
            cur -= 1;
        } else {
            break;
        }
    }

    res[cur as usize] += n - sum;

    for p in 0..b {
        print!("{}", res[p as usize]);
        if p != b - 1 {
            print!(" ");
        }
    }
    print!("\n");
}

fn main() {
    let stdin = std::io::stdin();
    let mut input = stdin.lock().lines();

    let t = input.next().unwrap().unwrap().trim().parse::<u8>().unwrap();

    assert_eq!(u128::BITS, 128);

    for _ in 0..t {
        let v: Vec<u128> = input
            .next()
            .unwrap()
            .unwrap()
            .trim()
            .split(' ')
            .map(|s| s.to_string().parse::<u128>().unwrap())
            .collect();

        let (n, k, b): (u128, u128, u128) = (v[0], v[1] as u128, v[2] as u128);

        bonetrousle(k, b, n);
    }
}
