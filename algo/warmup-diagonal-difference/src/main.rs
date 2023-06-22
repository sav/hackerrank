use std::io::{self, BufRead};

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

fn diagonal_difference(matrix: &[Vec<i32>]) -> i32 {
    let rows: usize = matrix.len();
    let mut row: usize = 0;
    let mut primary: i32 = 0;
    let mut secondary: i32 = 0;
    while row < rows {
        let cols: &Vec<i32> = &matrix[row];
        assert_eq!(cols.len(), rows); // square matrix
        primary += cols[row];
        secondary += cols[rows - row - 1];
        row += 1;
    }
    return (primary - secondary).abs();
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let n = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim()
        .parse::<i32>()
        .unwrap();

    let mut arr: Vec<Vec<i32>> = Vec::with_capacity(n as usize);

    for i in 0..n as usize {
        arr.push(Vec::with_capacity(n as usize));

        arr[i] = stdin_iterator
            .next()
            .unwrap()
            .unwrap()
            .trim_end()
            .split(' ')
            .map(|s| s.to_string().parse::<i32>().unwrap())
            .collect();
    }

    println!("{}", diagonal_difference(&arr));
}
