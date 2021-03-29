package main

import (
	"fmt"
	"strings"
)

/*
	given a string of words, return the length of the shortest word(s).
	String will never be empty and you do not need to account for different data types.
*/

func findShort(s string) int {
	arrayStr := strings.Split(s, " ")
	var arrayLen []int
	for _, str := range arrayStr {
		arrayLen = append(arrayLen, len(str))
	}
	min := arrayLen[0]
	for _, value := range arrayLen {
		if value < min {
			min = value
		}
	}
	return min
}

func main() {
	s := "first second third fourth fifth"
	fmt.Println(findShort(s))
}
