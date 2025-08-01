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

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		Run(in, out)
	}

}
func Run(in *bufio.Reader, out *bufio.Writer) {
	var n int
	var digits []int
	var result int
	fmt.Fscan(in, &n)
	digits = *Digits(&digits, n)

	l := len(digits)
	if l == 1 {
		result = 0
		fmt.Fprintln(out, result)
		return
	}

	for i := l - 1; i >= 1; i-- {
		if digits[i] < digits[i-1] {
			digits = append(digits[:i], digits[i+1:]...)
			break
		}
	}
	if len(digits) == l {
		digits = digits[1:]
	}
	for i := len(digits) - 1; i >= 0; i-- {
		result = 10*result + digits[i]
	}

	fmt.Fprintln(out, result)
}

func Digits(digits *[]int, n int) *[]int {
	if n/10 == 0 {
		*digits = append(*digits, n%10)
		return digits
	}
	*digits = append(*digits, n%10)
	return Digits(digits, n/10)
}
