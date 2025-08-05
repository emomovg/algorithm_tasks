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

	var n, target int
	fmt.Fscan(in, &n, &target)
	var nums []int
	var num int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &num)
		nums = append(nums, num)
	}

	fmt.Fprintln(out, TwoSum(nums, target))
}

func TwoSum(nums []int, target int) []int {
	mapDiff := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, ok := mapDiff[nums[i]]; ok {
			return []int{index, i}
		}
		diff := target - nums[i]
		mapDiff[diff] = i
	}
	return []int{0, 0}
}
