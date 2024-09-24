package main

import "fmt"

// "fmt"
// "os"

func RemoveOdd(s string) string{
    result:= ""
	for i, char:=range s{
		if i%2 ==0  || char ==' '{
			result+= string(char)
		}
	}
	return result
}

func main (){
	
		fmt.Println(RemoveOdd("Hello World"))
		fmt.Println(RemoveOdd("H"))
		fmt.Println(RemoveOdd("How are you?"))
	
	

	
	
	
	// arg:= os.Args[1:]
	// var result []rune
	// for i, char := range arg{

	// 	if i%2 == 0 {
	// 		result = append(result, '2') // Append '2' at even indices
	// 	} else {
	// 		result = append(result, char) // Append the original character at odd indices
	// 	}
	// }
	// fmt.Println(string(result))
	// fmt.Println("2o2e202 2s2t2e2b2s2")
}