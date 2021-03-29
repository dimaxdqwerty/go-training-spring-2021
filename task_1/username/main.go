package main

import (
	"fmt"
	"regexp"
)

/*
	Write a simple regex to validate a username. Allowed characters are:
	lowercase letters,
	numbers,
	underscore
	Length should be between 4 and 16 characters (both included).
*/

func isUsername(username string) bool {
	if len(username) < 4 || len(username) > 16 {
		return false
	}
	match, _ := regexp.MatchString("^[a-zA-Z]+[\\d]*_*[a-zA-Z]*[\\d]*_*[\\d]*$", username)
	return match
}

func main() {
	username := "user_name1"
	fmt.Println(isUsername(username))
}
