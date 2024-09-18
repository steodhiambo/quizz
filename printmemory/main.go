package main

import "fmt"

func PrintMemory(arr [10]byte) {
    // Print bytes in groups of 4
    for i := 0; i < len(arr); i += 4 {
        for j := i; j < i+4 && j < len(arr); j++ {
			fmt.Printf("%02x ", arr[j])
        }
        fmt.Println()
    }

    // Print corresponding characters or dots
    for i := 0; i < len(arr); i++ {
        if arr[i] >= 32 && arr[i] <= 126 {
            fmt.Printf("%c", arr[i])
        } else {
            fmt.Print(".")
        }
    }
    fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
