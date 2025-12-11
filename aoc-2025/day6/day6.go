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
	fmt.Printf("Part 1: %d\n", partOne(nums, operators))
}

func partOne(nums []int, operators []bool) int {
	total := 0
	linesCount := len(nums) / len(operators)

	for i, o := range operators {
		subTotal := 0
		if o {
			subTotal = 1
			for j := range linesCount {
				subTotal *= nums[i+(len(operators)*j)]
			}
		} else {
			for j := range linesCount {
				subTotal += nums[i+(len(operators)*j)]
			}
		}
		total += subTotal
	}

	return total
}

func parseInput() ([]int, []bool, error) {
	var in []int
	var operators []bool
	var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for {
		scanner.Scan()
		var word []byte = scanner.Bytes()
		if len(word) == 0 {
			break
		}
		if word[0] == '+' || word[0] == '*' {
			operators = append(operators, word[0] == '*')
		} else {
			in = append(in, utils.BytesToInt(word))
		}
	}
	err := scanner.Err()
	if err != nil {
		return nil, nil, err
	}
	return in, operators, nil
}
