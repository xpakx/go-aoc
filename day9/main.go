package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code, day 9")
	fmt.Println("=====================")
	first := solve("input.txt")
	fmt.Print("*  ")
	fmt.Println(first)
}

func solve(filename string) int {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		nums := ParseList(line)
		prediction := Predict(nums)
		fmt.Println(prediction)
		fmt.Println()
		result += prediction

	}
	readFile.Close()

	return result
}

func ParseList(list string) []int {
	splitted := strings.Split(strings.Join(strings.Fields(list), " "), " ")
	listFin := []int{}
	for _, x := range splitted {
		n, _ := strconv.Atoi(x)
		listFin = append(listFin, n)
	}
	return listFin
}

func Predict(list []int) int {
	zeroes := false
	lists := make([][]int, 0)
	lists = append(lists, list)
	for !zeroes {
		oldList := lists[len(lists)-1]
		lastLen := len(oldList)
		zeroes = true
		newList := make([]int, lastLen-1)
		for i:=0; i<lastLen-1; i++ {
			newList[i] = oldList[i+1]-oldList[i]
			if newList[i] != 0 {
				zeroes= false
			}
		}
		lists = append(lists, newList)
	}
	last := 0
	for i:=len(lists)-2; i>=0; i-- {
		fmt.Println(last)
		last = lists[i][len(lists[i])-1] + last

	}
	return last
}
