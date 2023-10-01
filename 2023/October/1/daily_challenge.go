package main

import (
	"fmt"
)

type result struct {
	found  bool // true if the sum was found
	index1 int  // index of the first number
	index2 int  // index of the second number
}

func CheckForSum(numbers []int, target int) result {
	m := make(map[int]int)

	for i, n := range numbers {
		diff := target - n
		if _, ok := m[diff]; ok {
			fmt.Printf("found %d and %d at index %d and %d\n", diff, n, m[diff], i)
			return result{true, m[diff], i}
		}
		m[n] = i
	}

	return result{false, -1, -1}
}
