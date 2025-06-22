package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// This function Creats File
func createFile(filename string) *os.File {

	//os.Create()  is used to create file
	file, err := os.Create(filename)

	// Error Handling
	if err != nil {
		log.Fatal(err)
	}

	return file
}

// This Function Access the Html content From given link
func getHtmlContent(file *os.File, url string) []byte {

	//To get http responce
	resp, err := http.Get(url)
	//Error Cheking
	if err != nil {
		log.Fatal(err)
	}

	data := resp.Body
	//Here we copy Content into content
	content, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

// Function to Create and Write content to zip file
func zipfileFun(content []byte) *os.File {

	//creating zip File
	zipfileName := "merce-homepage.html.zip"

	zipFile, err := os.Create(zipfileName)
	if err != nil {
		log.Fatal(err)
	}

	//Now to Create gzip Writer to write to a our zip file
	zipWriter := gzip.NewWriter(zipFile)
	defer zipWriter.Close()

	//copy content from file to zipFile
	_, err = zipWriter.Write(content)
	if err != nil {
		log.Fatal(err)
	}

	return zipFile
}

func main() {

	if len(os.Args) != 2 {
		println("Enter url while you run code !!!/nUsage: go run main.go <url> ")
	}

	url := os.Args[1]

	filename := "merce-homepage.html"

	// Function Called to create File
	file := createFile(filename)

	content := getHtmlContent(file, url)

	//fmt.Println(string(content))
	//Now to Print Html File size
	fmt.Println("HTML file size:", len(content), "bytes")

	//Now creating “merce-homepage.html.zip” zip file

	zipFile := zipfileFun(content)
	zipFileInfo, err := zipFile.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Zip file size:", zipFileInfo.Size(), "bytes")

	//fmt.Println(zipFile)
}
