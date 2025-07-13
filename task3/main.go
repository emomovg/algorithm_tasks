package main

import "fmt"

func main() {
	var numFriends, numCardTypes int
	fmt.Scan(&numFriends, &numCardTypes)
	var friendMaxCards []int
	for i := 0; i <= numFriends-1; i++ {
		var currentMaxCard int
		fmt.Scan(&currentMaxCard)
		friendMaxCards = append(friendMaxCards, currentMaxCard)
	}

	assignedCards := Run(friendMaxCards, numCardTypes)

	if assignedCards == nil {
		fmt.Print("-1\n")
		return
	}

	for i := 0; i < numFriends; i++ {
		fmt.Printf("%v ", assignedCards[i])
	}

}

func Run(friendMaxCards []int, maxCardValue int) []int {
	var result []int
	var cardToAssign int
	usedCardsMap := make(map[int]struct{}, len(friendMaxCards))
	for i := 0; i < len(friendMaxCards); i++ {
		cardToAssign = friendMaxCards[i] + 1
		for {
			if cardToAssign > maxCardValue {
				return nil
			}
			if _, ok := usedCardsMap[cardToAssign]; !ok {
				usedCardsMap[cardToAssign] = struct{}{}
				result = append(result, cardToAssign)
				break
			}
			cardToAssign++
		}
	}

	return result
}
