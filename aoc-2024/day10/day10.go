package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var tMap [][]byte
    var err error
    if tMap, err = readInput(); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Part One: %d\n", countTrailhead(tMap, true))
    fmt.Printf("Part Two: %d\n", countTrailhead(tMap, false))
}

func countTrailhead(tMap [][]byte, isPartOne bool) int {
    var result int = 0

    var visited [][]int = make([][]int, len(tMap))
    for i := 0; i < len(tMap); i += 1 {
        visited[i] = make([]int, len(tMap[0]))
    }
    var checker int = 0

    for y, line := range(tMap) {
        for x, digit := range(line) {
            if (digit == 0) {
                // another tricks: use [][]int instead, so no need to clear visited every time
                //var visited [][]bool = make([][]bool, len(tMap))
                /*
                for i := 0; i < len(tMap); i += 1 {
                    visited[i] = make([]bool, len(tMap[0]))
                }
                */
                var count int = 0
                dfs(tMap, x, y, checker, visited, &count, isPartOne)
                result += count
            }
            checker += 1
        }
    }

    return result
}

func dfs(tMap [][]byte, x, y, checker int, visited [][]int, count *int, isPartOne bool) {

    if (x < 0 || y < 0 || x >= len(tMap[0]) || y >= len(tMap)) {
        return
    }
    if (tMap[y][x] == 9 && visited[y][x] != checker) {
        if (isPartOne) {visited[y][x] = checker}
        *count += 1
    }
    if (x > 0 && (tMap[y][x - 1] - tMap[y][x] == 1)) {
        dfs(tMap, x - 1, y, checker, visited, count, isPartOne)
    }
    if (x < len(tMap[0]) - 1 && (tMap[y][x + 1] - tMap[y][x] == 1)) {
        dfs(tMap, x + 1, y, checker, visited, count, isPartOne)
    }
    if (y > 0 && (tMap[y - 1][x] - tMap[y][x] == 1)) {
        dfs(tMap, x, y - 1, checker, visited, count, isPartOne)
    }
    if (y < len(tMap) - 1 && (tMap[y + 1][x] - tMap[y][x] == 1)) {
        dfs(tMap, x, y + 1, checker, visited, count, isPartOne)
    }
}

func readInput() ([][]byte, error) {
    var tMap [][]byte
    var err error
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        var line string = scanner.Text()
        if (len(line) == 0) {
            break
        }
        var l []byte
        for _, digit := range(line) {
            l = append(l, byte(digit - '0'))
        }
        tMap = append(tMap, l)
    }

    err = scanner.Err()
    if (err != nil) {
        return tMap, err
    }

    return tMap, nil
}
