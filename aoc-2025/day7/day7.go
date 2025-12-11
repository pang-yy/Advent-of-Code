package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputs, err := parseInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	p1, p2 := simulate(inputs)

	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
}

func simulate(inputs []string) (int, int) {
	splitCount := 0
	beamIdx := []int{ len(inputs[0]) / 2 }
	beamWeight := []int{ 1 }

	for i := 1; i < len(inputs); i++ {
		nextBeamIdx := []int{}
		nextBeamWeight := []int{}
		for j, idx := range beamIdx {
			if inputs[i][idx] == '^' {
				if (len(nextBeamIdx) == 0) || (nextBeamIdx[len(nextBeamIdx) - 1] != idx - 1) {
					nextBeamIdx = append(nextBeamIdx, idx - 1)
					nextBeamWeight = append(nextBeamWeight, beamWeight[j])
				} else {
					nextBeamWeight[len(nextBeamWeight) - 1] += beamWeight[j]
				}
				nextBeamIdx = append(nextBeamIdx, idx + 1)
				nextBeamWeight = append(nextBeamWeight, beamWeight[j])
				splitCount += 1
			} else if (len(nextBeamIdx) == 0) || (nextBeamIdx[len(nextBeamIdx) - 1] != idx) {
				nextBeamIdx = append(nextBeamIdx, idx)
				nextBeamWeight = append(nextBeamWeight, beamWeight[j])
			} else {
				nextBeamWeight[len(nextBeamWeight) - 1] += beamWeight[j]
			}
		}
		beamIdx = nextBeamIdx
		beamWeight = nextBeamWeight
	}

	totalWeight := 0
	for _, w := range beamWeight {
		totalWeight += w
	}

	return splitCount, totalWeight
}

func parseInput() ([]string, error) {
	var in []string
	var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		var line string = scanner.Text()
		if len(line) == 0 {
			break
		}
		in = append(in, line)
	}
	return in, nil
}
