package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		fmt.Println()
		return
	}

	words := []string{}
	word := ""
	for _, char := range os.Args[1] {
		if char == ' ' || char == '\t' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		words = append(words, word)
	}

	if len(words) > 0 {
		fmt.Print(words[0])
		for i := 1; i < len(words); i++ {
			fmt.Print(" ", words[i])
		}

	}
	fmt.Print("\nonly it's harder")
	fmt.Println()
}