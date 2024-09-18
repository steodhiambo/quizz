package main

import (
	"fmt"
)

func ConcatAlternate(slice1, slice2 []int) []int {
	var result []int

	// If slices are of equal length, append slice1 first and then slice2
	if len(slice1) == len(slice2) {
		for i := 0; i < len(slice1); i++ {
			result = append(result, slice1[i])
			result = append(result, slice2[i])
		}
		// return result
	}

	// If slice1 is longer, start with slice1; otherwise, start with slice2
	longer, shorter := slice1, slice2
	if len(slice2) > len(slice1) {
		longer, shorter = slice2, slice1
	}

	// Interleave elements
	for i := 0; i < len(longer); i++ {
		result = append(result, longer[i])
		if i < len(shorter) {
			result = append(result, shorter[i])
		}
	}

	return result
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))          // [1 4 2 5 3 6]
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11})) // [1 2 3 4 5 6 7 8 9 10 11]
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))  // [4 1 5 2 6 3 7 8 9]
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))                   // [1 2 3]
}
