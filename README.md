# Sudoku solver
Doesn't everyone have their own version of a sudoku solver? Well here's mine.

# Usage
```bash
$ cat inputs/sudoku.1.in | cargo run
```

# How it works
Backtracking algorithm

# Example run
```shell
$ cat inputs/sudoku.1.in | cargo run
Input
5.. 467 3.9
9.3 81. 427
174 2.3 ...

231 976 854
857 124 .9.
496 3.8 172

... .89 26.
782 641 ..5
.1. ... 7.8

Output
528 467 319
963 815 427
174 293 586

231 976 854
857 124 693
496 358 172

345 789 261
782 641 935
619 532 748
```
