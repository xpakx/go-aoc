package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code, day 20")
	fmt.Println("=====================")
	board := Parse("input.txt")
	fmt.Println(board)
	fmt.Print("*  ")
	fmt.Println(SolveFirst(board, 64))
}

func Parse(filename string) [][]rune {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := make([][]rune, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		result = append(result, []rune(line))
	}
	readFile.Close()

	return result
}

type Pos struct {
	i int
	j int
}

func SolveFirst(board [][]rune, maxSteps int) int {
	visited := make([][]int, len(board))
	start := Pos{0, 0}
	for i, row := range board {
		visited[i] = make([]int, len(row))
		for j := range row {
			visited[i][j] = -1
			if row[j] == 'S' {
				start.i, start.j = i, j
			}
		}
	}

	list := make([]Pos, 0)
	list = append(list, start)
	for steps:=0; steps<=maxSteps; steps++ {
		newList := make([]Pos, 0)
		for _, pos := range list {
			if visited[pos.i][pos.j] >= 0 {
				continue
			}
			visited[pos.i][pos.j] = steps
			if pos.i > 0 && board[pos.i-1][pos.j] == '.' && visited[pos.i-1][pos.j] < 0 {
				newList = append(newList, Pos{pos.i-1, pos.j})
			}
			if pos.i < len(board)-1 && board[pos.i+1][pos.j] == '.' && visited[pos.i+1][pos.j] < 0 {
				newList = append(newList, Pos{pos.i+1, pos.j})
			}
			if pos.j > 0 && board[pos.i][pos.j-1] == '.' && visited[pos.i][pos.j-1] < 0 {
				newList = append(newList, Pos{pos.i, pos.j-1})
			}
			if pos.j < len(board[0])-1 && board[pos.i][pos.j+1] == '.' && visited[pos.i][pos.j+1] < 0 {
				newList = append(newList, Pos{pos.i, pos.j+1})
			}
		}
		list = nil
		list = newList

	}

	counter := 0
	for i := range board {
		for j := range board[i] {
			if visited[i][j] % 2 == 0 {
				fmt.Print("O")
				counter++
			} else {
				fmt.Print(string(board[i][j]))
			}
		}
		fmt.Println()
	}
	fmt.Println(counter)
	return 0
}
