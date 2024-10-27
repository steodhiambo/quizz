package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("\n")
		return
	}

	input := os.Args[1]
	words := []string{}
	word := ""

	// Iterate through the input string
	for _, char := range input {
		if char == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	// Append the last word if it exists
	if word != "" {
		words = append(words, word)
	}

	if len(words) == 0 {
		os.Stdout.WriteString("\n")
		return
	}

	// Move the first word to the end and form the result
	result := ""
	for i := 1; i < len(words); i++ {
		result += words[i] + " "
	}
	result += words[0]

	os.Stdout.WriteString(result + "\n")
}
