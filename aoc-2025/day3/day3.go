package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	banks, err := parseInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Part 1: %d\n", findLargestDigits(banks, 2))
	fmt.Printf("Part 2: %d\n", findLargestDigits(banks, 12))
}

func findLargestDigits(banks []string, size int) int {
	total := 0
	for _, bank := range banks {
		batteries_stack := make([]byte, size) // monotonic decreasing stack
		stack_ptr := 0
		for i, c := range bank {
			for (stack_ptr > 0) &&
					(len(bank) - i > (size - stack_ptr)) &&
					(batteries_stack[stack_ptr - 1] < byte(c - '0')) {
				stack_ptr -= 1
			}
			if stack_ptr < size {
				batteries_stack[stack_ptr] = byte(c - '0')
				stack_ptr += 1
			}
		}
		mul := 1
		t := 0
		for i := (size - 1); i >= 0; i-- {
			t += (int(batteries_stack[i]) * mul)
			mul *= 10
		}
		total += t
	}
	return total
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
