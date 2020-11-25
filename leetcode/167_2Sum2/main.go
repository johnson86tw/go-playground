package main

import "fmt"

func main() {
	nums := []int{0, 0, 3, 4}
	target := 4
	fmt.Println(twoSumTwo(nums, target))

}

// Two pointers
func twoSum2(numbers []int, target int) []int {
	var res []int
	for l, r := 0, len(numbers)-1; l < r; {
		sum := numbers[l] + numbers[r]
		if sum == target {
			res = []int{l + 1, r + 1}
			break
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return res
}

// Hash table (比上一個解法多了 O(n) 的空間複雜度)
func twoSumTwo(numbers []int, target int) []int {
	m := map[int]int{}

	for i, num := range numbers {
		another := target - num

		if _, ok := m[another]; ok {
			return []int{m[another] + 1, i + 1}
		}

		// 此行會隨著陣列長度消耗記憶體
		m[num] = i
	}

	return nil
}
