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

	var n, k int

	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &k)
		var sl []int
		var el int
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &el)
			sl = append(sl, el)
		}
		fmt.Fprintln(out, Diff(sl, k))
	}

}

func Diff(sl []int, k int) string {
	for i := 0; i < len(sl)-k+1; i++ {
		el := sl[i]
		for j := i; j < k+i; j++ {
			sl[j] -= el
			if sl[j] < 0 {
				return "NO"
			}
		}
	}
	for i := 0; i < len(sl); i++ {
		if sl[i] != 0 {
			return "NO"
		}
	}

	return "YES"
}
