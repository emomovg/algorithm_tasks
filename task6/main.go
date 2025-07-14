package main

import "fmt"

func main() {
	var letterCount, wordCount int
	fmt.Scan(&letterCount, &wordCount)
	var letters []string
	var words []string
	var st string
	var result string
	for i := 0; i < letterCount; i++ {
		fmt.Scan(&st)
		letters = append(letters, st)
	}

	for i := 0; i < wordCount; i++ {
		fmt.Scan(&st)
		words = append(words, st)
	}

	for i := 0; i < wordCount; i++ {
		result = Run(letters, words[i])
		fmt.Println(result)
	}
}

func Run(letters []string, word string) string {
	literals := []rune(word)
	litCount := make(map[string]int, len(literals))
	letCount := make(map[string]int, len(literals))

	for i := 0; i < len(literals); i++ {
		litCount[string(literals[i])]++
	}

	for i := 0; i < len(letters); i++ {
		letCount[letters[i]]++
	}

	for key, value := range litCount {
		if value != letCount[key] {
			return "NO"
		}
	}

	return "YES"
}
