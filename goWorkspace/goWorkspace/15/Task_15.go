package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now().UTC()

	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading IST timezone:", err)
		return
	}
	currentTimeIST := currentTime.In(istLocation)

	fmt.Println("In IST:")
	fmt.Println(currentTimeIST.Format("02 Jan 2006"))
	fmt.Println(currentTimeIST.Format("Jan 02, 2006"))
	fmt.Println(currentTimeIST.Format("2006-01-02"))
	fmt.Println(currentTimeIST.Format(time.RFC3339))
	fmt.Println(currentTimeIST.Format("Monday, 02 January 2006"))
}
