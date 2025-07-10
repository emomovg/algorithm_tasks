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
	var x, y int
	var arr []user
	var result []int
	var last int
	var number int
	for i := 0; i < q; i++ {
		fmt.Scan(&x, &y)
		if x == 1 {
			number++
			arr = append(arr, user{numberRequest: number, lastNotification: y})
		}
		if x == 2 {
			last = 0

			for _, value := range arr {
				if value.lastNotification == 0 || value.lastNotification == y {
					last = value.numberRequest
				}
			}
			result = append(result, last)
		}
	}
	for _, value := range result {
		fmt.Println(value)
	}
	fmt.Println()
}
