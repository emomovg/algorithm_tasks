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

	var n, t int
	fmt.Fscan(in, &n, &t)
	symbols := make(map[string]int, n)
	var symbol string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &symbol)
		symbols[symbol]++
	}

	var pass string
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &pass)

		result := IsValidPass(symbols, pass, n)
		fmt.Fprintln(out, result)
	}
}

func IsValidPass(symbols map[string]int, pass string, difSymbolsLen int) string {
	run := []rune(pass)

	if len(run) != difSymbolsLen {
		return "NO"
	}
	passMap := make(map[string]int, len(run))
	for _, value := range run {
		if _, ok := symbols[string(value)]; !ok {
			return "NO"
		}
		passMap[string(value)]++
	}

	for key, value := range symbols {
		if _, ok := passMap[key]; !ok {
			return "NO"
		}
		if value != passMap[key] {
			return "NO"
		}
	}

	return "YES"
}
