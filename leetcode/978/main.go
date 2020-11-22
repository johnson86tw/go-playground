package main

import "fmt"

func main() {
	A := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	fmt.Println("ans: ", maxTurbulenceSize(A))
}

func maxTurbulenceSize(A []int) int {
	if len(A) == 1 {
		return 1
	}

	l, r, res := 0, 0, 1
	flag := -1

	for l < len(A) {
		if r >= len(A)-1 {
			break
		}
		fmt.Println("------")

		a := A[r]
		b := A[r+1]
		fmt.Println("AvsB: ", a, b)

		if a == b {
			flag = -1
			l = r
		}

		if flag == -1 {
			if a > b {
				flag = 0
				r++
			} else if a < b {
				flag = 1
				r++
			} else {
				r++
				l++
			}
			continue
		}

		flag = flip(flag)

		fmt.Println("flag", flag)
		if (flag == 0 && a > b) || (flag == 1 && a < b) {
			r++
			res = max(res, r-l+1)

		} else if l < r {
			l = r
		} else {
			r++
		}
		fmt.Println(l, r)

	}

	return res

}

func flip(f int) int {
	if f == 0 {
		return 1
	} else if f == 1 {
		return 0
	} else {
		return -1
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
