package main

import "fmt"

func main() {
	// fmt.Println(permute([]int{2, 5, 3, 6}))
	s := []int{1, 2, 3}
	bar(s)
	fmt.Println(s)
	fmt.Printf("%p", s)

}
func foo(s []int) {
	s = append(s, 5)
}

func bar(s []int) {
	s = s[:]
	fmt.Printf("%p", s)

}
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	res, p, index := [][]int{}, []int{}, 0
	used := make([]bool, len(nums))

	// 之所以只有 res 需要用 pointer 而 used 不用
	// 是因為 res 會一直 append 擴增，會改變 slice header，所以需要用指標
	// used 不用的原因在於 used 是做直接修改值的動作，此動作本身就會修改到 slice 指向的底層 array
	// 因此不必特別使用指標存取
	genPermutation(nums, index, p, &res, used)

	return res
}

func genPermutation(nums []int, index int, p []int, res *[][]int, used []bool) {
	if index == len(nums) {
		// 為什麼一定要 copy？
		// 因為 slice p 的 array 一直在更動
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}

	for i, num := range nums {
		if !used[i] {
			(used)[i] = true
			p = append(p, num)
			genPermutation(nums, index+1, p, res, used)
			(used)[i] = false
			p = p[:len(p)-1]
		}
	}
}
