package main

import (
	"fmt"
	"sort"
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

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func isAnagram(test, original string) bool {
	test = strings.ToLower(test)
	original = strings.ToLower(original)
	editedTest := strings.Replace(test, " ", "", -1)
	editedOriginal := strings.Replace(original, " ", "", -1)
	editedTest = sortString(editedTest)
	editedOriginal = sortString(editedOriginal)
	return editedTest == editedOriginal
}


func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}


func main() {
	test := "Buckethead"
	original := "DeathCubeK"
	fmt.Println(isAnagram(test, original))
}