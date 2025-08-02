package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		productMap := make(map[string]int)
		intMap := make(map[int]struct{})
		var n int
		fmt.Fscan(in, &n)
		var product string
		var price int
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &product, &price)
			str := product + ":" + strconv.Itoa(price)
			productMap[str] = price
			intMap[price] = struct{}{}
		}
		var resultString string
		fmt.Fscan(in, &resultString)

		sl := strings.Split(resultString, ",")
		fmt.Fprintln(out, Check(intMap, productMap, sl))

	}
}

func Check(intMap map[int]struct{}, productMap map[string]int, resultString []string) string {
	resultMap := make(map[int]struct{})

	for _, el := range resultString {
		if value, ok := productMap[el]; ok {
			if _, ok1 := resultMap[value]; ok1 {
				return "NO"
			} else {
				resultMap[value] = struct{}{}
				delete(productMap, el)
			}
		} else {
			return "NO"
		}
	}

	if len(intMap) != len(resultMap) {
		return "NO"
	}

	return "YES"
}
