package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Pair struct {
	x int
	y int
}

func main() {
    var rawInputs []string
    var inputs [][2]Pair
    var err error
    if rawInputs, err = readInput(); err == nil {
        if inputs, err = parseInput(rawInputs); err != nil {
            fmt.Println(err)
            return
        }
    } else {
        fmt.Println(err)
        return
    }

	width := 101
	height := 103
	factor := calculatePartOne(inputs, 100, width, height)
    fmt.Printf("Part One: %d\n", factor)

	runSimulation(inputs, 10000, width, height)
}

// Part One

func calculatePartOne(robots [][2]Pair, time, width, height int) int {
	finalPos := calculatePos(robots, time, width, height)
	quadrants := calculateQuadrant(finalPos, width, height)
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func calculatePos(robots [][2]Pair, time, width, height int) []Pair {
	var pos []Pair = []Pair{}
	for _, r := range(robots) {
		nx := (r[0].x + (r[1].x * time)) % width
		ny := (r[0].y + (r[1].y * time)) % height
		if (nx < 0) {
			nx += width
		}
		if (ny < 0) {
			ny += height
		}
		pos = append(pos, Pair{x: nx, y: ny})
	}
	return pos
}

func calculateQuadrant(pos []Pair, width, height int) [4]int {
	var quadrants [4]int = [4]int{0, 0, 0, 0}
	for _, p := range(pos) {
		if ((p.x >= 0 && p.x <= (width / 2 - 1)) &&
            (p.y >= 0 && p.y <= (height / 2 - 1))) {
			quadrants[0] += 1
        } else if ((p.x >= (width / 2 + 1) && (p.x <= width - 1)) &&
                    (p.y >= 0 && p.y <= (height / 2 - 1))) {
            quadrants[1] += 1
        } else if ((p.x >= 0 && p.x <= (width / 2 - 1)) &&
                    (p.y >= (height / 2 + 1) && p.y <= height - 1)) {
            quadrants[2] += 1
        } else if ((p.x >= (width / 2 + 1) && (p.x <= width - 1)) &&
                    (p.y >= (height / 2 + 1) && (p.y <= height - 1))) {
            quadrants[3] += 1
        }
	}
	return quadrants
}

// Part Two

func runSimulation(robots [][2]Pair, time, width, height int) {
	displayFloor(robots, width, height)
	fmt.Println(strings.Repeat("=", 50))
	for i := range(time) {
		robots = nextState(robots, width, height)
		if ((i + 1 - height) % width == 0) {
			displayFloor(robots, width, height)
			fmt.Printf("Count: %d\n", i + 1)
			fmt.Println(strings.Repeat("=", 50))
		}
	}
}

func nextState(robots [][2]Pair, width, height int) [][2]Pair {
	for i, r := range(robots) {
		r[0].x = (r[0].x + r[1].x) % width
		r[0].y = (r[0].y + r[1].y) % height
		if (r[0].x < 0) {
			r[0].x += width
		}
		if (r[0].y < 0) {
			r[0].y += height
		}
		robots[i] = r
	}
	return robots
}

func displayFloor(robots [][2]Pair, width, height int) {
	posMap := make(map[Pair]int, width * height)
	for _, r := range(robots) {
		pos := r[0]
		_, ok := posMap[pos]
		if ok {
			posMap[pos] += 1
		} else {
			posMap[pos] = 1
		}
	}
	for y := range(height) {
		for x:= range(width) {
			count, ok := posMap[Pair{x: x, y: y}]
			if (ok && count > 0) {
				fmt.Print(count)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func parseInput(rawInputs []string) ([][2]Pair, error) {
    var parseResult [][2]Pair = [][2]Pair{}
    //var parseResult [][4]int = [][4]int{}
    var err error
    for _, line := range(rawInputs) {
        var in [2]Pair = [2]Pair{Pair{}, Pair{}}
        //var in [4]int = [4]int{}
        var temp []string = strings.Split(line, " ")

		for i, pv := range(temp) {
			var nums []string = strings.Split(strings.Split(pv, "=")[1], ",")
			_, err = fmt.Sscanf(nums[0], "%d", &in[i].x)
			if (err != nil) {
				return [][2]Pair{}, err
				//return [][4]int{}, err
			}

			_, err = fmt.Sscanf(nums[1], "%d", &in[i].y)
			if (err != nil) {
				//return [][4]int{}, err
				return [][2]Pair{}, err
			}
		}
        parseResult = append(parseResult, in)
    }
    return parseResult, nil
}

func readInput() ([]string, error) {
    var input []string
    var err error
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        var line string = scanner.Text()

        if (len(line) == 0) {
            break
        }

        input = append(input, line)
    }

    err = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return input, err
    }

    return input, nil
}
