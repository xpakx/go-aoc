package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	//"strings"
)

func main() {
	fmt.Println("Advent of Code, day 20")
	fmt.Println("=====================")
	modules := Parse("input.txt")
	fmt.Println(modules)
	fmt.Print("*  ")
	fmt.Println(SolveFirst(modules))
	modules = Parse("input.txt")
	fmt.Print("** ")
	fmt.Println(SolveSecond(modules))
}

func Parse(filename string) map[string]Module {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	modules := make(map[string]Module, 0)
	inputs := make(map[string][]string, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitted := strings.Split(line, " -> ")
		id := splitted[0]
		addresses := strings.Split(splitted[1], ", ")
		if line[0] == '&' {
			id = id[1:]
			modules[id] = ConjunctionModule{make(map[string]int, 0), addresses}
		} else if line[0] == '%' {
			id = id[1:]
			modules[id] = FlipFlopModule{false, addresses}
		} else {
			modules[id] = BroadcastModule{addresses}
		}

		for _, a := range addresses {
			if _, ok := inputs[a]; !ok {
				inputs[a] = make([]string, 0)
			}
			inputs[a] = append(inputs[a], id)
		}
	}
	readFile.Close()

	for key, module := range modules {
		if con, ok := module.(ConjunctionModule); ok {
			if _, ok := inputs[key]; ok {
				for _, in := range inputs[key] {
					con.last[in] = Low
				}
			}
		}
	}
	return modules
}

type Output struct {
	id  string
	signal int
}

type Input struct {
	id string
	signal int
}

const (
	Low  = 0
	High = 1
)

type Module interface {
	Proccess(signal Input) (Module, []Output)
	Successors() []string
}

type FlipFlopModule struct {
	on bool
	addresses []string
}

func (module FlipFlopModule) Proccess(signal Input) (Module, []Output) {
	result := make([]Output, 0)
	if signal.signal == High {
		return module, result
	}
	ret := Low 
	if !module.on {
		ret = High
	}
	module.on = !module.on
	for _, addr := range module.addresses {
		result = append(result, Output{addr, ret})
	}
	return module, result
}

func (module FlipFlopModule) Successors() []string {
	return module.addresses
}

type BroadcastModule struct {
	addresses []string
}

func (module BroadcastModule) Proccess(signal Input) (Module, []Output) {
	result := make([]Output, 0)
	for _, addr := range module.addresses {
		result = append(result, Output{addr, signal.signal})
	}
	return module, result
}

func (module BroadcastModule) Successors() []string {
	return module.addresses
}

type ConjunctionModule struct {
	last map[string]int
	addresses []string
}

func (module ConjunctionModule) Proccess(signal Input) (Module, []Output) {
	result := make([]Output, 0)
	module.last[signal.id] = signal.signal
	ret := Low
	for _, value := range module.last {
		if value == Low {
			ret = High
			break
		}
	}
	for _, addr := range module.addresses {
		result = append(result, Output{addr, ret})
	}
	return module, result
}

func (module ConjunctionModule) Successors() []string {
	return module.addresses
}


func ParsePart(line string) int {
	reg := regexp.MustCompile("^{x=(\\d+),m=(\\d+),a=(\\d+),s=(\\d+)}$")
	match := reg.FindStringSubmatch(line)
	x, _ := strconv.Atoi(match[1])
	return x
}

type Packet struct {
	from string
	to string
	signal int
}

func SolveFirst(modules map[string]Module) int {
	low, high := 0, 0
	for i:=0; i<1000; i++ {
		nlow, nhigh := Cycle(modules)
		low +=nlow
		high +=nhigh
	}
	return low*high
}

func Cycle(modules map[string]Module) (int, int) {
	queue := make([]Packet, 0);
	queue = append(queue, Packet{"button", "broadcaster", Low})
	low := 0
	high := 0
	for len(queue) > 0 {
		packet := queue[0]
		queue = queue[1:]
		if packet.signal == Low {
			low++
		} else {
			high++
		}
		if _, ok := modules[packet.to]; ok {
			input := Input{packet.from, packet.signal}
			new_mod, out := modules[packet.to].Proccess(input)
			modules[packet.to] = new_mod
			for _, output := range out {
				queue = append(queue, Packet{packet.to, output.id, output.signal})
			}
		}
	}
	return low, high
}

func SolveSecond(modules map[string]Module) int {
	return 0

}
