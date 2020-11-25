package main

import "fmt"

func main() {
	A := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	fmt.Println(maxTurbulenceSize(A))
}

func maxTurbulenceSize(A []int) int {
	if len(A) == 1 {
		return 1
	}

	l, r, res := 0, 0, 1
	flag, n := A[1]-A[0], A[0]

	for l < len(A) {
		fmt.Println("l:r", l, r)
		if r < len(A)-1 && ((A[r+1] > A[r] && flag > 0) || (A[r+1] < A[r] && flag < 0) || (l == r)) {
			r++
			flag = n - A[r]
			n = A[r]
		} else {
			if r != l && flag != 0 {
				res = max(res, r-l+1)
			}
			l++
		}
	}
	return res
}

func max(a, b int) int {
	if b > a {
		return b
	}

	return a
}
