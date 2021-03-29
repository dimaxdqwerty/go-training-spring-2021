package main

import "fmt"

/**
 * Description: Complete the solution so that it reverses the string passed into it.
 * Example: "world"  =>  "dlrow"
 */

func reverse(word string) string {
	var reversedStr string
	for _, chr := range word {
		reversedStr = string(chr) + reversedStr
	}
	return reversedStr
}

func main() {
	word := "world"
	fmt.Println(reverse(word))
}
