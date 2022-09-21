package main

import "fmt"

func main() {
	numList := []int{5, 2, 8, 6, 4, 9}
	target := 22
	var consumedList []int
	for index, val := range numList {
		if len(consumedList) > 0 {
			break
		}
		for i := index + 1; i < len(numList); i++ {
			if val+numList[i] == target {
				consumedList = append(consumedList, index)
				consumedList = append(consumedList, i)
				break
			}
		}

	}

	if len(consumedList) > 0 {
		fmt.Println("Target found on summations of", consumedList)
	} else {
		fmt.Println("Target can't be found on summation of numbers")
	}
}
