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
	for i := 0; i < t; i++ {
		var n, m int

		fmt.Fscan(in, &n, &m)

		if n == 1 {
			fmt.Fprint(out, "1\n1 1 R")
		} else if m == 1 {
			fmt.Fprint(out, "1\n1 1 D")
		} else {
			fmt.Fprint(out, "2\n1 1 D\n", n, " ", m, " U")
		}

		if i < t-1 {
			fmt.Fprintln(out) // Перенос только между тестами
		}
	}

}
