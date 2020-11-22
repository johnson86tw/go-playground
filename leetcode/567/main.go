package main

import "fmt"

func main() {
	s1 := "abe"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}

	freq := [256]int{}

	for i := range s1 {
		freq[s1[i]-'a']++
	}

	r, l := 0, 0
	count := len(s1)
	for r < len(s2) {
		if freq[s2[r]-'a'] > 0 {
			freq[s2[r]-'a']--
			r++
			count--

			if count == 0 {
				return true
			}
		} else if l == r {
			r++
			l++
		} else if l < r && freq[s2[r]-'a'] == 0 {
			freq[s2[l]-'a']++
			count++
			l++
		}

	}
	return false
}

// author version
func checkInclusion2(s1 string, s2 string) bool {
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}

	freq := [256]int{}

	for i := range s1 {
		freq[s1[i]-'a']++
	}

	r, l := 0, 0
	count := len(s1)
	for r < len(s2) {
		if freq[s2[r]-'a'] > 0 {
			count--
		}

		if count == 0 {
			return true
		}

		freq[s2[r]-'a']--
		r++

		if r-l == len(s1) {
			// 代表曾經大於零，否則會是負的
			if freq[s2[l]-'a'] >= 0 {
				count++
			}

			freq[s2[l]-'a']++
			l++
		}

	}
	return false
}
