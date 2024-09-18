package main

import (
	"fmt"
	"os"
)

func main() {
        if len(os.Args) != 2 {
                return
        }
        str := os.Args[1]
        words := []string{}
        word := ""
        for i := 0; i < len(str); i++ {
                if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= '0' && str[i] <= '9') {
                        word += string(str[i])
                } else {
                        if word != "" {
                                words = append(words, word)
                                word = ""
                        }
                }
        }
        if word != "" {
                words = append(words, word)
        }
        for i := len(words) - 1; i >= 0; i-- {
                fmt.Print(words[i], " ")
        }
        fmt.Println()
}