package main

import "fmt"

func main() {
	fmt.Println(permute([]int{2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) < 1 {
		return [][]int{}
	}
	index, p, res := 0, []int{}, [][]int{}
	used := make([]bool, len(nums))
	genPermutation(nums, index, p, &res, &used)

	return res
}

func genPermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	// 判斷遞迴停止的條件
	if index == len(nums) {
		// 為什麼要用 tmp
		tmp := []int{}
		tmp = append(tmp, p...)
		tmp = append(tmp, nums[len(nums)-1])
		*res = append(*res, tmp)
		return
	}
	// for迴圈與遞迴
	for i, num := range nums {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, num)
			genPermutation(nums, index+1, p, res, used)
			(*used)[i] = false
			p = p[:len(nums)-1]
		}
	}

}
