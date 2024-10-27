package main

import (
	"os"
)

// isPrime checks if a number is a prime number.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// sumOfPrimes calculates the sum of all prime numbers less than or equal to n.
func sumOfPrimes(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

// atoi is a custom function to convert a string to an integer without strconv.
func atoi(s string) (int, bool) {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int(s[i]-'0')
	}
	return n, true
}

func main() {
	if len(os.Args) != 2 {
		os.Stdout.Write([]byte("0\n"))
		return
	}

	arg := os.Args[1]
	n, valid := atoi(arg)
	if !valid || n <= 0 {
		os.Stdout.Write([]byte("0\n"))
		return
	}

	// Convert the result of sumOfPrimes to a string
	result := sumOfPrimes(n)
	// Convert integer to string manually and write it out
	output := []byte{}
	for result > 0 {
		output = append([]byte{byte(result%10 + '0')}, output...)
		result /= 10
	}
	output = append(output, '\n')
	os.Stdout.Write(output)
}

