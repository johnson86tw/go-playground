package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	res := threeSumClosest([]int{-4, -1, 1, 2}, 1)
	fmt.Println(res)
}

func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt16
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			if i-1 >= 0 && nums[i] == nums[i-1] {
				continue
			}

			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res = sum
					diff = abs(sum - target)
				}

				if sum == target {
					return res
				} else if sum < target {
					j++
				} else {
					k--
				}
			}
		}
	}

	return res
}

// 解法二 暴力解法 O(n^3)
func solution2(nums []int, target int) int {
	res, difference := 0, math.MaxInt16
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if abs(nums[i]+nums[j]+nums[k]-target) < difference {
					difference = abs(nums[i] + nums[j] + nums[k] - target)
					res = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
