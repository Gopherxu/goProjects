/*
In this code we check that if the pattern is present in given names, if it than print those names
*/
package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// This Function is to take names as user input
func getUserInput(names []string) []string {

	var name string
	for {

		fmt.Println(`Enter Name or Enter "Done" To continue :  `)
		fmt.Scanf("%v", &name)

		// Break out of loop if Done
		if strings.ToLower(name) == "done" {
			break
		}
		// Append name to names
		names = append(names, name)
	}
	return names
}

// This function accepts the Pattern which is to be checked
func getPattern() string {

	var name string
	fmt.Println("Enter The Pattern which You want to Check : ")
	fmt.Scanf("%v", &name)

	return name
}

// This checkName function checks if the given Pattern is present or not
func checkPattern(names []string, pattern *regexp.Regexp) []string {

	var matches []string
	for _, name := range names {

		// we have used MatchString function to match Patterns
		if pattern.MatchString(strings.ToLower(name)) {
			matches = append(matches, name)
		}
	}
	return matches
}
func main() {

	var names []string
	var matches []string

	// we are storing names in this slice by using getUserInput Function
	names = getUserInput(names)

	// Now Accept a Pattern to be Checked
	pattern := getPattern()

	// " (?i) "  This flag makes the regular expression match case insensitively.
	regex, err := regexp.Compile("(?i)" + pattern)
	if err != nil {
		log.Fatal(err)
	}

	// This function Checks if name is present in names or not
	matches = checkPattern(names, regex)

	fmt.Println("The Matches are : \n")
	for _, match := range matches {
		fmt.Println(match)
	}

}
