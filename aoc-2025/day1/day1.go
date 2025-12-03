package main

import (
	"bufio"
	"fmt"
	"os"
)

type Rotation struct {
	dir bool
	distance int
}

func main() {
	rotations, err := parseInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Part 1: %d\n", cal1(rotations))
	fmt.Printf("Part 2: %d\n", cal2(rotations))
}

func cal1(rotations []Rotation) int {
	var pos int = 50
	var ans int = 0
	for _, rot := range rotations {
		if !rot.dir { // left
			pos -= rot.distance
		} else { // right
			pos += rot.distance
		}
		pos %= 100
		if pos == 0 {
			ans += 1
		}
	}
	return ans
}

func cal2(rotations []Rotation) int {
	var pos int = 50
	var ans int = 0
	for _, rot := range rotations {
		rd := rot.distance % 100
		ans += rot.distance / 100 // every full cycle will has a pass
		prevPos := pos
		if !rot.dir { // left
			pos -= rd
			if pos < 0 {
				if prevPos != 0 { // increment only if not start at 0
					ans += 1
				}
				pos += 100
			} else if pos == 0 { // initial position == distance to the left
				ans += 1
			}
		} else { // right
			pos += rd
			if pos > 99 {
				// increment only if not start at 0
				if prevPos != 0 {
					ans += 1
				}
				pos -= 100
			}/* else if pos == 0 { // initial position + rd == 100 (but this case actually covered by condition if clause)
				ans += 1
			}*/
		}
	}
	return ans
}

func parseInput() ([]Rotation, error) {
	var in []Rotation
	var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		var line string = scanner.Text()
		if len(line) == 0 {
			break
		}
		dir := line[0] == 'R'
		var num int
		_, err := fmt.Sscanf(line[1:], "%d", &num)
		if err != nil {
			return nil, err
		}
		in = append(in, Rotation{ dir: dir, distance: num })
	}
	return in, nil
}
