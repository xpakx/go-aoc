package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code, day 19")
	fmt.Println("=====================")
	workflows, parts := Parse("input.txt")
	fmt.Print("*  ")
	fmt.Println(Solve(workflows, parts))
}

func Parse(filename string) (map[string]Workflow, []Part) {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	workflows := make(map[string]Workflow, 0)
	parts := make([]Part, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}
		if line[0] == '{' {
			parts = append(parts, ParsePart(line))
		} else {
			workflow := ParseWorkflow(line)
			workflows[workflow.id] = workflow
		}
	}
	readFile.Close()
	return workflows, parts
}

type Workflow struct {
	id  string
	rules []Rule
}

const (
	Lower = -1
	Higher = 1
	Jump = 0
)

type Rule struct {
	accept bool
	reject bool
	address string
	field string
	cond_type int
	cond_value int
}

type Part struct {
	x int
	m int
	a int
	s int
}

func ParseWorkflow(line string) Workflow {
	reg := regexp.MustCompile("^([a-z]+){(.+)}$")
	match := reg.FindStringSubmatch(line)
	id := match[1]
	rules := make([]Rule, 0)
	rulesString := strings.Split(match[2], ",")
	for _, rule := range rulesString {
		rule_split := strings.Split(rule, ":")
		if len(rule_split) == 1 {
			if rule_split[0] == "A" {
				rules = append(rules, Rule{true, false, "", "", Jump, 0})
			} else if rule_split[0] == "R" {
				rules = append(rules, Rule{false, true, "", "", Jump, 0})
			} else {
				rules = append(rules, Rule{false, false, rule_split[0], "", Jump, 0})
			}
		} else {
			accept := false 
			reject := false 
			address := ""
			if rule_split[1] == "A" {
				accept = true
			} else if rule_split[1] == "R" {
				reject = true
			} else {
				address = rule_split[1]
			}
			cond_type := Jump
			cond_value := 0
			field := ""
			if strings.Contains(rule_split[0], "<") {
				cond := strings.Split(rule_split[0], "<")
				field = cond[0]
				cond_type = Lower
				cond_value, _ = strconv.Atoi(cond[1])

			} else {
				cond := strings.Split(rule_split[0], ">")
				field = cond[0]
				cond_type = Higher
				cond_value, _ = strconv.Atoi(cond[1])
			}


			rules = append(rules, Rule{accept, reject, address, field, cond_type, cond_value})
		}
	}
	return Workflow{id, rules}
}


func ParsePart(line string) Part {
	reg := regexp.MustCompile("^{x=(\\d+),m=(\\d+),a=(\\d+),s=(\\d+)}$")
	match := reg.FindStringSubmatch(line)
	x, _ := strconv.Atoi(match[1])
	m, _ := strconv.Atoi(match[2])
	a, _ := strconv.Atoi(match[3])
	s, _ := strconv.Atoi(match[4])
	return Part{x, m, a, s}
}

func Solve(workflows map[string]Workflow, parts []Part) int {
	result := 0
	
	for _, part := range parts {
		id := "in"
		finished := false 
		for !finished {
			workflow := workflows[id]
			for _, rule := range workflow.rules {
				apply := rule.cond_type == 0
				if rule.cond_type == 1 {
					if rule.field == "x" {
						apply = part.x > rule.cond_value
					} else if rule.field == "m" {
						apply = part.m > rule.cond_value
					} else if rule.field == "a" {
						apply = part.a > rule.cond_value
					} else if rule.field == "s" {
						apply = part.s > rule.cond_value
					}
				} else {
					if rule.field == "x" {
						apply = part.x < rule.cond_value
					} else if rule.field == "m" {
						apply = part.m < rule.cond_value
					} else if rule.field == "a" {
						apply = part.a < rule.cond_value
					} else if rule.field == "s" {
						apply = part.s < rule.cond_value
					}
				}
				if apply {
					if rule.accept {
						finished = true
						result += part.x+part.m+part.a+part.s
						break
					}
					if rule.reject {
						finished = true
						break
					}
					id = rule.address
					break
				}
			}

		}
	}
	return result
}
