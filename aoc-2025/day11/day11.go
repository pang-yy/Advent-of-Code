package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Devices map[string][]string

type Pair struct {
	name  string
	ready bool
}

func main() {
	devices, err := parseInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Started")
	fmt.Printf("Part 1: %d\n", partOne(devices))
	fmt.Printf("Part 2: %d\n", partTwo(devices))
}

func partOne(devices Devices) int {
	// return partOneDFS(devices, "you", "out")
	return partOneDP(devices, "you", "out")
}

func partTwo(devices Devices) int {
	b := partOneDP(devices, "fft", "dac")
	if b != 0 {
		a := partOneDP(devices, "svr", "fft")
		c := partOneDP(devices, "dac", "out")
		return a * b * c
	}
	d := partOneDP(devices, "svr", "dac")
	e := partOneDP(devices, "dac", "fft")
	f := partOneDP(devices, "fft", "out")
	return d * e * f
}

// assume acyclic
func partOneDFS(devices Devices, start, end string) int {
	stack := []Pair{
		{start, false},
	}
	visited := map[string]int{
		end: 1,
	}

	for len(stack) != 0 {
		curr := stack[len(stack)-1]

		if curr.ready {
			stack = stack[:len(stack)-1]

			path_count := 0
			for _, o := range devices[curr.name] {
				path_count += visited[o]
			}
			visited[curr.name] = path_count
		} else {
			if _, exists := visited[curr.name]; exists {
				stack = stack[:len(stack)-1]
				continue
			}

			visited[curr.name] = -1 // visiting this path
			stack[len(stack)-1].ready = true

			for _, o := range devices[curr.name] {
				if val, exists := visited[o]; !exists {
					stack = append(stack, Pair{o, false})
				} else if val == -1 { // child/neighbour not done, but is in visiting state => cycle back
					panic("cycle detected")
				}
			}
		}
	}

	return visited[start]
}

// assume acyclic
func partOneDP(devices Devices, start, end string) int {
	indegree := make(map[string]int, len(devices))
	queue := []string{}

	pathCounts := make(map[string]int, len(devices))

	// the "1" spreads from "start" in all direction, and converges at the "end" to give total
	pathCounts[start] = 1

	for _, outs := range devices {
		for _, out := range outs {
			indegree[out] += 1
		}
	}

	for name := range devices {
		if indegree[name] == 0 {
			queue = append(queue, name)
		}
	}

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, out := range devices[curr] {
			indegree[out] -= 1
			if indegree[out] == 0 {
				queue = append(queue, out)
			}
			pathCounts[out] += pathCounts[curr]
		}
	}

	return pathCounts[end]
}

func parseInput() (Devices, error) {
	in := make(map[string][]string)
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		sArr := strings.Split(line, ": ")
		deviceName := sArr[0]
		outs := strings.Split(sArr[1], " ")
		in[deviceName] = outs
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}
