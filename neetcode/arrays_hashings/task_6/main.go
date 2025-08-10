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

	var n, num int
	var nums []int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &num)
		nums = append(nums, num)
	}

	nums = ProductExceptSelf(nums)
	for _, num = range nums {
		fmt.Fprintf(out, "%v ", num)
	}

}

func ProductExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	prefix := 1
	for i, n := range nums {
		ret[i] = prefix
		prefix *= n
	}
	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= suffix
		suffix *= nums[i]
	}

	return ret
}
