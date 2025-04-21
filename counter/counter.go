package counter

import (
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func CountWords(lines []string) (map[string]int, int) {
	wordCounts := make(map[string]int)
	totalWords := 0

	for _, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			cleaned := strings.ToLower(strings.Trim(word, ".,!?\"'():;"))
			if cleaned == "" {
				continue
			}
			wordCounts[cleaned]++
			totalWords++
		}
	}
	return wordCounts, totalWords
}

func GetTopWords(wordCounts map[string]int, topN int) []WordCount {
	var wordList []WordCount
	for word, count := range wordCounts {
		wordList = append(wordList, WordCount{Word: word, Count: count})
	}

	sort.Slice(wordList, func(i, j int) bool {
		return wordList[i].Count > wordList[j].Count
	})

	if len(wordList) < topN {
		topN = len(wordList)
	}
	return wordList[:topN]
}

func normalizeWord(word string) string {
	word = strings.ToLower(word)
	word = strings.Trim(word, ".,!?\"'():;")
	return word
}