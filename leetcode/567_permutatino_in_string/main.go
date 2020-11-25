package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	freq := [256]int{}

	for i := range s1 {
		freq[s1[i]-'a']++
	}

	l, r := 0, 0
	count := len(s2)

	for r < len(s2) {
		if freq[s2[r]-'a'] > 0 {
			r++
			count--

			if count == 0 {
				return true
			}
		} else {
			l++
			count++
		}
	}

	return false
}
