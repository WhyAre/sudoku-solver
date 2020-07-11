package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const noOfRow = 9

// Check whether val can be placed on board[y][x] without causing any problems
func isValidMove(board [][]int, x int, y int, val int) bool {
	// Loop row
	for i := 0; i < len(board[y]); i++ {
		if board[y][i] == val {
			return false
		}
	}

	// Loop column
	for i := 0; i < len(board); i++ {
		if board[i][x] == val {
			return false
		}
	}

	// Check block
	x = (x / 3) * 3
	y = (y / 3) * 3
	for i := y; i < y+3; i++ {
		for n := x; n < x+3; n++ {
			if board[i][n] == val {
				return false
			}
		}
	}
	return true
}

func findEmptyCell(board [][]int) (int, int, bool) {
	var x, y int
	for y = 0; y < len(board); y++ {
		for x = 0; x < len(board[y]); x++ {
			// Ignore cells that are filled
			if board[y][x] == 0 {
				return x, y, true
			}
		}
	}
	return 0, 0, false
}

func solve(board [][]int) {
	x, y, ok := findEmptyCell(board)
	if !ok { // If cannot find an empty cell, that means that the puzzle is solved
		printSlice(board)
		return
	}

	// Test all possible values
	for testVal := 1; testVal <= 9; testVal++ {
		if isValidMove(board, x, y, testVal) {
			board[y][x] = testVal
			solve(board)
			board[y][x] = 0
		}
	}
}

func printSlice(board [][]int) {
	fmt.Println("\nSolution")
	fmt.Println(strings.Repeat("=", 19))
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			fmt.Printf("%d ", board[y][x])
			if (x+1)%3 == 0 {
				fmt.Print(" ")
			}
		}
		if (y+1)%3 == 0 {
			fmt.Println()
		}
		fmt.Println()
	}
}

func main() {
	board := [][]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < noOfRow; i++ {
		fmt.Printf("Row %d > ", i+1)
		if !scanner.Scan() {
			return
		}
		row := []int{}
		line := scanner.Text()
		for _, v := range line {
			val, err := strconv.Atoi(string(v))
			if err != nil {
				log.Println("Error reading board")
				return
			}

			row = append(row, val)
		}
		board = append(board, row)
	}
	solve(board)
}
