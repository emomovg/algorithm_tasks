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
		var result int
		var str string
		oddMap := make(map[string]int, n)
		evenMap := make(map[string]int, n)
		eqMap := make(map[string]int, n)
		for q := 0; q < n; q++ {
			fmt.Fscan(in, &str)
			if len(str) == 1 {
				result += oddMap[str]
				oddMap[str]++
				continue
			}
			even, odd := getWords(str)

			result += evenMap[even]
			result += oddMap[odd]
			result -= eqMap[str]

			evenMap[even]++
			oddMap[odd]++
			eqMap[str]++
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
