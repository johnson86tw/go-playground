package main

func main() {

}

func subsets(nums []int) [][]int {
	res := [][]int{}

	for i := 0; i <= len(nums); i++ {
		p := []int{}
		genSubsets(nums, i, p, &res, 0)
	}

	return res
}

func genSubsets(nums []int, limit int, p []int, res *[][]int, start int) {
	if len(p) == limit {
		tmp := make([]int, len(p))
		copy(tmp, p)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(nums)-(limit-len(p))+1; i++ {
		p = append(p, nums[i])
		genSubsets(nums, limit, p, res, i+1)
		p = p[:len(p)-1]
	}
}

// 解法二
// func subsets1(nums []int) [][]int {

// }

// 解法三：位运算的方法
func subsets2(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	sum := 1 << uint(len(nums))
	for i := 0; i < sum; i++ {
		stack := []int{}
		tmp := i                              // i 从 000...000 到 111...111
		for j := len(nums) - 1; j >= 0; j-- { // 遍历 i 的每一位
			if tmp&1 == 1 {
				stack = append([]int{nums[j]}, stack...)
			}
			tmp >>= 1
		}
		res = append(res, stack)
	}
	return res
}
