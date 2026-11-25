package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Region struct {
	width, height int
	quantity      [6]int
}

type Shape struct {
	raw        [3]string
	blockCount int
}

func main() {
	regions, shapes, err := parseInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Started")
	fmt.Printf("Part 1: %d\n", partOne(regions, shapes))
}

func partOne(regions []Region, shapes [6]Shape) int {
	needToCheck := []Region{}
	count := 0

	defCannotFit := func(r Region) bool {
		total := 0
		for i, val := range r.quantity {
			total += val * shapes[i].blockCount
		}
		return total > (r.width * r.height)
	}

	defCanFit := func(r Region) bool {
		w := r.width / 3
		h := r.height / 3
		return (w * h) >= (r.quantity[0] + r.quantity[1] + r.quantity[2] + r.quantity[3] + r.quantity[4] + r.quantity[5])
	}
	// canFit := func (r Region) bool {
	// 	return false
	// }

	for _, r := range regions {
		if defCannotFit(r) {
			continue
		}
		if defCanFit(r) {
			count += 1
		} else {
			needToCheck = append(needToCheck, r)
		}
	}

	for _, r := range needToCheck {
		fmt.Println(r.width, r.height, r.quantity)
		// ...
	}

	return count
}

func parseInput() ([]Region, [6]Shape, error) {
	shapes := [6]Shape{}
	regions := []Region{}
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	for i := range 6 {
		shape := Shape{
			[3]string{},
			0,
		}

		scanner.Scan() // index
		for j := range 3 {
			scanner.Scan()
			line := scanner.Text()
			shape.blockCount += strings.Count(line, "#")
			shape.raw[j] = line
		}
		scanner.Scan() // newline
		shapes[i] = shape
	}

	for scanner.Scan() {
		region := Region{}
		line := scanner.Text()
		inputs := strings.Split(line, ": ")

		dimStr := strings.Split(inputs[0], "x")
		w, _ := strconv.Atoi(dimStr[0])
		h, _ := strconv.Atoi(dimStr[1])
		region.width = w
		region.height = h

		quantStr := strings.Split(inputs[1], " ")
		for i, qs := range quantStr {
			n, _ := strconv.Atoi(qs)
			region.quantity[i] = n
		}

		regions = append(regions, region)
	}

	if err := scanner.Err(); err != nil {
		return nil, [6]Shape{}, err
	}

	return regions, shapes, nil
}
