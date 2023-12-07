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
		fmt.Println("Type: ", hand.PrintableHandType())
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
	handType := CalculateHandType(cards)
	return Hand{cards, bid, handType}
}

func CalculateHandType(cards []string) int {
	hash := make(map[string]int)
	for _, card := range cards {
		if _, ok := hash[card]; ok {
			hash[card] += 1
		} else {
			hash[card] = 1
		}
	}
	values := make([]int, 6)
	for i := range values {
		values[i] = 0
	}

	for i := range hash {
		values[hash[i]] += 1
	}

	if values[5] == 1 {
		return Five
	}
	if values[4] == 1 {
		return Four
	}
	if values[3] == 1 && values[2] == 1 {
		return House
	}
	if values[3] == 1 {
		return Three
	}
	if values[2] == 2 {
		return TwoPairs
	}
	if values[2] == 1 {
		return Pair
	}
	return HighCard
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
	handType int
}

func (hand Hand) PrintableHandType() string {
	cards := map[int]string{
		HighCard: "High card",
		Pair: "Pair",
		TwoPairs: "Two pairs",
		Three: "Three of a kind",
		House: "Full house",
		Four: "Four of a kind",
		Five: "Five of a kind",
	}
	return cards[hand.handType]
}
