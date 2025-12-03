package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	ranges, err := parseInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Printf("Part 1: %d\n", cal1(ranges))
	//fmt.Printf("Part 2: %d\n", cal2(ranges))
	fmt.Printf("Part 1: %d\n", cal_conc(ranges, isSym))
	fmt.Printf("Part 2: %d\n", cal_conc(ranges, isBruteForceRepeat))

}

func intLen(n int) int {
	if n == 0 {
		return 1
	}
	l := 0
	for n > 0 {
		n /= 10
		l += 1
	}
	return l
}

func isSym(n, l int) bool {
	if l%2 != 0 {
		return false
	}
	var left, right int
	var mul int = 1
	left = n
	for range l/2 {
		right += ((left % 10) * mul)
		left /= 10
		mul *= 10
	}
	return left == right
}

func isBruteForceRepeat(n, l int) bool {
	for i := 1; i <= l/2; i++ {
		if l % i == 0 { // possible
			allSame := true
			num := n
			ref := 0
			mul := 1
			for range i {
				ref += ((num % 10) * mul)
				num /= 10
				mul *= 10
			}
			for num > 0 {
				mul = 1
				refAgainst := 0
				for range i {
					refAgainst += ((num % 10) * mul)
					num /= 10
					mul *= 10
				}
				if ref != refAgainst {
					allSame = false
					break
				}
			}
			if allSame {
				return true
			}
		}
	}
	return false
}

func cal1(ranges [][]int) int {
	total := 0
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			if l := intLen(i); isSym(i, l) {
				total += i
			}
		}
	}
	return total
}

func cal2(ranges [][]int) int {
	total := 0
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			if l := intLen(i); isBruteForceRepeat(i, l) {
				total += i
			}
		}
	}
	return total
}

func cal_conc(ranges [][]int, prop func(int,int)bool) int {
	total := 0
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	var resChan chan int = make(chan int, 100)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for t := range resChan {
			total += t
		}
	}()

	for _, r := range ranges {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			sub_total := 0
			for i := r[0]; i <= r[1]; i++ {
				if l := intLen(i); prop(i, l) {
					sub_total += i
				}
			}
			resChan <- sub_total
		}()
	}

	wg2.Wait()
	close(resChan)
	wg.Wait()
	return total
}

func parseInput() ([][]int, error) {
	var in [][]int
	var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var line string = scanner.Text()
	for r := range strings.SplitSeq(line, ",") {
		l := strings.Split(r, "-")

		var n1, n2 int
		_, err := fmt.Sscanf(l[0], "%d", &n1)
		if err != nil {
			return nil, err
		}
		_, err = fmt.Sscanf(l[1], "%d", &n2)
		if err != nil {
			return nil, err
		}

		in = append(in, []int{n1, n2})
	}
	return in, nil
}
