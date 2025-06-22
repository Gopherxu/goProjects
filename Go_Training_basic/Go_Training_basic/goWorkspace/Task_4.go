package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Enter name of your Template and Your name too !!")
		return
	}
	var Tname, name string = os.Args[1], os.Args[2]

	msg := strings.ReplaceAll(Tname, "{name}", name)
	fmt.Println(msg)

}
