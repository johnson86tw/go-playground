package main

import "fmt"

func main() {
	x := 'a'
	fmt.Println(x)

	y := "a"
	fmt.Println(y[0])
}

// func lengthOfLongestSubstring(s string) int {
// 	// 如果是空字串，回傳0
// 	// 建立一個 256 bytes 的 bool array 檢查是否出現過
// 	// 建立 left, right, and result
// 	// 迴圈直到 left 大於等於字串長度

// 	// 如果右邊的字出現過 左邊的字為 false 左++
// 	// 否則 右邊的字為 true 右++
// 	// 判斷是否更新 result
// 	// 判斷是否跳出迴圈
// }

func lengthOfLongestSubstring1(s string) int {
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
		if right >= len(s) {
			break
		}
	}
	return result
}

// 解法二 滑动窗口
func lengthOfLongestSubstring2(s string) int {
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
