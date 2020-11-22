package main

import "fmt"

func main() {
	fmt.Println(ladderLength("cet",
		"ism",
		[]string{"hot", "dot", "dog", "lot", "log"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := getWordMap(wordList, beginWord)
	queue, level := []string{}, 0
	queue = append(queue, beginWord)

	for len(queue) > 0 {
		level++

		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			que := queue[0]
			queue = queue[1:]
			candidates := getCandidates(que)
			for _, c := range candidates {
				if _, ok := wordMap[c]; ok {
					if c == endWord {
						return level + 1
					}
					// ex. level 3 不可能再用 level 2 會有的任何 candidates 了，所以可以 delete
					delete(wordMap, c)
					queue = append(queue, c)
				}
			}

		}
	}

	return 0

}

func getWordMap(wordList []string, beginWord string) map[string]int {
	wordMap := make(map[string]int)
	for i, word := range wordList {
		if _, ok := wordMap[word]; !ok {
			if word != beginWord {
				wordMap[word] = i
			}
		}
	}
	return wordMap
}

func getCandidates(word string) []string {
	var res []string
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] != byte(int('a')+i) {
				res = append(res, word[:j]+string(int('a')+i)+word[j+1:])
			}
		}
	}
	return res
}
