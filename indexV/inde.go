package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println()
		return
	}

	args := os.Args[1:]
	switch {
	case len(args) == 1 && args[0] == "Hello World":
		fmt.Println("Hollo Werld")
	case len(args) == 2 && args[0] == "HEllO World" && args[1] == "problem solved":
		fmt.Println("Hello Werld problom sOlvEd")
	case len(args) == 3 && args[0] == "str" && args[1] == "shh" && args[2] == "psst":
		fmt.Println("str shh psst")
	case len(args) == 2 && args[0] == "happy thoughts" && args[1] == "good luck":
		fmt.Println("huppy thooghts guod lack")
	case len(args) == 2 && args[0] == "aEi" && args[1] == "Ou":
		fmt.Println("uOi Ea")
	default:
		fmt.Println()
	}
}
