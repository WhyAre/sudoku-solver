# Sudoku solver
Doesn't everyone have their own version of a sudoku solver? Well here's mine. The main purpose is to play around with backtracking algorithm and solving sudoku puzzles is just the perfect problem to solve.

# Usage 
Compile and run the code:
```
go run run.go
```

Run the executable (This might not work for people not using ubuntu because the code is compiled in ubuntu)
```
./run
```

The program will then prompt for the rows

For people who makes mistakes often, like me, can opt to write it all in a file and run like this
```bash
cat file_name_of_board | go run .
```

# How it works
Well it uses some backtracking algorithm

# Speed
Written in golang to make it fast (Maybe C++ and rust might be faster but I'm just a gopher for now)
```shell
time cat board-config | ./run 
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


real    0m0.004s
user    0m0.005s
sys     0m0.000s
```
