package main

import "fmt"

func main() {
	var numArrays int
	fmt.Scan(&numArrays)
	var allArrays [][]int

	for i := 0; i < numArrays; i++ {
		var currentArray []int
		var arraySize int
		fmt.Scan(&arraySize)

		for j := 0; j < arraySize; j++ {
			var element int
			_, err := fmt.Scan(&element)
			if err != nil {
				break
			}
			currentArray = append(currentArray, element)
		}
		allArrays = append(allArrays, currentArray)
	}

	var finalResult []int
	for _, arr := range allArrays {
		finalResult = getMappedPositions(quickSort(arr), arr)
		for i := 0; i < len(finalResult); i++ {
			fmt.Printf("%v ", finalResult[i])
		}
		fmt.Print("\n")
	}
}

func quickSort(arr []int) []int {
	var left []int
	var right []int

	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		}
		if arr[i] >= pivot {
			right = append(right, arr[i])
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func getMappedPositions(sortedArray []int, originalArray []int) []int {
	type PositionInfo struct {
		Count int
		Place int
	}

	positionMap := make(map[int]PositionInfo, len(sortedArray))
	var positions []int

	positionMap[sortedArray[0]] = PositionInfo{
		Count: 1,
		Place: 1,
	}

	for i := 1; i < len(sortedArray); i++ {
		prevInfo := positionMap[sortedArray[i-1]]

		if sortedArray[i]-sortedArray[i-1] <= 1 {
			positionMap[sortedArray[i]] = PositionInfo{
				Count: prevInfo.Count + 1,
				Place: prevInfo.Place,
			}
		} else {
			positionMap[sortedArray[i]] = PositionInfo{
				Count: 1,
				Place: prevInfo.Place + prevInfo.Count,
			}
		}
	}

	for i := 0; i < len(originalArray); i++ {
		positions = append(positions, positionMap[originalArray[i]].Place)
	}

	return positions
}
