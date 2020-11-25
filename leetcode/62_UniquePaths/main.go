package main

import (
	"fmt"
)

func main() {
	// uniquePaths(3, 2)
	fmt.Println(uniquePaths(3, 2))
	// fmt.Println(fibonacci2(5))

}

func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < m; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[n-1][m-1]
}

// Divide and Conquer
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Dynamic programming
func fibonacci2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	// fmt.Println(dp)

	return dp[n]
}

// func uniquePaths2(m int, n int) uint {
// 	if m <= 0 || n <= 0 {
// 		log.Panic("parameter cannot be zero")
// 	}
// 	r, l := uint(m-1), uint(n-1)
// 	divisor := factorial(r + l)
// 	fmt.Println(divisor)
// 	dividend := (factorial(r) * factorial(l))
// 	fmt.Println(dividend)
// 	return divisor / dividend
// }

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorial2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] * i
	}

	return dp[n]
}
