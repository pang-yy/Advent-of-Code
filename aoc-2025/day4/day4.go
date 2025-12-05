package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, err := parseInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, count := findAccess(grid)
	fmt.Printf("Part 1: %d\n", count)
	fmt.Printf("Part 2: %d\n", findFinalAccess(grid))
}

func findAccess(grid [][]bool) ([][]bool, int) {
	total := 0
	dir := [8][2]int{ // {x, y}
		{-1,-1},{0,-1},{1,-1},{-1,0},
		{1,0},{-1,1},{0,1},{1,1},
	}
	newGrid := make([][]bool, len(grid))
	for i := range len(grid) {
		newGrid[i] = make([]bool, len(grid[0]))
		for j := range len(grid[0]) {
			newGrid[i][j] = grid[i][j]
		}
	}

	for rIdx, row := range grid {
		for cIdx, col := range row {
			if !col {
				continue
			}
			count := 0
			for _, d := range dir {
				dx := cIdx+d[0]
				dy := rIdx+d[1]
				if (dx >= 0 && dx < len(row) && dy >= 0 && dy < len(grid) && grid[dy][dx]) {
					count += 1
				}
			}
			if count < 4 {
				total += 1
				newGrid[rIdx][cIdx] = false
			}
		}
	}

	return newGrid, total
}

func findFinalAccess(grid [][]bool) int {
	total := 0
	count := 0
	for {
		grid, count = findAccess(grid)
		if count == 0 {
			break
		}
		total += count
	}
	return total
}

func parseInput() ([][]bool, error) {
	var in [][]bool
	var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		var line string = scanner.Text()
		if len(line) == 0 {
			break
		}
		row := make([]bool, len(line))
		for i, c := range line {
			row[i] = (c == '@')
		}
		in = append(in, row)
	}
	return in, nil
}
