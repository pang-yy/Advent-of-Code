package utils

// https://hackmd.io/@pangyy/cs2040c#Weighted-Union

type UnionFind struct {
	parents []int
	sizes   []int
	isFullyConnected bool
}

func NewUnionFind(length int) *UnionFind {
	uf := UnionFind{
		parents: make([]int, length),
		sizes:   make([]int, length),
	}

	for i := range length {
		uf.parents[i] = i
		uf.sizes[i] = 1
	}

	return &uf
}

func (uf *UnionFind) FindRoot(p int) int {
	root := p
	for uf.parents[root] != root {
		uf.parents[root] = uf.parents[uf.parents[root]]
		root = uf.parents[root]
	}
	return root
}

// Are p and q in the same set
func (uf *UnionFind) Find(p, q int) bool {
	for uf.parents[p] != p {
		p = uf.parents[p]
	}
	for uf.parents[q] != q {
		q = uf.parents[q]
	}
	return p == q

	// If want to use path compression
	//return uf.FindRoot(p) == uf.FindRoot(q)
}

// Replace sets containing p and q with their union
func (uf *UnionFind) Union(p, q int) {
	for uf.parents[p] != p {
		p = uf.parents[p]
	}
	for uf.parents[q] != q {
		q = uf.parents[q]
	}

	// Weighted Union
	// Attach the smaller tree to the root of the larger tree
	if uf.sizes[p] > uf.sizes[q] {
		uf.parents[q] = p
		uf.sizes[p] += uf.sizes[q]
		if uf.sizes[p] == len(uf.parents) {
			uf.isFullyConnected = true
		}
	} else {
		uf.parents[p] = q
		uf.sizes[q] += uf.sizes[p]
		if uf.sizes[q] == len(uf.parents) {
			uf.isFullyConnected = true
		}
	}
}

func (uf *UnionFind) IsFullyConnected() bool {
	return uf.isFullyConnected
}

// For AoC 2025 Day 8
func (uf *UnionFind) MultiplySizeOfRoot() int {
	top1 := 0
	top2 := 0
	top3 := 0
	for i, par := range uf.parents {
		if par == i {
			if uf.sizes[i] > top1 {
				top3 = top2
				top2 = top1
				top1 = uf.sizes[i]
			} else if uf.sizes[i] > top2 {
				top3 = top2
				top2 = uf.sizes[i]
			} else if uf.sizes[i] > top3 {
				top3 = uf.sizes[i]
			}
		}
	}

	return top1 * top2 * top3
}
