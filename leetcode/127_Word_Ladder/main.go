package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	// fmt.Println(getCandidates("hot"))
	// fmt.Println(getWordMap([]string{"hot", "dot", "dog", "lot", "log", "cog"}, "hit"))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := getWordMap(wordList, beginWord)
	queue, depth := []string{}, 0
	queue = append(queue, beginWord)

	for len(queue) > 0 {
		depth++
		qlen := len(queue)

		for i := 0; i < qlen; i++ {
			q := queue[0]
			queue = queue[1:]

			candidates := getCandidates(q)
			for _, c := range candidates {
				if _, ok := wordMap[c]; ok {
					if c == endWord {
						return depth + 1
					}
					delete(wordMap, c)
					queue = append(queue, c)
				}
			}
		}
	}
	return 0
}

func getWordMap(wordList []string, beginWord string) map[string]int {
	res := map[string]int{}
	for i, word := range wordList {
		if _, ok := res[word]; !ok {
			res[word] = i
		}
	}
	return res
}

func getCandidates(word string) []string {
	res := []string{}
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] != byte(int('a'+i)) {
				res = append(res, word[:j]+string(int('a'+i))+word[j+1:])
			}
		}
	}
	return res
}
