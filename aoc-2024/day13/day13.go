package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
    "errors"
)

type Pair struct {
    first int
    second int
}

type Machine struct {
    buttonA Pair
    buttonB Pair
    prize Pair
}

func main() {
    var machines []Machine
    var err error
    if machines, err = readInput(); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Part One: %d\n", toMat(machines, 0))
    fmt.Printf("Part Two: %d\n", toMat(machines, 10000000000000))
}

func toMat(machines []Machine, displace int) int {
    var total = 0
    const costPerA int = 3
    const costPerB int = 1

    var x [2]int
    var err error
    for _, machine := range(machines) {
        if x, err = sol2([2][2]int{
            {machine.buttonA.first, machine.buttonB.first}, 
            {machine.buttonA.second, machine.buttonB.second},
        }, [2]int{machine.prize.first + displace, machine.prize.second + displace}); err == nil {
            total += ((x[0] * costPerA) + (x[1] * costPerB))
        }
    }

    return total
}

func sol2(A [2][2]int, b [2]int) ([2]int, error) { // Ax = b, A: 2x2, b: 2x1
    var A_inv [2][2]int
    var det_A_inv int = 0
    var x [2]int
    var err error
    if A_inv, det_A_inv, err = inverse2(A); err != nil {
        return [2]int{}, err
    }
    x[0] = ((A_inv[0][0] * b[0]) + (A_inv[0][1] * b[1]))
    x[1] = ((A_inv[1][0] * b[0]) + (A_inv[1][1] * b[1]))
    if (x[0] % det_A_inv != 0 || x[1] % det_A_inv != 0) {
        return [2]int{}, errors.New("not an integer")
    }
    x[0] /= det_A_inv
    x[1] /= det_A_inv
    return x, nil
}

func inverse2(matrix [2][2]int) ([2][2]int, int, error) {
    var det int = determinant2(matrix)
    if (det == 0) {
        return [2][2]int{}, 0, errors.New("no inverse")
    }
    matrix[0][0], matrix[1][1] = matrix[1][1], matrix[0][0]
    matrix[0][1], matrix[1][0] = -matrix[0][1], -matrix[1][0]
    return matrix, det, nil
}

func determinant2(matrix [2][2]int) int {
    return (matrix[0][0] * matrix[1][1]) - (matrix[0][1] * matrix[1][0])
}

func readInput() ([]Machine, error) {
    var machines []Machine
    var err error
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    var temp [3]Pair
    var t1 int
    var t2 int
    var line []string
    var re regexp.Regexp = *regexp.MustCompile(`X[+=]|Y[+=]`)
    for {
        for i := 0; i < 3; i += 1 {
            scanner.Scan()
            line = strings.Split(scanner.Text(), ", ")
            if _, err = fmt.Sscanf(re.Split(line[0], -1)[1], "%d", &t1); err != nil {
                return []Machine{}, err
            }
            if _, err = fmt.Sscanf(re.Split(line[1], -1)[1], "%d", &t2); err != nil {
                return []Machine{}, err
            }
            temp[i] = Pair{first: t1, second: t2}
        }
        machines = append(machines, Machine{buttonA: temp[0], buttonB: temp[1], prize: temp[2]})
        if (!scanner.Scan()) { // empty line after 3 scans 
            break
        }
    }

    err = scanner.Err()
    if (err != nil) {
        return machines, err
    }

    return machines, nil
}
