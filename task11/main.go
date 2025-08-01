package main

import (
	"bufio"
	"fmt"
	"math"
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
	var n int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)

		var nums1, nums2 []int
		var num int
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &num)
			nums1 = append(nums1, num)
		}
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &num)
			nums2 = append(nums2, num)
		}
		fmt.Fprintln(out, Run(nums1, nums2))
	}
}

func Run(a, b []int) int {
	result := 1
	for i := 0; i < len(a); i++ {
		del := b[i]/(i+1) - (a[i]-1)/(i+1)

		result *= del
	}
	result = result % int(math.Pow(10, 9)+7)
	return result
}
