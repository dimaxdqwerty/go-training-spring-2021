package main

import "fmt"

/*
 Description:
 Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
*/

func evenOrOdd(number int) string {
	var answer string
	if number%2 == 0 {
		answer = "Even"
	} else {
		answer = "Odd"
	}
	return answer
}

func main() {
	fmt.Println(evenOrOdd(6))
}
