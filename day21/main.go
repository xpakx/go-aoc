package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code, day 20")
	fmt.Println("=====================")
	start, board := GetStart(Parse("input.txt"))

	fmt.Print("*  ")
	fmt.Println(SolveFirst(board, start, 64, 0))
	result := SolveSecond(board, start, 26501365)
	fmt.Print("**  ")
	fmt.Println(result)
}

func CalcFromCorner(board [][]rune, steps int, modulo int) int {
	result :=  Solve(board, Pos{0,0}, steps, modulo)  
	result += Solve(board, Pos{len(board)-1, len(board)-1}, steps, modulo)
	result += Solve(board, Pos{len(board)-1, 0}, steps, modulo)
	result += Solve(board, Pos{0, len(board)-1}, steps, modulo)
	return result
}

func CalcFromEdge(board [][]rune, start int, steps int, modulo int) int {
	t2 := Solve(board, Pos{len(board)-1, start}, steps, modulo)
	t2 += Solve(board, Pos{0, start}, steps, modulo)
	t2 += Solve(board, Pos{start, len(board)-1}, steps, modulo)
	t2 += Solve(board, Pos{start, 0}, steps, modulo)
	return t2
}

type Solver struct {
	board [][]rune
	evenBoard int
	oddBoard int
	corners int
	invCorners int
	ends int
}

func Prepare(board [][]rune, start int, startPos Pos) Solver {
	return Solver{
		board,
		SolveFirst(board, startPos, 1000, 0),
		SolveFirst(board, startPos, 1000, 1),
		CalcFromCorner(board, start+1, 0),
		CalcFromCorner(board, len(board)+start, 1),
		CalcFromEdge(board, start, len(board), 1),
	}
}

func (s Solver) Solve(n int) int {
	result := 0
	result = (n-1)*(n-1)*s.oddBoard
	result += n*n*s.evenBoard
	result += n*s.corners
	result += (n-1)*s.invCorners
	result += s.ends
	return result
}

func SolveSecond(board [][]rune, startPos Pos, steps int) int {
	// Because there are very little obstacles on the board, the explored “boards” are diamond-shaped
	// and the amount of full boards we visited is equal to the number of dots within a given taxicab distance
	// So explored area is f(s) = s² + (s-1)². 


	// The board is a square of 131/131 rows, and row/column we start on are without any obstacles. Because 
	// width/height is odd and we start at the center, the next square will have reachable non-obstacle 
	// spaces inverted. 

	// What's more, because the visitable fields depend on remaining steps s² is a number of boards 
	// with “even” parity (i.e. reachable from start when steps are even)
	// and (s-1)² of boards with “odd” parity (i.e. reachable from start when steps are odd)
	fmt.Println("Steps modulo length =", steps%(len(board)))


	// steps from puzzle are n*131+65, so we could look only at distances conforming to this equation
	// Because in the cornes/inverted corners and at the end of path remaining steps will be always the same,
	// regardless of n, the parity of these boards will be always the same. Because there are very little 
	// obstacles and steps always reach edge/corner at the same point and with the same remaining steps,
	// we could just precompute incomplete boards
	length := len(board)
	start := length/2
	t65 := SolveFirst(board, startPos, 65, 0)
	solver := Prepare(board, start, startPos)

	// To make sure it works
	fmt.Println("after 65", t65, "steps")
	for i:=1;i<=3;i++ {
		fmt.Println("after", i*length+start, solver.Solve(i), "steps")
	}

	return solver.Solve((steps-start)/length)
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

func GetStart(board [][]rune) (Pos, [][]rune) {
	for i, row := range board {
		for j := range row {
			if row[j] == 'S' {
				board[i][j] = '.'
				return Pos{i, j}, board
			}
		}
	}
	return Pos{0, 0}, board
}

func SolveFirst(board [][]rune, start Pos, maxSteps int, modulo int) int {
	return Solve(board, start, maxSteps+1, (modulo+1)%2) // we must make one step from „nowhere” to start to preserve numbering
}

func Solve(board [][]rune, start Pos, maxSteps int, modulo int) int {
	visited := make([][]int, len(board))
	for i, row := range board {
		visited[i] = make([]int, len(row))
		for j := range row {
			visited[i][j] = -1
			if row[j] == 'S' {
				board[i][j] = '.'
			}
		}
	}

	list := make([]Pos, 0)
	list = append(list, start)
	for steps:=0; steps<maxSteps; steps++ {
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
			if visited[i][j] % 2 == (modulo + maxSteps) % 2 {
				counter++
			}
		}
	}
	return counter
}
