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

	var n int
	fmt.Fscan(in, &n)
	var str string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &str)
		result := Run(str)
		fmt.Fprintln(out, result)
	}
}

func Run(str string) string {
	run := []rune(str)
	if len(run) == 1 {
		return "YES"
	}
	symbol := string(run[0])

	if string(run[len(run)-1]) != symbol {
		return "NO"
	}

	for i := 1; i < len(run)-1; i++ {
		if string(run[i]) != symbol {
			if string(run[i-1]) != symbol || string(run[i+1]) != symbol {
				return "NO"
			}
		}
	}

	return "YES"
}
