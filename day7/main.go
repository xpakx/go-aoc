package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code, day 7")
	fmt.Println("=====================")
	input := GetInput("input.txt")
	first := solve(input)
	fmt.Print("*  ")
	fmt.Println(first)
}

func solve(lines []string) int {
	hands := ParseHands(lines)
	for _, hand := range hands {
		fmt.Print("Hand:")
		for _, card := range hand.cards {
			fmt.Print(" ", card)
		}
		fmt.Println()
		fmt.Println("Bid: ", hand.bid)
	}
	return 0
}

func GetInput(filename string) []string {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	readFile.Close()
	return lines
}

func ParseHands(lines []string) []Hand {
	var hands []Hand
	for _, line := range lines {
		hand := ParseHand(line)
		hands = append(hands, hand)
	}
	return hands
}

func ParseHand(line string) Hand {
	splitted := strings.Split(line, " ")
	cards := strings.Split(splitted[0], "")
	bid, _ := strconv.Atoi(splitted[1])
	return Hand{cards, bid}
}

const (
	HighCard int = 0
	Pair = 1
	TwoPairs = 2
	Three = 3
	House = 4
	Four = 5
	Five = 6
)

type Hand struct {
	cards []string
	bid int
	// result int
}
