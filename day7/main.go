package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {
	fmt.Println("Advent of Code, day 7")
	fmt.Println("=====================")
	input := GetInput("input.txt")
	first := solveFirst(input)
	second := solveSecond(input)
	fmt.Print("*  ")
	fmt.Println(first)
	fmt.Print("** ")
	fmt.Println(second)
}

func solveFirst(lines []string) int {
	hands := ParseHands(lines)
	sort.Slice(hands, func(i, j int) bool {
		return compare(hands[i], hands[j], false)
	})
	result := 0
	for i := range hands {
		result += (i+1)*hands[i].bid
	}
	return result
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

func compare(hand1 Hand, hand2 Hand, jokers bool) bool {
	if hand1.handType != hand2.handType {
		return hand1.handType < hand2.handType
	}
	for i := range hand1.cards {
		if hand1.cards[i] != hand2.cards[i] {
			if jokers {
				return compareCardsWithJokers(hand1.cards[i], hand2.cards[i])
			} else {
				return compareCards(hand1.cards[i], hand2.cards[i])
			}
		}
	}
	return false
}

func compareCards(card1 string, card2 string) bool {
	cards := map[string]int{
		"2": 0,
		"3": 1,
		"4": 2,
		"5": 3,
		"6": 4,
		"7": 5,
		"8": 6,
		"9": 7,
		"T": 8,
		"J": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}
	return cards[card1] < cards[card2]
}

func compareCardsWithJokers(card1 string, card2 string) bool {
	cards := map[string]int{
		"J": 0,
		"2": 1,
		"3": 2,
		"4": 3,
		"5": 4,
		"6": 5,
		"7": 6,
		"8": 7,
		"9": 8,
		"T": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}
	return cards[card1] < cards[card2]
}

func CalculateHandTypeWithJokers(cards []string) int {
	hash := make(map[string]int)
	jokers := 0
	for _, card := range cards {
		if card == "J" {
			jokers += 1
		} else if _, ok := hash[card]; ok {
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
		if jokers > 0 {
			return Five
		}
		return Four
	}
	if values[3] == 1 && values[2] == 1 {
		return House
	}
	if values[3] == 1 {
		if jokers == 1 {
			return Four
		}
		if jokers > 1 {
			return Five
		}
		return Three
	}
	if values[2] == 2 {
		if jokers > 0 {
			return House
		}
		return TwoPairs
	}
	if values[2] == 1 {
		if jokers == 1 {
			return Three
		}
		if jokers == 2 {
			return Four
		}
		if jokers > 2 {
			return Five
		}
		return Pair
	}
	if jokers > 3 {
		return Five
	}
	if jokers == 3 {
		return Four
	}
	if jokers == 2 {
		return Three
	}
	if jokers == 1 {
		return Pair
	}
	return HighCard
}

func solveSecond(lines []string) int {
	hands := ParseHands(lines)
	for i := range hands {
		hands[i].handType = CalculateHandTypeWithJokers(hands[i].cards)
	}

	sort.Slice(hands, func(i, j int) bool {
		return compare(hands[i], hands[j], true)
	})
	result := 0
	for i := range hands {
		result += (i+1)*hands[i].bid
	}
	return result
}
