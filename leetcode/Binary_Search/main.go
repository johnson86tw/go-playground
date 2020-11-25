package main

import "fmt"

func main() {
	n := []int{1, 3, 5, 7, 9, 10, 11, 12}
	fmt.Println(binarySearch(n, 11))

}

func binarySearch(nums []int, target int) int {
	high, low := len(nums)-1, 0

	for low <= high {
		mid := low + (high-low)>>1
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func myBinarySearch(nums []int, target int) int {
	i := len(nums) / 2
	if i == 0 && target != nums[0] {
		return -1
	}

	center := nums[i]
	if target > center {
		return binarySearch(nums[i+1:], target)
	} else if target < center {
		return binarySearch(nums[:i], target)
	} else {
		return i
	}
}
