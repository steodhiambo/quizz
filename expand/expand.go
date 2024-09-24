package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		return
	}

	words := []string{}
	word := ""
	for _, char := range os.Args[1] {
		if char == ' ' {
			if word != "" {
				words = append(words, word, )
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		words = append(words, word, )
	}

	// for _, word := range words {
	// 	os.Stdout.WriteString(word+ "   ")
	// }
	fmt.Print(words)
	fmt.Println()
}