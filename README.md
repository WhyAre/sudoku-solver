# Sudoku solver
Doesn't everyone have their own version of a sudoku solver? Well here's mine.

# Usage 
```bash
$ cat board-config | go run .
```

# How it works
Backtracking algorithm

# Example run
```shell
$ cat board-config | go run .
Solution
===================
5 3 4  6 7 8  1 9 2  
6 7 2  1 9 5  3 4 8  
1 9 8  3 4 2  5 6 7  

8 5 9  7 6 1  4 2 3  
4 2 6  8 5 3  9 7 1  
7 1 3  9 2 4  8 5 6  

9 6 1  5 3 7  2 8 4  
2 8 7  4 1 9  6 3 5  
3 4 5  2 8 6  7 1 9  


Solution
===================
5 3 4  6 7 8  9 1 2  
6 7 2  1 9 5  3 4 8  
1 9 8  3 4 2  5 6 7  

8 5 9  7 6 1  4 2 3  
4 2 6  8 5 3  7 9 1  
7 1 3  9 2 4  8 5 6  

9 6 1  5 3 7  2 8 4  
2 8 7  4 1 9  6 3 5  
3 4 5  2 8 6  1 7 9  
```
