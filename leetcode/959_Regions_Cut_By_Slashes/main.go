package main

import "fmt"

func main() {
	fmt.Println(regionsBySlashes([]string{"/ ", " /"}))

}

func regionsBySlashes(grid []string) int {
	n := len(grid[0])
	boxNum := n * n
	chipNum := boxNum * 4

	uf := NewUnionFind(chipNum)

	for i := range grid {
		for j := range grid[i] {
			t := n*4*i + 4*j

			switch grid[i][j] {
			case ' ':
				uf.Union(t, t+1)
				uf.Union(t+1, t+2)
				uf.Union(t+2, t+3)
			case '/':
				uf.Union(t, t+3)
				uf.Union(t+1, t+2)
			case '\\':
				uf.Union(t, t+1)
				uf.Union(t+2, t+3)
			}

			// must have bottom box
			if i < n-1 {
				uf.Union(t+2, n*4*(i+1)+4*j)
			}

			// must have right box
			if j < n-1 {
				uf.Union(t+1, n*4*i+4*(j+1)+3)
			}

		}
	}

	return uf.TotalCount()
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
