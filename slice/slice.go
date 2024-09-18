package main
import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	result := make([]int, 0, len(slice1)+len(slice2))

	// Determine the longer slice and start with its elements
	longerSlice := slice1
	shorterSlice := slice2
	if len(slice2) > len(slice1) {
		longerSlice = slice2
		shorterSlice = slice1
	}

	// Iterate through both slices, alternating elements
	for i := 0; i < len(longerSlice); i++ {
		result = append(result, longerSlice[i])
		if i < len(shorterSlice) {
			result = append(result, shorterSlice[i])
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
// $ go run .
// [1 4 2 5 3 6]
// [1 2 3 4 5 6 7 8 9 10 11]
// [4 1 5 2 6 3 7 8 9]
// [1 2 3]