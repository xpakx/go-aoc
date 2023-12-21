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
	fmt.Println("Advent of Code, day 20")
	fmt.Println("=====================")
	modules := Parse("input.txt")
	fmt.Println(modules)
	fmt.Print("*  ")
	fmt.Println(SolveFirst(modules))
	modules = Parse("input.txt")
	result := SolveSecond(modules)
	fmt.Print("** ")
	fmt.Println(result)
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
		nlow, nhigh, _ := Cycle(modules, "broadcaster", "rx")
		low +=nlow
		high +=nhigh
	}
	return low*high
}

func Cycle(modules map[string]Module, start string, end string) (int, int, []Packet) {
	queue := make([]Packet, 0);
	queue = append(queue, Packet{"button", start, Low})
	low := 0
	high := 0
	endNode := make([]Packet, 0)
	for len(queue) > 0 {
		packet := queue[0]
		queue = queue[1:]
		if packet.to == end {
			endNode = append(endNode, packet)
		}
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
	return low, high, endNode
}

func SolveSecond(modules map[string]Module) int {
	comps := tarjan(modules)
	fmt.Println(comps)
	components := make([]ConnectedModule, 0)
	for _, component := range comps {
		if len(component) > 1 {
			inputs, outputs := InspectComponent(modules, component)
			fmt.Println()
			fmt.Println(component, inputs, outputs)
			modulesNew := make(map[string]Module)

			for key, value := range modules {
				modulesNew[key] = value
			}
			node := ConnectedModule{outputs, inputs, component, make(map[ConnectedKey]ConnectedValue), modulesNew, -1, 1}
			components = append(components, node)
		}
	}
	

	cycleLengths := make([]int, 0)

	for _, c := range components{
		i, j, value := c.FindCycle()
		length := j-i
		fmt.Println("Components ending at", c.outputs, "has cycle", i, "-", j, "of length", length)
		fmt.Println("Low signals in cycle:", value)
		cycleLengths = append(cycleLengths, length)
	}

	// there is only single output for each strongly connected component, each send repeteadly Low n steps before cycle end, and the 
	// cycle start at n step from start. also each strongly component is connected to single inverter, and all of those inverters 
	// connect to conjunction module, that is connected to rx, so it should be possible to calculate result with lcm 
	// Actually, all periods are prime, so lcm isn't even necessary 

	return LCM(cycleLengths)
}

func GCD(a int, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

func LCM(nums []int) int {
      result := nums[0] * nums[1] / GCD(nums[0], nums[1])
      for i := 2; i < len(nums); i++ {
              result = LCM([]int{result, nums[i]})
      }
      return result
}

func tarjan(modules map[string]Module) [][]string {
	result := make([][]string, 0)
	index := 0
	indices := make(map[string]int, 0)
	lowlinks := make(map[string]int, 0)
	onStack := make(map[string]bool, 0)
	stack := make([]string, 0);
	for key := range modules {
		if _, ok := indices[key]; !ok {
			i, stck, component := strongconnect(modules, stack, indices, index, lowlinks, onStack, result, key)
			index = i
			stack = stck
			result = component
		}
	}
	return result 
}

func strongconnect(
	modules map[string]Module,
	stack []string,
	indices map[string]int,
	index int,
	lowlinks map[string]int,
	onStack map[string]bool,
	components [][]string,
	key string) (int, []string, [][]string) {
		component := make([]string, 0)
		indices[key] = index
		lowlinks[key] = index
		index = index + 1
		stack = append(stack, key)
		onStack[key] = true
		if _, ok := modules[key]; ok {
			for _, w := range modules[key].Successors() {
				if _, ok := indices[w]; !ok {
					i, stck, cmp :=  strongconnect(modules, stack, indices, index, lowlinks, onStack, components, w)
					components = cmp
					stack = stck
					index = i 
					lowlinks[key] = Min(lowlinks[key], lowlinks[w])
				} else if value, ok := onStack[w]; ok && value {
					lowlinks[key] = Min(lowlinks[key], indices[w])
				}
			}
		}

		if lowlinks[key] == indices[key] {
			for true {
				w := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[w] = false
				component = append(component, w)
				if w == key {
					break
				}
			}
			components = append(components, component)
		}

		return index, stack, components
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func InspectComponent(modules map[string]Module, component []string) ([]string, []string) {
	inputs := make([]string, 0)
	outputs := make([]string, 0)
	for key, node := range modules {
		inComponent := false
		for _, a := range component {
			if key == a {
				inComponent = true
				break
			}
		}
		if inComponent {
			continue
		}
		for _, foreign := range node.Successors() {
			for _, inside := range component {
				if inside == foreign {
					inputs = append(inputs, foreign)
				}
			}
		}
	}
	for _, node := range component {
		for _, n := range modules[node].Successors() {
			insider := false
			for _, inside := range component {
				if inside == n {
					insider = true
					break
				}
			}
			
			if !insider {
				outputs = append(outputs, n)
			}

		}
	}

	return inputs, outputs
}

type ConnectedModule struct {
	outputs []string
	inputs []string
	nodes []string
	stateMap map[ConnectedKey]ConnectedValue
	originalModules map[string]Module
	state int
	steps int
}

type ConnectedKey struct {
	state int
	signal Input
}

type ConnectedValue struct {
	state int
	outputs []int
	first int
}

func (module ConnectedModule) Proccess(signal Input) (Module, []Output) {
	result := make([]Output, 0)
	hash := module.GetStateHash()
	output := make([]int, 0)
	if res, ok := module.stateMap[ConnectedKey{hash, signal}]; ok {
		output = res.outputs
		module.state = res.state
	} else {
		// TODO
		_, _, packets := Cycle(module.originalModules, module.inputs[0], module.outputs[0])
		newHash := module.GetStateHash()
		newSignal := -1
		if len(packets) > 0 {
			newSignal = packets[0].signal
			output = append(output, newSignal)
		}
		module.stateMap[ConnectedKey{hash, signal}] = ConnectedValue{newHash, output, module.steps}
	}
	for _, out := range output {
		for _, addr := range module.outputs {
			result = append(result, Output{addr, out})
		}
	}
	module.steps += 1
	return module, result
}

type ConnectedResult struct {
	values []int
	pos int
}

func (module ConnectedModule) FindCycle() (int, int, []ConnectedResult) {
	for module.state < 0 {
		a, _ := module.Proccess(Input{"broadcast", Low})
		if md, ok := a.(ConnectedModule); ok {
			module = md
		}
	}
	value := module.stateMap[ConnectedKey{module.state, Input{"broadcast", Low}}]
	output := make([]ConnectedResult, 0)
	for _, v := range module.stateMap {
		if len(v.outputs) > 0 {
			for _, a := range v.outputs {
				if a == 0 {
					output = append(output, ConnectedResult{v.outputs, v.first})

				}
			}
		}
	}
	return value.first, module.steps, output
}

func (module ConnectedModule) Successors() []string {
	return module.outputs
}

func (module ConnectedModule) GetStateHash() int {
	if module.state != -1 {
		return module.state
	}
	hash := 1
	for _, node := range module.nodes {
			hash = hash << 1
		if m, ok := module.originalModules[node].(FlipFlopModule); ok {
			if m.on {
				hash++
			}
		} else if m, ok := module.originalModules[node].(ConjunctionModule); ok {
			for _, a := range m.addresses {
				hash  = hash << 1
				if m.last[a] == Low {
					hash++
				}
			}

		}
	}
	return hash
}
