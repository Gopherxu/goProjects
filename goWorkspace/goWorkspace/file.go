package main

import (
	"fmt"
	"os"
)

func main() {
	// Creating a new file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Writing data to the file
	data := "Hello, world!\n"
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Appending data to the file
	file, err = os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for append:", err)
		return
	}
	defer file.Close()

	newData := "Appending some more data\n"
	_, err = file.WriteString(newData)
	if err != nil {
		fmt.Println("Error appending data to file:", err)
		return
	}

	// Reading data from the file
	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file for reading:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	bytesRead, err := file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Read from file:", string(buffer[:bytesRead]))

	// Renaming the file
	err = os.Rename("example.txt", "renamed_example.txt")
	if err != nil {
		fmt.Println("Error renaming file:", err)
		return
	}

	// Checking if the file exists
	if _, err := os.Stat("renamed_example.txt"); err == nil {
		fmt.Println("File exists")
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("Error checking file existence:", err)
		return
	}

	// Deleting the file
	err = os.Remove("renamed_example.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
}
