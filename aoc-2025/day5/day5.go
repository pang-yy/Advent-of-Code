package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	intervals, ids, err := parseInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	intervals = mergeAndSortIntervals(intervals)

	fmt.Printf("Part 1: %d\n", partOne(intervals, ids))
	fmt.Printf("Part 2: %d\n", partTwo(intervals))
}

func partOne(intervals [][2]int, ids []int) int {
	count := 0
	for _, id := range ids {
		if binarySearch(intervals, id) {
			count += 1
		}
	}
	return count
}

func partTwo(intervals [][2]int) int {
	total := 0
	for _, interval := range intervals {
		total += (interval[1] - interval[0]) + 1
	}
	return total
}

func mergeAndSortIntervals(intervals [][2]int) [][2]int {
	sortedIntervals := slices.SortedFunc(slices.Values(intervals), func(r1, r2 [2]int) int {
		return cmp.Compare(r1[0], r2[0])
	}) // sorted by starting point, do not modify original slice
	newIntervals := [][2]int{}

	currInt := [2]int{sortedIntervals[0][0], sortedIntervals[0][1]}
	for i := 1; i < len(sortedIntervals); i++ {
		if currInt[1] < sortedIntervals[i][0] { // disjoint
			newIntervals = append(newIntervals, currInt)
			currInt = [2]int{sortedIntervals[i][0], sortedIntervals[i][1]}
		} else {
			currInt[1] = max(currInt[1], sortedIntervals[i][1])
		}
	}
	newIntervals = append(newIntervals, currInt)

	return newIntervals
}

func binarySearch(intervals [][2]int, target int) bool {
	low := 0
	high := len(intervals) - 1
	for low <= high {
		mid := (low + high) / 2
		midInt := intervals[mid]
		if target >= midInt[0] && target <= midInt[1] {
			return true
		}
		if midInt[0] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func parseInput() ([][2]int, []int, error) {
	var intervals [][2]int
	var ids []int
	var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		var line string = scanner.Text()
		if len(line) == 0 {
			break
		}
		endpoint := [2]int{}
		_, err := fmt.Sscanf(line, "%d-%d", &endpoint[0], &endpoint[1])
		if err != nil {
			return nil, nil, err
		}
		intervals = append(intervals, endpoint)
	}
	for {
		scanner.Scan()
		var line string = scanner.Text()
		if len(line) == 0 {
			break
		}
		var n int
		_, err := fmt.Sscanf(line, "%d", &n)
		if err != nil {
			return nil, nil, err
		}
		ids = append(ids, n)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return intervals, ids, nil
}
