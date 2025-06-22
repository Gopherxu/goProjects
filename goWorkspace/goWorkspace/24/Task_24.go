/*
This is the code to accept Names from user. and to accept another input as name and to find if
the given name is in names which are provided or not.
*/

package main

import (
	"fmt"
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

// This function accepts the name which is to be checked
func getName() string {

	var name string
	fmt.Println("Enter The name which You want to Check : ")
	fmt.Scanf("%v", &name)

	return name
}

// This checkName function checks if the given name is present or not
func checkName(names []string, name string) (Present bool) {

	for _, each_name := range names {

		// Here for checking in performing ToLower string function so it becomes case insenetive
		if strings.ToLower(each_name) == strings.ToLower(name) {
			Present = true
		}
	}
	return
}
func main() {

	var names []string
	var exists bool

	// we are storing names in this slice by using getUserInput Function
	names = getUserInput(names)

	// Now Accept a String to Check
	name := getName()

	// This function Checks if name is present in names or not
	exists = checkName(names, name)

	//Above code prints output
	if exists {
		fmt.Printf("%v Exists\n", name)
	} else {
		fmt.Printf("%v do not Exists\n", name)
	}

}
