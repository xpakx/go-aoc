package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {
	fmt.Println("Advent of Code, day 15")
	fmt.Println("=====================")
	input := ParseFirst("input.txt")
	fmt.Print("*  ")
	fmt.Println(SolveFirst(input))
	input2 := ParseSecond("input.txt")
	fmt.Print("** ")
	fmt.Println(SolveSecond(input2))
}

func SolveFirst(input [][]rune) int {
	result := 0
	for _, code := range input {
		result += CalcHash(code)
	}
	return result
}

func CalcHash(value []rune) int {
	c := 0
	for i := range value {
		c = (c+int(value[i]))*17%256
	}
	return c
}

func ParseFirst(filename string) [][]rune {
	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}
	contentString := string(content)
	contentString  = strings.TrimSuffix(contentString, "\n")
	codes := strings.Split(contentString, ",")
	input := make([][]rune, 0)
	for i := range codes {
		input = append(input, []rune(codes[i]))
	}

	return input
}

type Instruction struct {
	hash int
	register string
	value int
	delete bool
}

func ParseSecond(filename string) []Instruction {
	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}
	contentString := string(content)
	contentString  = strings.TrimSuffix(contentString, "\n")
	codes := strings.Split(contentString, ",")
	input := make([]Instruction, 0)
	for i := range codes {
		splitByEqual := strings.Split(codes[i], "=")
		if len(splitByEqual) > 1 {
			instrCode := CalcHash([]rune(splitByEqual[0]))
			value, _ := strconv.Atoi(splitByEqual[1])
			instr := Instruction{instrCode, splitByEqual[0], value, false}
			input = append(input, instr)
		} else {
			splitByMinus := strings.Split(codes[i], "-")
			instrCode := CalcHash([]rune(splitByMinus[0]))
			instr := Instruction{instrCode, splitByMinus[0], 0, true}
			input = append(input, instr)
		}
	}
	return input
}

type SortedLens struct {
	box int
	focal_length int
	order int
}

func SolveSecond(instructions []Instruction) int {
	boxes := make(map[int]map[string]int, 0)
	orders := make(map[int]map[string]int, 0)
	max_order := make(map[int]int, 0)
	for _, inst := range instructions {
		if _, ok := boxes[inst.hash]; ok {
			if _, ok := boxes[inst.hash][inst.register]; ok && inst.delete {
				delete(boxes[inst.hash], inst.register)
				delete(orders[inst.hash], inst.register)
			} else if !inst.delete {
				if _, ok := boxes[inst.hash][inst.register]; !ok {
					ord := max_order[inst.hash] + 1
					orders[inst.hash][inst.register] = ord
					max_order[inst.hash] = ord
				}
				boxes[inst.hash][inst.register] = inst.value
			}
		} else if !inst.delete {
			lst := make(map[string]int, 0)
			lst[inst.register] = inst.value
			boxes[inst.hash] = lst
			lst2 := make(map[string]int, 0)
			orders[inst.hash] = lst2
			max_order[inst.hash] = 1
		}
	}

	result := 0
	for key, value := range boxes {
		box_number := key+1
		if len(value) > 0 {
			lens := make([]SortedLens, 0)
			for k, focal_length := range value {
				lens = append(lens, SortedLens{box_number, focal_length, orders[key][k]})
			}
			sort.Slice(lens, func(i, j int) bool {
				return lens[i].order < lens[j].order
			})
			for i, n := range lens {
				result += n.box*n.focal_length*(i+1)
			}
		}
	}
	return result
}
