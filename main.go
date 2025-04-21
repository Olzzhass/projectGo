package main

import (
	"log"
	"os"
	"projectGolang/counter"
	"projectGolang/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <path-to-text-file>")
	}
	filePath := os.Args[1]

	lines, err := file.ReadFileLines(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	wordCounts, totalWords := counter.CountWords(lines)
	topWords := counter.GetTopWords(wordCounts, 5)

	log.Printf("Total words in file: %d\n", totalWords)
	log.Println("Top 5 most frequent words:")
	for i, word := range topWords {
		log.Printf("%d. %s (%d occurrences)", i+1, word.Word, word.Count)
	}
}
