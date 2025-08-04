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

	var n int
	fmt.Fscan(in, &n)
	var sl []int
	var num int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &num)
		sl = append(sl, num)
	}

	result := ContainsDuplicate(sl)

	fmt.Fprintln(out, result)
}

func ContainsDuplicate(nums []int) bool {
	numsMap := make(map[int]struct{}, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, ok := numsMap[nums[i]]; ok {
			return true
		}

		numsMap[nums[i]] = struct{}{}
	}

	return false
}
