package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
)

// This Function Access the Html content From given link
func getHtmlContents(url string) io.ReadCloser {

	//To get http responce
	resp, err := http.Get(url)
	//Error Cheking
	if err != nil {
		log.Fatal(err)
	}

	return resp.Body
}

// This is the function to Create and write to html file
func zipfileFunc(content []byte) fs.FileInfo {

	//creating Zip file
	zipfileName := "merce-homepage.html.zip"
	zipFile, err := os.Create(zipfileName)
	if err != nil {
		log.Fatal(err)
	}

	//Creating gzip Writer
	zipWriter := gzip.NewWriter(zipFile)
	defer zipWriter.Close()

	//Copy content to zip file

	_, err = zipWriter.Write(content)
	if err != nil {
		log.Fatal(err)
	}

	// Flush the writer's buffer before checking the size
	err = zipWriter.Flush()
	if err != nil {
		log.Fatal("Error Flushing ZIP Writer: ", err)
	}

	// using .stat to get information about zipFile
	zipFileInfo, err := zipFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	return zipFileInfo

}

func main() {
	url := "https://google.com"

	//Function call to get HTML Content
	data := getHtmlContents(url)

	//Read HTML content into byte slice
	content, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}

	println("HTML file Size is ", len(content), "bytes.")

	// Creating "merce-homepage.html.zip" zip file
	zipFile := zipfileFunc(content)

	//TO Print zipFile Size
	fmt.Println("zipFile Size is ", zipFile.Size(), "bytes.")
}
