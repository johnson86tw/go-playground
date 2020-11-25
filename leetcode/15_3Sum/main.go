package main

import (
	"fmt"
	"sort"
)

func main() {
	ary := []int{-1, 0, 1, 2, -1, -4, -2, -2, 4}
	fmt.Println(threeSum(ary))

}

func threeSum(nums []int) [][]int {
	// initiate answer
	res := [][]int{}
	// 建立 counter，用 map 整理重複數字的數量
	counter := map[int]int{}
	for _, n := range nums {
		counter[n]++
	}
	// 挑出 uniqNum 的陣列
	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	// 用 sort.Ints 由小到大排序 uniqNums
	sort.Ints(uniqNums)

	// 用雙迴圈配對同一個陣列不重複的兩個數
	for i := 0; i < len(uniqNums); i++ {
		// 找出 三個一樣的 (都是0)
		if uniqNums[i] == 0 && counter[uniqNums[i]] > 2 {
			res = append(res, []int{0, 0, 0})
		}

		for j := i + 1; j < len(uniqNums); j++ {
			// 找出 二個一樣的
			if uniqNums[i]*2+uniqNums[j] == 0 && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}

			if uniqNums[j]*2+uniqNums[i] == 0 && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[j], uniqNums[j], uniqNums[i]})
			}

			// 找出 都不一樣的
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}

	return res
}
