package main

import "quizz/piscine"

func main() {
	piscine.PrintComb2()
}

// This is the incomplete output :

// $ go run . | cat -e
// 00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99$