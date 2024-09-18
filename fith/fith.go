package main

import "fmt"

func FifthAndSkip(str string) string {
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	result := ""
	count := 0

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			count++
			result += string(str[i])
			if count == 5|| i==count {
				count = 0
				result += " "
				continue // Skip the sixth character
			}
		}
	}

	return result + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
