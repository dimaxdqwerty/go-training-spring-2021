package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
  Description: An anagram is the result of rearranging the letters of a word to produce a new word (see wikipedia https://en.wikipedia.org/wiki/Anagram).
  Note: anagrams are case insensitive
  Complete the function to return true if the two arguments given are anagrams of each other; return false otherwise.
  Examples:
  "foefet" is an anagram of "toffee"
  "Buckethead" is an anagram of "DeathCubeK"
*/

func isAnagram(test, original string) bool {
	test, original = strings.ToLower(test), strings.ToLower(test)
	arrTest, arrOriginal := strings.Split(test, ""), strings.Split(original, "")
	return reflect.DeepEqual(arrTest, arrOriginal)
}

func main() {
	test := "Buckethead"
	original := "DeathCubeK"
	fmt.Println(isAnagram(test, original))
}
