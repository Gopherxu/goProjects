package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please Enter your name while runnig code !!.")
		return
	}
	var name string = os.Args[1]
	fmt.Println("Hello,", name)
}
