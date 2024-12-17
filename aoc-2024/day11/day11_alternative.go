package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    var stones []int
    var err error
    if stones, err = readInput(); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Part One: %d\n", countStones(stones, 25))
    fmt.Printf("Part Two: %d\n", countStones(stones, 75))
    fmt.Printf("Part X: %d\n", countStones(stones, 300))
}

func countStones(stones []int, blinkTimes int) int {
    var cache map[int]int = make(map[int]int)
    for _, val := range(stones) {
        cache[val] = 1
    }

    for i := 0; i < blinkTimes; i += 1 {
        cache = blink(cache)
    }

    var total int = 0
    for _, count := range(cache) {
        total += count
    }
    return total
}

func blink(inital map[int]int) map[int]int {
    var updatedCache map[int]int = make(map[int]int)
    for key, count := range(inital) {
        if (key == 0) {
            updatedCache[1] += count
        } else if numberOfDigits := countDigits(key); (numberOfDigits % 2 == 0) {
            var left, right int = splitNumber(key, numberOfDigits)
            updatedCache[left] += count
            updatedCache[right] += count
        } else {
            updatedCache[key * 2024] += count
        }
    }
    return updatedCache
}

func splitNumber(val, numberOfDigits int) (int, int) {
    var left int = val
    var right int = 0
    var multipleOfTen int = 1

    for i := 0; i < numberOfDigits / 2; i += 1 {
        right = (left % 10 * multipleOfTen) + right
        left /= 10
        multipleOfTen *= 10
    }

    return left, right
}

func countDigits(val int) int {
    var i int = 0
    for val > 0 {
        i += 1
        val /= 10
    }
    return i
}

func readInput() ([]int, error) {
    var stones []int
    var err error
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    scanner.Scan()
    var line string = scanner.Text()

    var t int
    for _, sv := range(strings.Split(line, " ")) {
        if _, err = fmt.Sscanf(sv, "%d", &t); err != nil {
            return []int{}, err
        }
        stones = append(stones, t)
    }

    err = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return stones, err
    }

    return stones, nil
}
