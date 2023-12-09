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
	first, second := solve("input.txt")
	fmt.Print("*  ")
	fmt.Println(first)
	fmt.Print("** ")
	fmt.Println(second)
}

func solve(filename string) (int, int) {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	first := 0
	second := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		nums := ParseList(line)
		prediction := Predict(nums)
		first += prediction
		secondPrediction := PredictBackward(nums)
		second += secondPrediction

	}
	readFile.Close()

	return first, second
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
		last = lists[i][len(lists[i])-1] + last

	}
	return last
}

func PredictBackward(list []int) int {
	reversed := make([]int, len(list))
	for i := range list {
		reversed[i] = list[len(list)-1-i]
	}
	return Predict(reversed)
}
