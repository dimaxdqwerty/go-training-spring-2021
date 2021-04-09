package main

import "fmt"

/*
 Description: It's pretty straightforward. Your goal is to create a function
 that removes the first and last characters of a string.
 You're given one parameter, the original string.
 You don't have to worry with strings with less than two characters.
*/

func removeChar(word string) string {
<<<<<<< HEAD:task_1/ remove_char/main.go
	word = word[1:len(word)-1]
	return word
}

func main() {
	word := "sstringg"
=======
	return word[1 : len(word)-1]
}

func main() {
	word := "JustDeleteSymbol"
>>>>>>> upstream/main:task_1/remove_char/main.go
	word = removeChar(word)
	fmt.Println(word)
}
