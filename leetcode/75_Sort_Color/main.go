package main

import "fmt"

func main() {
	nums := []int{2, 1, 1, 2, 0, 0, 2}
	sortColors(nums)
	fmt.Println(nums)

}

// 三路快排
func sortColors1(nums []int) {
	l, c, r := []int{}, []int{}, []int{}
	for _, n := range nums {
		if n == 0 {
			l = append(l, n)
		} else if n == 1 {
			c = append(c, n)
		} else if n == 2 {
			r = append(r, n)
		}
	}

	l = append(l, c...)
	l = append(l, r...)
	nums = nums[:0]
	nums = append(nums, l...)
}

// 游標移動法
func sortColors(nums []int) {
	r, w, b := 0, 0, 0
	for _, n := range nums {
		if n == 0 {
			nums[b] = 2
			b++
			nums[w] = 1
			w++
			nums[r] = 0
			r++
		}

		if n == 1 {
			nums[b] = 2
			b++
			nums[w] = 1
			w++
		}

		if n == 2 {
			b++
		}
	}
}
