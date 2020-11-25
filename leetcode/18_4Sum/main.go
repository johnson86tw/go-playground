package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))

}

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	counter := map[int]int{}
	for _, n := range nums {
		counter[n]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}

	// 排序的用意？
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if uniqNums[i]*4 == target && counter[uniqNums[i]] >= 4 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[i]})
		}

		for j := i + 1; j < len(uniqNums); j++ {
			if uniqNums[i]*3+uniqNums[j] == target && counter[uniqNums[i]] >= 3 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[j]})
			}

			if uniqNums[i]*2+uniqNums[j]*2 == target && counter[uniqNums[i]] >= 2 && counter[uniqNums[j]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			if uniqNums[i]+uniqNums[j]*3 == target && counter[uniqNums[j]] >= 3 {
				res = append(res, []int{uniqNums[j], uniqNums[j], uniqNums[j], uniqNums[i]})
			}

			for k := j + 1; k < len(uniqNums); k++ {
				if uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target && counter[uniqNums[i]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}
				if uniqNums[i]+uniqNums[j]*2+uniqNums[k] == target && counter[uniqNums[j]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[k]})
				}
				if uniqNums[i]+uniqNums[j]+uniqNums[k]*2 == target && counter[uniqNums[k]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[k]})
				}

				d := target - uniqNums[i] - uniqNums[j] - uniqNums[k]
				if counter[d] > 0 && d > uniqNums[k] {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], d})
				}
			}
		}
	}

	return res

}
