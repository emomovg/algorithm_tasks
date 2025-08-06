package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var out *bufio.Writer
	in := os.Stdin
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	fmt.Fscan(in, &n)

	var strs []string
	var str string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &str)
		strs = append(strs, str)
	}

	fmt.Fprintln(out, GroupAnagrams(strs))
}

func GroupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	var result [][]string
	anagramMap[sortStr(strs[0])] = []string{strs[0]}
	for i := 1; i < len(strs); i++ {
		st := sortStr(strs[i])
		anagramMap[st] = append(anagramMap[st], strs[i])
	}
	for _, anagrams := range anagramMap {
		result = append(result, anagrams)
	}
	return result
}

func sortStr(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
