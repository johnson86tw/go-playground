package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("bbb"))

}

func lengthOfLongestSubstring(s string) int {
	col := []map[string]int{}
	ls := strings.Split(s, "")

	m := make(map[string]int)
	for i, c := range ls {
		if len(m) == 0 || m[c] == 0 {
			// map 裡面不存在
			m[c]++
		} else if m[c] > 0 {
			// map 裡面已經存在
			col = append(col, m)
			m = make(map[string]int)
			m[ls[i-1]]++
			m[c]++
		}

		if i == len(ls)-1 {
			col = append(col, m)
		}
	}

	res := 0
	for _, m := range col {
		if len(m) > res {
			res = len(m)
		}
	}

	return res

}

func solution1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将X标记为 false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// 解法二 滑动窗口
func solution2(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
