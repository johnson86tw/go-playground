package main

import (
	"strings"
)

func main() {
	// input := "/a/./b/////../../"
	// fmt.Println(simplifyPath(input))

	// s := []int{1, 5, 3, 6, 7, 4, 3, 2, 6}
	// i := 5

	// ss := remove(s, i)
	// fmt.Println(s)
	// fmt.Println(ss)

}

func remove(ary []int, i int) []int {
	copy(ary[:i], ary[:i+1])
	ary = ary[:len(ary)-1]
	return ary
}

func simplifyPath(path string) string {
	stack := []string{}
	arr := strings.Split(path, "/")

	for _, cur := range arr {

		if cur == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}

		if cur != "." && len(cur) > 0 {
			stack = append(stack, cur)
		}

	}

	return "/" + strings.Join(stack, "/")
}
