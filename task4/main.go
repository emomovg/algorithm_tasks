package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	var n, q int
	var queue []int
	var t int

	fmt.Scan(&n, &q)

	for j := 0; j < q; j++ {
		fmt.Scan(&t)
		queue = append(queue, t)
	}

	fmt.Println(Run(queue, n))

}

type Person struct {
	index  int
	window int
	action string
}

func Run(queue []int, queueCount int) string {
	if len(queue) > queueCount {
		return "x"
	}
	persons := make([]Person, 0, len(queue))
	for i := 0; i < len(queue); i++ {
		persons = append(persons, Person{
			index:  i,
			window: queue[i],
		})
	}

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].window < persons[j].window
	})
	if persons[0].window > 1 {
		persons[0].window--
		persons[0].action = "-"
	} else {
		persons[0].action = "0"
	}
	for i := 1; i < len(persons); i++ {
		if persons[i].window-persons[i-1].window == 1 {
			persons[i].action = "0"
		}

		if persons[i].window-persons[i-1].window > 1 {
			persons[i].window--
			persons[i].action = "-"
		}

		if persons[i].window-persons[i-1].window == 0 {
			persons[i].window++
			persons[i].action = "+"
		}

		if persons[i].window-persons[i-1].window < 0 {
			return "x"
		}

		if persons[i].window > queueCount {
			return "x"
		}
	}

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].index < persons[j].index
	})

	result := strings.Builder{}
	for _, p := range persons {
		result.WriteString(p.action)
	}

	return result.String()
}
