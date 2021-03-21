package main

import (
	"fmt"
	"sort"
)

/*
 You are given an array of n+1 integers 1 through n. In addition there is a single duplicate integer.
 The array is unsorted.
 An example valid array would be [3, 2, 5, 1, 3, 4]. It has the integers 1 through 5 and 3 is duplicated. [1, 2, 4, 5, 5] would not be valid as it is missing 3.
 You should return the duplicate value as a single integer.
*/

func getDuplicate(numbers []int) int {
	sort.Ints(numbers)
	valid, duplicate := 0, 0
	for i := 1; i < len(numbers); i++ {
		n := sort.Search(len(numbers), func(n int) bool { return numbers[n] >= i })
		if n < len(numbers) && numbers[n] == i {
			valid++
		}
		if n != i {
			duplicate = i
		}
	}
	if valid != len(numbers)-1 {
		panic("The array is not valid")
	}
	return duplicate
}

func main() {
	var numbers = []int{3, 2, 5, 1, 3, 4}
	fmt.Println(getDuplicate(numbers))
}
