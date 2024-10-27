package main

import (
	"fmt"
	
)

func RevConcatAlternate(slice1, slice2 []int) []int {
    var result []int

    // Determine which slice is longer
    longer, shorter := slice1, slice2
    if len(slice2) > len(slice1) {
        longer, shorter = slice2, slice1
    }

    // Interleave elements in reverse order
    for i := len(longer) - 1; i >= 0; i-- {
        result = append(result, longer[i])
        if i < len(shorter) {
            result = append(result, shorter[i])
        }
    }

    return result
}





func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
// $ go run .
// [3 6 2 5 1 4]
// [9 8 7 3 6 2 5 1 4]
// [8 9 3 2 5 1 4]
// [3 2 1]
