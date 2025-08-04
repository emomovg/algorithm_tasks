package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var out *bufio.Writer
	in := os.Stdin
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var s, t string

	fmt.Fscan(in, &s, &t)

	fmt.Fprintln(out, IsAnagram(s, t))
}

func IsAnagram(s string, t string) bool {
	srun := []rune(s)
	trun := []rune(t)
	if len(srun) != len(trun) {
		return false
	}
	diffSt := make(map[rune]int, len(srun))

	for i := 0; i < len(trun); i++ {
		diffSt[trun[i]]++
	}

	for i := 0; i < len(srun); i++ {
		diffSt[srun[i]]--
	}

	for _, count := range diffSt {
		if count != 0 {
			return false
		}
	}

	return true

}
