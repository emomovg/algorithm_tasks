package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var t int
	fmt.Fscan(in, &t)

	for p := 0; p < t; p++ {
		var n int
		fmt.Fscan(in, &n)
		var words []string
		var str string
		for q := 0; q < n; q++ {
			fmt.Fscan(in, &str)
			words = append(words, str)
		}
		var result int
		le := len(words)
		for i := 0; i < le; i++ {

			for j := i + 1; j < le; j++ {
				even, odd := getWords(words[i])
				even1, odd1 := getWords(words[j])
				if (even == even1 && even != "") || (odd == odd1 && odd != "") {
					result++
				}
			}
		}
		fmt.Fprintln(out, result)
	}

}

func getWords(word string) (string, string) {
	var evenWord, oddWord []rune
	run := []rune(word)

	for i := 0; i < len(run); i++ {
		if (i+1)%2 == 0 {
			evenWord = append(evenWord, run[i])
		} else {
			oddWord = append(oddWord, run[i])
		}
	}

	return string(evenWord), string(oddWord)
}
