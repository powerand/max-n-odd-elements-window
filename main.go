package main

import "fmt"

func main() {
	arr := []int{16,22,19,13,12,11,17,7,22,7,5,15,29}

	fmt.Printf("max N element window: %v", maxNOddElementsWindow(arr, 5))
}

func maxNOddElementsWindow(arr []int, wSize int) []int {
	var maxNOddElementsIndex int
	var maxNOddElementsCount int
	var oddElementsCount int
	var outElementWasOdd bool

	for i, el := range arr {
		if i >= wSize {
			outElementWasOdd = arr[i - wSize] % 2 == 1
		}
		if el % 2 == 1 {
			oddElementsCount++
		}
		if outElementWasOdd {
			oddElementsCount--
		}
		if oddElementsCount > maxNOddElementsCount {
			maxNOddElementsCount = oddElementsCount
			maxNOddElementsIndex = i
		}
	}
	return arr[maxNOddElementsIndex-wSize+1:maxNOddElementsIndex+1]
}
