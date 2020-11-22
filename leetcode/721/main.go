package main

import (
	"fmt"
	"sort"
)

// UnionFind ...
type UnionFind struct {
	parent, rank []int // 指標、團體大小
	count        int   // 團體數
}

// Init ...
func (uf *UnionFind) Init(n int) {
	uf.count = n

	uf.parent = make([]int, n)
	uf.rank = make([]int, n)

	for i := range uf.parent {
		uf.parent[i] = i
	}
}

// Find ...
func (uf *UnionFind) Find(p int) int {
	root := p

	for root != uf.parent[root] {
		root = uf.parent[root]
	}

	// path compression
	// for p != uf.parent[p] {
	// 	tmp := uf.parent[p]
	// 	uf.parent[p] = root
	// 	p = tmp
	// }

	return root
}

// Union ...
func (uf *UnionFind) Union(p, q int) {
	rootp := uf.Find(p)
	rootq := uf.Find(q)

	if rootp == rootq {
		return
	}

	// 只有當後者等級大於前者等級，前者才會融入後者
	if uf.rank[rootq] > uf.rank[rootp] {
		uf.parent[rootp] = rootq
	} else {
		// 否則後者要融入前者團體
		uf.parent[rootq] = rootp
		// 若雙方等級相同，進行合併後，該root等級會提升
		if uf.rank[rootq] == uf.rank[rootp] {
			uf.rank[rootp]++
		}
	}

	// 每一次合併都會減少團體數量
	uf.count--
}

func main() {
	accounts := [][]string{{"Johnson", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}
	fmt.Println(accountsMerge(accounts))
}

func accountsMerge(accounts [][]string) [][]string {
	res := [][]string{}

	idToName := map[int]string{}
	emailToID := map[string]int{}
	idToEmails := map[int][]string{}

	uf := UnionFind{}
	uf.Init(len(accounts))

	for id, acc := range accounts {
		// 製作 idToName
		idToName[id] = acc[0]

		// 進行合併
		for i := 1; i < len(acc); i++ {
			if pid, ok := emailToID[acc[i]]; ok {
				uf.Union(id, pid)
			}

			// 製作 emailToID
			emailToID[acc[i]] = id
		}
	}

	// 使用 uf 與 emailToID 製作 idToEmails
	for email, id := range emailToID {
		rootID := uf.Find(id)
		idToEmails[rootID] = append(idToEmails[rootID], email)
	}

	// 使用 idToEmails 製作 res
	for id, emails := range idToEmails {
		sort.Strings(emails)
		res = append(res, append([]string{idToName[id]}, emails...))
	}

	return res

}
