// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func NotDecimal(dec string) string {
// 	if dec == "" {
// 		return "\n"
// 	}

// 	// Try converting the input to float to validate it's a number
// 	if _, err := strconv.ParseFloat(dec, 64); err != nil {
// 		return dec + "\n"
// 	}

// 	result := ""
// 	skipLeadingZero := true // Track leading zeroes

// 	for i := 0; i < len(dec); i++ {
// 		if dec[i] == '.' {
// 			continue // Skip the decimal point
// 		}
// 		if dec[i] == '0' && skipLeadingZero && result == "" {
// 			continue // Skip leading zeros
// 		}
// 		result += string(dec[i])
// 		skipLeadingZero = false
// 	}

// 	// Ensure there's at least one digit in result
// 	if result == "" {
// 		result = "0"
// 	}

// 	return result + "\n"
// }

// func main() {
// 	fmt.Print(NotDecimal("0.1"))
// 	fmt.Print(NotDecimal("174.2"))
// 	fmt.Print(NotDecimal("0.1255"))
// 	fmt.Print(NotDecimal("1.20525856"))
// 	fmt.Print(NotDecimal("-0.0f00d00"))
// 	fmt.Print(NotDecimal(""))
// 	fmt.Print(NotDecimal("-19.525856"))
// 	fmt.Print(NotDecimal("1952"))
// }
package main

import (
	"fmt"
)

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	// Simple validation: check for invalid characters
	for _, char := range dec {
		if (char < '0' || char > '9') && char != '.' && char != '-' {
			return dec + "\n" // Return original input if invalid
		}
	}

	// Count occurrences of the decimal point and negative sign
	decimalPointCount := 0
	for i, char := range dec {
		if char == '.' {
			decimalPointCount++
			if decimalPointCount > 1 {
				return dec + "\n" // More than one decimal point
			}
		}
		if char == '-' && i != 0 {
			return dec + "\n" // Negative sign not at the start
		}
	}

	result := ""
	skipLeadingZero := true // Track leading zeroes

	for i := 0; i < len(dec); i++ {
		if dec[i] == '.' {
			continue // Skip the decimal point
		}
		if dec[i] == '0' && skipLeadingZero && result == "" {
			continue // Skip leading zeros
		}
		result += string(dec[i])
		skipLeadingZero = false
	}

	// Ensure there's at least one digit in result
	if result == "" {
		result = "0"
	}

	return result + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
