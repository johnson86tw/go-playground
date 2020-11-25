package main

import (
	"fmt"
)

// type UnionFind struct {
// 	parent []int
// 	n      int
// }

// func (uf *UnionFind) init(n int) {
// 	uf.n = n
// 	uf.parent = make([]int, n)
// 	for i := 0; i < n; i++ {
// 		uf.parent[i] = i
// 	}
// }

// func (uf *UnionFind) find(p int) int {
// 	root := p

// 	for root != uf.parent[root] {
// 		root = uf.parent[root]
// 	}

// 	return root
// }

// func (uf *UnionFind) union(p, q int) {
// 	rootp := uf.find(p)
// 	rootq := uf.find(q)

// 	if rootp != rootq {
// 		uf.parent[rootq] = rootp
// 	}
// }

type UnionFind struct {
	parent, rank []int
	count        int
}

// Init define
func (uf *UnionFind) Init(n int) {
	uf.count = n
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
	}
}

// Find define
func (uf *UnionFind) Find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	// compress path
	for p != uf.parent[p] {
		tmp := uf.parent[p]
		uf.parent[p] = root
		p = tmp
	}
	return root
}

// Union define
func (uf *UnionFind) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return
	}
	if uf.rank[qroot] > uf.rank[proot] {
		uf.parent[proot] = qroot
	} else {
		uf.parent[qroot] = proot
		if uf.rank[proot] == uf.rank[qroot] {
			uf.rank[proot]++
		}
	}
	uf.count--
}

func main() {
	clues := [][]int{{1, 3}, {2, 4}, {4, 6}, {5, 4}}
	unionFind(clues)
}

func unionFind(pair [][]int) {
	uf := &UnionFind{}
	uf.Init(6)
	fmt.Printf("initial UnionFind: %v\n", *uf)

	for _, c := range pair {
		uf.Union(c[0]-1, c[1]-1)
	}

	fmt.Println(*uf)

}
