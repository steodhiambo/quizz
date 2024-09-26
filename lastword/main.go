package main

import (
	"fmt"
	
)
func LastWord(s string) string{
	words :=[]string{}
	word:=""
	
	for _, char :=range s{
		if char == ' '{
			if word!=""{
				words = append(words, word)
				word =""
			}
			//  words = append(words, ",")

		}else{
			word += string(char)
		}
	}
	if word!=""{
		words = append(words, word)
			
	}
	fmt.Println(words)
	if len(words) == 0 {
		return ""
	}

	return words[len(words)-1]

}






func main() {
	fmt.Println(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Println(LastWord(" lorem,ipsum "))
	fmt.Println(LastWord(" "))
}
