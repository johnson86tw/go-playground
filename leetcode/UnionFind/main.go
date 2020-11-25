package main

import "fmt"

func main() {

	uf := NewUnionFind(6)
	uf.Union(3, 5)
	uf.Union(1, 4)
	uf.Union(5, 2)
	uf.Union(1, 5)

	uf.Union(3, 1)

	fmt.Printf("%v\n", *uf)
}

// UnionFind ...
type UnionFind struct {
	root, size []int
	count      int
}

// NewUnionFind ...
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{}
	uf.init(n)
	return uf
}

func (uf *UnionFind) init(n int) {
	uf.count = n
	uf.root = make([]int, n)
	uf.size = make([]int, n)

	for i := range uf.root {
		uf.root[i] = i
		uf.size[i] = 1
	}
}

// Find define
func (uf *UnionFind) Find(p int) int {
	if p > len(uf.root)-1 {
		return -1
	}

	for p != uf.root[p] {
		uf.root[p] = uf.root[uf.root[p]]
		p = uf.root[p]
	}

	return p
}

// Union define
func (uf *UnionFind) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)

	if pRoot == qRoot {
		return
	}

	if uf.size[qRoot] > uf.size[pRoot] {
		uf.root[pRoot] = qRoot
		uf.size[qRoot] += uf.size[pRoot]
	} else {
		uf.root[qRoot] = uf.root[pRoot]
		uf.size[pRoot] += uf.size[qRoot]
	}

	uf.count--
}

// TotalCount ...
func (uf *UnionFind) TotalCount() int {
	return uf.count
}
