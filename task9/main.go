package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	var question string
	for i := 0; i < t; i++ {
		var facts []string
		reader := bufio.NewReader(os.Stdin)
		question, _ = reader.ReadString('\n')

		question = strings.TrimSpace(question)
		for i := 0; i < 3; i++ {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')

			text = strings.TrimSpace(text)
			facts = append(facts, text)
		}
		fmt.Fprintln(out, Run(question, facts))
	}
}

func Run(question string, facts []string) int {
	ages := make(map[string]int, 3)
	quesSl := strings.Fields(question)
	friendName := strings.ReplaceAll(quesSl[len(quesSl)-1], "?", "")
	i := 0
	for {
		j := i % 3
		factSl := strings.Fields(facts[j])

		switch factSl[4] {
		case "old":
			ages[factSl[0]], _ = strconv.Atoi(factSl[2])
		case "age":
			if factSl[3] == "same" {
				if _, ok := ages[factSl[6]]; ok {
					ages[factSl[0]] = ages[factSl[6]]
				}
			}
		case "younger":
			if _, ok := ages[factSl[6]]; ok {
				age, _ := strconv.Atoi(factSl[2])
				ages[factSl[0]] = ages[factSl[6]] - age
			}
			if _, ok := ages[factSl[0]]; ok {
				age, _ := strconv.Atoi(factSl[2])
				ages[factSl[6]] = ages[factSl[0]] + age
			}
		case "older":
			if _, ok := ages[factSl[6]]; ok {
				age, _ := strconv.Atoi(factSl[2])
				ages[factSl[0]] = ages[factSl[6]] + age
			}
			if _, ok := ages[factSl[0]]; ok {
				age, _ := strconv.Atoi(factSl[2])
				ages[factSl[6]] = ages[factSl[0]] - age
			}
		}

		if _, ok := ages[friendName]; ok {
			break
		}
		i++
	}
	return ages[friendName]
}
