package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Read the input line containing the array
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	// Parse the array of integers
	numsStr := strings.Split(line, " ")
	nums := make([]int, len(numsStr))
	for i, numStr := range numsStr {
		nums[i], _ = strconv.Atoi(numStr)
	}

	// Read the value of k
	kLine, _ := reader.ReadString('\n')
	k, _ := strconv.Atoi(strings.TrimSpace(kLine))

	// Call the function and get the result
	result := TopKFrequent(nums, k)

	// Write the result to the output
	for i, num := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, num)
	}
	fmt.Fprintln(writer)
}

type result struct {
	Num   int
	Count int
}

func TopKFrequent(nums []int, k int) []int {
	numMap := make(map[int]*result)
	var res []result
	var es []int
	for _, num := range nums {
		if _, ok := numMap[num]; !ok {
			numMap[num] = &result{num, 1}
		} else {
			t := numMap[num]
			t.Count++
		}
	}

	for _, value := range numMap {
		res = append(res, *value)
	}

	sortRes := SortRes(res)
	for i := len(sortRes) - 1; i >= len(sortRes)-k; i-- {
		es = append(es, sortRes[i].Num)
	}

	return es
}

func SortRes(res []result) []result {
	var left []result
	var right []result
	if len(res) <= 0 {
		return res
	}

	m := res[0]

	for i := 1; i < len(res); i++ {
		if res[i].Count < m.Count {
			left = append(left, res[i])
		} else {
			right = append(right, res[i])
		}
	}

	return append(append(SortRes(left), m), SortRes(right)...)
}
