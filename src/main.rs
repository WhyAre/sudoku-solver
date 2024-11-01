use std::io::stdin;
use std::io::BufRead;

const EMPTY: char = '.';

fn print_board(board: &Vec<Vec<char>>) {
    board.iter().enumerate().for_each(|(row_ind, row)| {
        if row_ind > 0 && row_ind % 3 == 0 {
            println!("");
        }

        row.iter().enumerate().for_each(|(col_ind, val)| {
            if col_ind > 0 && col_ind % 3 == 0 {
                print!(" ");
            }
            print!("{}", val);
        });

        println!("");
    });
}

fn is_valid(r: usize, c: usize, val: char, board: &Vec<Vec<char>>) -> bool {
    // Row
    if board[r].iter().any(|&x| x == val) {
        return false;
    }

    // Col
    if board.iter().any(|x| x[c] == val) {
        return false;
    }

    // Check small box
    let start_r = r / 3 * 3;
    let start_c = c / 3 * 3;

    (start_r..start_r + 3)
        .flat_map(|r| {
            (start_c..start_c + 3)
                .map(|c| board[r][c])
                .collect::<Vec<_>>()
        })
        .all(|e| e != val)
}

fn solved(r: usize, c: usize, rows: usize) -> bool {
    return r == rows && c == 0;
}

fn get_next_cell(r: usize, c: usize, cols: usize) -> (usize, usize) {
    if c + 1 == cols {
        (r + 1, 0)
    } else {
        (r, c + 1)
    }
}

fn solve(r: usize, c: usize, board: &mut Vec<Vec<char>>) -> bool {
    if solved(r, c, board.len()) {
        print_board(board);
        return true;
    }

    if board[r][c] != EMPTY {
        let (nr, nc) = get_next_cell(r, c, board[r].len());
        return solve(nr, nc, board);
    }

    ('1'..='9').any(|val| {
        if !is_valid(r, c, val, board) {
            return false;
        }

        let (nr, nc) = get_next_cell(r, c, board[r].len());

        board[r][c] = val;
        let val = solve(nr, nc, board);
        board[r][c] = EMPTY;

        val
    })
}

fn main() {
    let iter = stdin().lock();
    let mut board: Vec<Vec<char>> = iter
        .lines()
        .map(|line| line.expect("Error reading line").trim().chars().collect())
        .collect();
    println!("Input");
    print_board(&board);

    println!("\nOutput");
    let has_sol = solve(0, 0, &mut board);
    if !has_sol {
        println!("No solution");
    }
}
