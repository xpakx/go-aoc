package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

func main() {
	fmt.Println("Advent of Code, day 23")
	fmt.Println("=====================")
	start, end, graph := Parse("input.txt")

	first, second := Solve(start, end, graph)
	fmt.Print("*  ")
	fmt.Println(first)
	fmt.Print("** ")
	fmt.Println(second)
}

type Path struct {
	last Pos
	curr Pos
	length int
	from Pos
}

type Edge struct {
	to Pos
	length int
}

func Parse(filename string) (Pos, Pos, map[Pos][]Edge) {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	board := make([][]rune, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		board = append(board, []rune(line))
	}
	readFile.Close()

	start := Pos{0, 0}
	end := Pos{0, 0}
	for i, s := range board[0] {
		if s != '#' {
			start = Pos{0, i}
			break
		}
	}
	for i, s := range board[len(board)-1] {
		if s != '#' {
			end = Pos{len(board)-1, i}
			break
		}
	}

	visited := make(map[Pos]struct{}, 0)
	paths := make([]Path, 0)
	paths = append(paths, Path{start, start, 0, start})
	result := make(map[Pos][]Edge, 0)

	for len(paths) > 0 {
		newPaths := make([]Path, 0)
		for _, p := range paths {
			visited[p.curr] = struct{}{}

			if p.curr == end {
				fmt.Println("testend")
				result[p.from] = append(result[p.from], Edge{end, p.length})
				continue
			}

			left := Pos{p.curr.x, p.curr.y-1}
			right := Pos{p.curr.x, p.curr.y+1}
			up := Pos{p.curr.x-1, p.curr.y}
			down := Pos{p.curr.x+1, p.curr.y}
			if left.y >=0 && left != p.last && board[left.x][left.y] == '.' {
				newPaths = append(newPaths, Path{p.curr, left, p.length+1, p.from})
			} else if right.y <= len(board)-1 && right != p.last && board[right.x][right.y] == '.' {
				newPaths = append(newPaths, Path{p.curr, right, p.length+1, p.from})
			} else if down.x <= len(board)-1 && down != p.last && board[down.x][down.y] == '.' {
				newPaths = append(newPaths, Path{p.curr, down, p.length+1, p.from})
			} else if up.x >=0 && up != p.last && board[up.x][up.y] == '.' {
				newPaths = append(newPaths, Path{p.curr, up, p.length+1, p.from})
			} else if p.curr != p.last {
				if left.y >=0 && left != p.last && board[left.x][left.y] == '<' {
					node := Pos{left.x, left.y-1}
					result[p.from] = append(result[p.from], Edge{node, p.length+2})
					if _, ok := visited[node]; !ok {
						newPaths = append(newPaths, Path{node, node, 0, node})
					}
				} else if right.y <= len(board)-1 && right != p.last && board[right.x][right.y] == '>' {
					node := Pos{right.x, right.y+1}
					result[p.from] = append(result[p.from], Edge{node, p.length+2})
					if _, ok := visited[node]; !ok {
						newPaths = append(newPaths, Path{node, node, 0, node})
					}
				} else if down.x <= len(board)-1 && down != p.last && board[down.x][down.y] == 'v' {
					node := Pos{down.x+1, down.y}
					result[p.from] = append(result[p.from], Edge{node, p.length+2})
					if _, ok := visited[node]; !ok {
						newPaths = append(newPaths, Path{node, node, 0, node})
					}
				} else if up.x >=0 && up != p.last && board[up.x][up.y] == '^' {
					node := Pos{up.x-1, up.y}
					result[p.from] = append(result[p.from], Edge{node, p.length+2})
					if _, ok := visited[node]; !ok {
						newPaths = append(newPaths, Path{node, node, 0, node})
					}
				}
			} else if p.curr == p.last {
				if left.y >=0 && board[left.x][left.y] == '<' {
					node := Pos{left.x, left.y-1}
					newPaths = append(newPaths, Path{left, node, p.length+2, p.from})
				} 
				if right.y <= len(board)-1 && board[right.x][right.y] == '>' {
					node := Pos{right.x, right.y+1}
					newPaths = append(newPaths, Path{right, node, p.length+2, p.from})
				} 
				if down.x <= len(board)-1 && board[down.x][down.y] == 'v' {
					node := Pos{down.x+1, down.y}
					newPaths = append(newPaths, Path{down, node, p.length+2, p.from})
				} 
				if up.x >=0 && board[up.x][up.y] == '^' {
					node := Pos{up.x-1, up.y}
					newPaths = append(newPaths, Path{right, node, p.length+2, p.from})
				}
			}

		}
		paths = nil
		paths = newPaths
	}

	for i := range board {
		for j:= range board[i] {
		   if _, ok := visited[Pos{i,j}]; ok {
			   fmt.Print("O")
		   } else {
			   fmt.Print(string(board[i][j]))
		   }
		}
		fmt.Println()
	}


	fmt.Println(start, end, board)
	fmt.Println(result)
	fmt.Println(len(result))
	return start, end, result
}

type Pos struct {
	x int
	y int
}

func Solve(start, end Pos, graph map[Pos][]Edge) (int, int) {
	return BellmanFord(start, end, graph), 0
}

func BellmanFord(start Pos, end Pos, graph map[Pos][]Edge) int {

    distance := make(map[Pos]int, 0)

    for key := range graph {
	    distance[key] = math.MaxInt
    }
    
    distance[start] = 0

    for range graph {
	    for key := range graph {
		    for j := range graph[key] {
			    if distance[key] < math.MaxInt {
				    distance[graph[key][j].to] = Min(distance[graph[key][j].to], distance[key] - graph[key][j].length)
			    }
		    }
	    }
    }
    return -distance[end]

}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
