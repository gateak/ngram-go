package main

import (
	"fmt"
	"strings"
)

// Tokenize splits the text into words.
func Tokenize(text string) []string {
	return strings.Fields(text)
}

// BuildNGrams creates n-grams from the given tokens.
func BuildNGrams(tokens []string, n int) [][]string {
	var ngrams [][]string
	for i := 0; i <= len(tokens)-n; i++ {
		ngrams = append(ngrams, tokens[i:i+n])
	}
	return ngrams
}

// CountNGrams counts the frequency of each n-gram.
func CountNGrams(ngrams [][]string) map[string]int {
	counts := make(map[string]int)
	for _, ngram := range ngrams {
		key := strings.Join(ngram, " ")
		counts[key]++
	}
	return counts
}

func main() {
	text := "this is a sample text for the n-gram model"
	tokens := Tokenize(text)
	ngrams := BuildNGrams(tokens, 2) // 2-grams
	counts := CountNGrams(ngrams)

	fmt.Println("N-grams and their counts:")
	for ngram, count := range counts {
		fmt.Printf("%s: %d\n", ngram, count)
	}
}

