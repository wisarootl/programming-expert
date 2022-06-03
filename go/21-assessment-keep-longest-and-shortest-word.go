// Keep Longest And Shortest Word

package main

import "fmt"

func keepLongestAndShortestWordHelper(words []string) []string {
	if len(words) == 0 {
		return words
	}

	longest := words[0]
	shortest := words[0]
	longestIdx := 0
	shortestIdx := 0

	for idx, word := range words {
		if len(word) > len(longest) {
			longest = word
			longestIdx = idx
		}
		if len(word) < len(shortest) {
			shortest = word
			shortestIdx = idx
		}
	}

	if longest == shortest {
		return []string{longest}
	} else if longestIdx > shortestIdx {
		return []string{shortest, longest}
	} else {
		return []string{longest, shortest}
	}
}

func KeepLongestAndShortestWord(wordSlices *[][]string) {
	// Keep Longest And Shortest Word
	for wordsIdx, words := range *wordSlices {
		newWords := keepLongestAndShortestWordHelper(words)
		(*wordSlices)[wordsIdx] = newWords
	}
}

func main() {
	words := &[][]string{ {"best", "course", "yes"}, {"hello"}, {"a", "ab", "abc", "abcd"}, {} }
	fmt.Println(words)
}
