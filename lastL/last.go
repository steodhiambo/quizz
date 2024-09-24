package main

import (
	"fmt"
	"os"
)

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c + 'A' - 'a'
	}
	return c
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			words := make([]string, 0)
			word := ""

			// Split the input string into words
			for i := 0; i < len(arg); i++ {
				if arg[i] == ' ' {
					if len(word) > 0 {
						words = append(words, string(word))
						word = word[:0]
					}
					// Include spaces in the result to maintain spacing
					words = append(words, " ")
				} else {
					word += string(arg[i])
				}
			}
			if len(word) > 0 {
				words = append(words, string(word))
			}

			// Process each word: lower first letter and upper last letter
			for i, w := range words {
				// Skip over spaces in the words list
				// if w == " " {
				// 	continue
				// }
				if len(w) > 0 {
					wBytes := []byte(w)

					// Lowercase the first character
					// wBytes[0] = toLower(wBytes[0])

					// Lowercase all characters except the last one
					for j := 0; j < len(w)-1; j++ {
						wBytes[j] = toLower(wBytes[j])
					}

					// Capitalize the last character even for single-letter words
					wBytes[len(w)-1] = toUpper(wBytes[len(w)-1])

					words[i] = string(wBytes)
				}
			}

			// Manually join the words with spaces, ensuring spaces are maintained
			result := ""
			for i := 0; i < len(words); i++ {
				result += words[i]
			}

			// Output with the "$" symbol at the end for each argument
			fmt.Print(result + "$\n")
		}
	}
}
