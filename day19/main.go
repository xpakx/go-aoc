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
	fmt.Println(SolveFirst(workflows, parts))
	fmt.Print("** ")
	fmt.Println(SolveSecond(workflows))
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

func SolveFirst(workflows map[string]Workflow, parts []Part) int {
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


func SolveSecond(workflows map[string]Workflow) int {
	id := "in"
	boundaries := Boundaries{0, 4001, 0, 4001, 0, 4001, 0, 4001}
	return SolveRec(workflows, id, boundaries)
}

type Boundaries struct {
	xMin int
	xMax int
	mMin int
	mMax int
	aMin int
	aMax int
	sMin int
	sMax int
}

func (bnd Boundaries) change(field string, cond_type int, cond_value int) Boundaries {
	xMin := bnd.xMin
	if field == "x" && cond_type == Higher && cond_value > xMin {
		xMin = cond_value
	}
	xMax := bnd.xMax
	if field == "x" && cond_type == Lower && cond_value < xMax {
		xMax = cond_value
	}
	mMin := bnd.mMin
	if field == "m" && cond_type == Higher && cond_value > mMin {
		mMin = cond_value
	}
	mMax := bnd.mMax
	if field == "m" && cond_type == Lower && cond_value < mMax {
		mMax = cond_value
	}
	aMin := bnd.aMin
	if field == "a" && cond_type == Higher && cond_value > aMin {
		aMin = cond_value
	}
	aMax := bnd.aMax
	if field == "a" && cond_type == Lower && cond_value < aMax {
		aMax = cond_value
	}
	sMin := bnd.sMin
	if field == "s" && cond_type == Higher && cond_value > sMin {
		sMin = cond_value
	}
	sMax := bnd.sMax
	if field == "s" && cond_type == Lower && cond_value < sMax {
		sMax = cond_value
	}
	return Boundaries{xMin, xMax, mMin, mMax, aMin, aMax, sMin, sMax}
}

func (bnd Boundaries) result() int {
	return (bnd.xMax-bnd.xMin-1)*(bnd.mMax-bnd.mMin-1)*(bnd.aMax-bnd.aMin-1)*(bnd.sMax-bnd.sMin-1)
}

func SolveRec(workflows map[string]Workflow, id string, boundaries Boundaries) int {
	workflow := workflows[id]
	result := 0
	bnd := boundaries
	for _, rule := range workflow.rules {
		if rule.cond_type == 0 {
			if rule.accept {
				result += bnd.result()
			} else if !rule.reject {
				result += SolveRec(workflows, rule.address, bnd)
			}
			break
		} else {
			newBoundaries := bnd.change(rule.field, rule.cond_type, rule.cond_value) 
			if rule.accept {
				result += newBoundaries.result()
			} else if ! rule.reject {
				result += SolveRec(workflows, rule.address, newBoundaries)
			}
			bnd = bnd.change(rule.field, -rule.cond_type, rule.cond_value+rule.cond_type)
		}
	}
	return result
}
