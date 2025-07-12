package main

import (
	"fmt"
)

type user struct {
	numberRequest    int
	lastNotification int
}

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	var requests [][]int
	for i := 0; i < q; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		requests = append(requests, []int{x, y})
	}
	finalResults := ProcessRequests(requests)

	for _, val := range finalResults {

		fmt.Println(val)
	}
}

func ProcessRequests(requests [][]int) []int {
	var users []user
	var results []int
	var currentRequestNumber int

	for _, req := range requests {
		x := req[0]
		y := req[1]

		if x == 1 {
			currentRequestNumber++
			users = append(users, user{numberRequest: currentRequestNumber, lastNotification: y})
		} else if x == 2 {
			var lastMatch int = 0

			for _, u := range users {
				if u.lastNotification == 0 || u.lastNotification == y {
					lastMatch = u.numberRequest
				}
			}
			results = append(results, lastMatch)
		}
	}
	return results
}
