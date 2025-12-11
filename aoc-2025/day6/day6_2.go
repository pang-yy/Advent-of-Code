package main

import (
	"aoc-2025/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	nums, operators, err := parseInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Part 2: %d\n", partTwo(nums, operators))
}

func partTwo(nums [][]byte, rawOps []byte) int {
	total := 0
	operators, idxs := getColLocations(rawOps)

	for i, op := range operators {
		if i == 0 {
			total += do(nums, 0, idxs[0], op)
		} else if i == len(operators) - 1 {
			total += do(nums, idxs[len(idxs) - 1]+1, len(nums[0]), op)
		} else {
			total += do(nums, idxs[i-1]+1, idxs[i], op)
		}
	}

	return total
}

func do(nums [][]byte, start, end int, op bool) int {
	total := 0
	if op {
		total = 1
	}

	for i := (end - 1); i >= start; i-- {
		numByte := []byte{}
		for _, line := range nums {
			numByte = append(numByte, line[i])
		}
		n := utils.BytesToInt(numByte)
		if op {
			total *= n
		} else {
			total += n
		}
	}

	return total
}

func getColLocations(rawOps []byte) ([]bool, []int) {
	operators := []bool{ rawOps[0] == '*' }
	idxs := []int{}

	for i := 1; i < len(rawOps) - 1; i++ {
		if (rawOps[i] == ' ') && (rawOps[i+1] != ' ') {
			operators = append(operators, rawOps[i+1] == '*')
			idxs = append(idxs, i)
		}
	}

	return operators, idxs
}

func parseInput() ([][]byte, []byte, error) {
	var in [][]byte
	var operators []byte
	var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		var line []byte = scanner.Bytes()
		if len(line) == 0 {
			break
		}
		if line[0] == '+' || line[0] == '*' {
			operators = make([]byte, len(line))
			copy(operators, line)
		} else {
			// scanner.Bytes() does not do allocation !
			lineCopy := make([]byte, len(line))
			copy(lineCopy, line)
			in = append(in, lineCopy)
		}
	}
	err := scanner.Err()
	if err != nil {
		return nil, nil, err
	}
	return in, operators, nil
}
