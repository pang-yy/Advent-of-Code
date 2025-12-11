package main

import (
	"aoc-2025/utils"
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
)

type Position struct { x,y,z int }
type Metadata struct { d, idx1, idx2 int }

func main() {
	boxes, err := parseInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Part 1: %d\n", partOne(boxes, 1000))
	fmt.Printf("Part 2: %d\n", partTwo(boxes))
}

func partOne(boxes []Position, iteration int) int {
	metaList := []Metadata{}
	indexMap := utils.NewUnionFind(len(boxes))

	for i := range len(boxes) {
		for j := i+1; j < len(boxes); j += 1 {
			d := distanceSquare(boxes[i], boxes[j])
			metaList = append(metaList, Metadata{ d: d, idx1: i, idx2: j })
		}
	}

	// TODO: replace with minHeap?
	slices.SortFunc(metaList, func(m1, m2 Metadata) int {
		return cmp.Compare(m1.d, m2.d)
	})

	for i := range iteration {
		comp := metaList[i]
		if !indexMap.Find(comp.idx1, comp.idx2) {
			indexMap.Union(comp.idx1, comp.idx2)
		}
	}

	return indexMap.MultiplySizeOfRoot()
}

func partTwo(boxes []Position) int {
	metaList := []Metadata{}
	indexMap := utils.NewUnionFind(len(boxes))

	for i := range len(boxes) {
		for j := i+1; j < len(boxes); j += 1 {
			d := distanceSquare(boxes[i], boxes[j])
			metaList = append(metaList, Metadata{ d: d, idx1: i, idx2: j })
		}
	}

	// TODO: replace with minHeap?
	slices.SortFunc(metaList, func(m1, m2 Metadata) int {
		return cmp.Compare(m1.d, m2.d)
	})

	lastIdx1 := 0
	lastIdx2 := 0
	for i := 0; !indexMap.IsFullyConnected(); i++ {
		comp := metaList[i]
		if !indexMap.Find(comp.idx1, comp.idx2) {
			indexMap.Union(comp.idx1, comp.idx2)
			lastIdx1 = comp.idx1
			lastIdx2 = comp.idx2
		}
	}

	return boxes[lastIdx1].x * boxes[lastIdx2].x
}

func distanceSquare(box1, box2 Position) int {
	const lol = 100
	dxS := (box1.x - box2.x) * (box1.x - box2.x)
	dyS := (box1.y - box2.y) * (box1.y - box2.y)
	dzS := (box1.z - box2.z) * (box1.z - box2.z)
	return dxS + dyS + dzS
}

func parseInput() ([]Position, error) {
	var in []Position
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		var raw []byte = scanner.Bytes()
		if len(raw) == 0 {
			break
		}
		rawCopy := make([]byte, len(raw))
		copy(rawCopy, raw)
		nums := bytes.Split(rawCopy, []byte{','})
		box := Position{
			x: utils.BytesToInt(nums[0]),
			y: utils.BytesToInt(nums[1]),
			z: utils.BytesToInt(nums[2]),
		}
		in = append(in, box)
	}
	return in, nil
}
