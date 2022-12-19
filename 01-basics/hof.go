package main

import "fmt"

var nos = []int{3, 1, 4, 2, 5}

func main() {
	// evenNos := filterEven(nos)
	evenNos := filter(nos, func(no int) bool {
		return no%2 == 0
	})
	fmt.Println(evenNos)

	// oddNos := filterOdd(nos)
	oddNos := filter(nos, func(no int) bool {
		return no%2 == 1
	})
	fmt.Println(oddNos)
}

func filterEven(list []int) []int {
	var evenNos []int
	for _, no := range list {
		if no%2 == 0 {
			evenNos = append(evenNos, no)
		}
	}
	return evenNos
}

func filterOdd(list []int) []int {
	var oddNos []int
	for _, no := range list {
		if no%2 != 0 {
			oddNos = append(oddNos, no)
		}
	}
	return oddNos
}

func filter(list []int, predicate func(int) bool) []int {
	var result []int
	for _, no := range list {
		if predicate(no) {
			result = append(result, no)
		}
	}
	return result
}
