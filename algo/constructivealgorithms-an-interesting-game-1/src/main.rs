use std::io::BufRead;

fn an_interesting_game_1(v: &mut Vec<i64>, player: u8) -> &str {
    let len = v.len();
    let (mut max, mut max_index): (i64, usize) = (v[0], 0);

    for i in 1..len as usize {
        if v[i] > max {
            max = v[i];
            max_index = i;
        }
    }

    for _ in max_index..len as usize {
        v.remove(max_index);
    }

    if v.len() > 0 {
        return an_interesting_game_1(v, (player + 1) % 2);
    }

    return if player == 0 { "ANDY" } else { "BOB" };
}

fn main() {
    let stdin = std::io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let _: u8 = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim()
        .parse::<u8>()
        .unwrap();

    let mut c: usize = 0;

    for line in stdin_iterator {
        if c % 2 == 1 {
            let mut v: Vec<i64> = line
                .unwrap()
                .trim()
                .split(' ')
                .map(|s| s.to_string().parse::<i64>().unwrap())
                .collect();
            println!("{}", an_interesting_game_1(&mut v, 1));
        }
        c += 1;
    }
}
