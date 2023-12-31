package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code, day 22")
	fmt.Println("=====================")
	blocks := Parse("input.txt")

	first, second := Solve(blocks)
	fmt.Print("*  ")
	fmt.Println(first)
	fmt.Print("** ")
	fmt.Println(second)
}

func Parse(filename string) map[int][]Block {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := make(map[int][]Block, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		block := GetBlock(line)
		lowerZ := Min(block.coordA.z, block.coordB.z)
		result[lowerZ] = append(result[lowerZ], block)
	}
	readFile.Close()

	return result
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func GetBlock(line string) Block {
	coords := strings.Split(line, "~")
	coordA := strings.Split(coords[0], ",")
	coordB := strings.Split(coords[1], ",")
	return Block{
		ListToPos(coordA),
		ListToPos(coordB),
	}
}

func ListToPos(pos []string) Pos {
	x, _ := strconv.Atoi(pos[0])
	y, _ := strconv.Atoi(pos[1])
	z, _ := strconv.Atoi(pos[2])
	return Pos{x, y, z}
}

type Pos struct {
	x int
	y int
	z int
}

type Block struct {
	coordA Pos
	coordB Pos
}

func Intersect(b1, b2 Block) bool {
	x1 := b1.coordA.x
	y1 := b1.coordA.y
	x2 := b1.coordB.x
	y2 := b1.coordB.y
	x3 := b2.coordA.x
	y3 := b2.coordA.y
	x4 := b2.coordB.x
	y4 := b2.coordB.y

	if x1 == x2 && y1 == y2 {
		if x3 == x4 && y3 == y4 {
			return x1 == x3 && y1 == y3
		}
		return PointOnLine(x1, y1, x3, y3, x4, y4) 
	}

	if x3 == x4 && y3 == y4 {
		return PointOnLine(x3, y3, x1, y1, x2, y2) 
	}

	denominator := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if denominator == 0 {
		if (BetweenI(x1, x3, x4) && BetweenI(y1, y3, y4)) || (BetweenI(x2, x3, x4) && BetweenI(y2, y3, y4)) {
			return true
		}
		if (BetweenI(x3, x1, x2) && BetweenI(y3, y1, y2)) || (BetweenI(x4, x1, x2) && BetweenI(y4, y1, y2)) {
			return true
		}
		return false
	}

	a := (x1*y2 - y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)
	b := (x1*y2 - y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)
	
	x0 := float64(a)/float64(denominator)
	y0 := float64(b)/float64(denominator)

	if Between(x0, x1, x2) && Between(y0, y1, y2) && Between(x0, x3, x4) && Between(y0, y3, y4) {
		return true
	}
	
	return false
}

func Between(x float64, x1, x2 int) bool {
	epsilon := 1E-03
	lower := Min(x1, x2)
	higher := Max(x1, x2)
	return x <= float64(higher) + epsilon && x >= float64(lower) - epsilon
}

func BetweenI(x, x1, x2 int) bool {
	lower := Min(x1, x2)
	higher := Max(x1, x2)
	return x <= higher && x >= lower
}

func PointOnLine(x, y, x1, y1, x2, y2 int) bool {
	zero := (x2 - x1) * (y - y1) - (x - x1) * (y2 - y1);
	if zero != 0 {
		return false
	}

	return y >= Min(y1, y2) && y <= Max(y1, y2) && x >= Min(x2, x1) && x <= Max(x2, x1)
}

func (block Block) MoveToZ(z int) Block {
	diff := Abs(block.coordA.z - block.coordB.z)
	if block.coordA.z > block.coordB.z {
		block.coordA.z = z+diff
		block.coordB.z = z
	} else {
		block.coordB.z = z+diff
		block.coordA.z = z
	}
	return block
}

func Solve(blocks map[int][]Block) (int, int) {
	stopped := make(map[int][]Block, 0)
	maxZ := 0

	for _, level := range blocks {
		for _, block := range level {
			currZ := Max(block.coordA.z, block.coordB.z)
			maxZ = Max(currZ, maxZ)
		}
	}


	supports := make(map[Block][]Block, 0)
	supportedBy := make(map[Block][]Block, 0)

	for z:=1; z<=maxZ; z++ {
		for _, block := range blocks[z] {
			changed := false
			for i:=z-1; i>=0; i-- {
				if i == 0 {
					block = block.MoveToZ(1)
					changed = true
				} 

				for _, b := range stopped[i] {
					if Intersect(b, block) {
						block = block.MoveToZ(i+1)
						supports[b] = append(supports[b], block)
						supportedBy[block] = append(supportedBy[block], b)
						changed = true
					}
				}

				if changed {
					currZ := Max(block.coordA.z, block.coordB.z)
					stopped[currZ] = append(stopped[currZ], block)
					break
				}
			}
		}
	}

	first := 0
	dangerous := make([]Block, 0)
	for _, level := range stopped {
		for _, block := range level {
			if len(supports[block]) == 0 {
				first++
			} else {
				ok := true
				for _, b := range supports[block] {
					if len(supportedBy[b]) == 1 {
						ok = false
						break
					}
				}
				if ok {
					first++
				} else {
					dangerous = append(dangerous, block)
				}
			}
		}
	}


	second := 0
	for _, block := range dangerous {
		lst := make([]Block, 0)
		all := make([]Block, 0)
		lst = append(lst, block)
		all = append(lst, block)
		for true {
			levelList := make([]Block, 0)
			for _, curr := range lst {
				for _, b := range supports[curr] {
					outsideSupport := false
					for _, supporter := range supportedBy[b] {
						supporterDesintegrated := false
						for _, last := range all {
							if last == supporter {
								supporterDesintegrated = true
								break
							}
						}
						if !supporterDesintegrated {
							outsideSupport = true
							break
						}
					}
					if !outsideSupport {
						added := false
						for _, onList := range all {
							if b == onList {
								added = true
								break
							}
						}
						if !added {
							levelList = append(levelList, b)
							all = append(all, b)
						}
					}
				}

			}
			second += len(levelList);
			lst = nil
			lst = levelList
			if len(levelList) == 0 {
				break
			}
		}
	}

	return first, second
}
