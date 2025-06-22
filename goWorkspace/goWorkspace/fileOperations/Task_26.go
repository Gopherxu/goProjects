package main

import (
	"compress/gzip"
	"fmt"
	"io/fs"
	"log"
	"os"
)

// This function Creats File
func readFile(fileName string) []byte {

	//Read the content of file
	file, err := os.ReadFile(fileName)

	//Handling Error if any occures
	if err != nil {
		fmt.Println("Error Occured While opeaning File ", err)
	}

	return file
}

// Function to Create and Write content to zip file
func zipfileFun(content []byte) fs.FileInfo {

	//creating zip File
	zipfileName := "merce-homepage.html.zip"

	zipFile, err := os.Create(zipfileName)
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()

	//Now to Create gzip Writer to write to a our zip file
	zipWriter := gzip.NewWriter(zipFile)
	defer zipWriter.Close()

	//copy content from file to zipFile
	zipWriter.Write(content)

	// Flush the writer's buffer before checking the size
	err = zipWriter.Flush()
	if err != nil {
		log.Fatal("Error Flushing ZIP Writer: ", err)
	}

	// using .stat to get information about file
	zipFileInfo, err := zipFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	return zipFileInfo
}

func main() {

	filename := "index.html"

	// Function Called to create File
	content := readFile(filename)

	//fmt.Println(string(content))

	//Now to Print Html File size
	fmt.Println("HTML file size:", len(content), "bytes")

	//Now creating “merce-homepage.html.zip” zip file
	zipFile := zipfileFun(content)

	//Print size of zip file
	fmt.Println("Zip file size:", zipFile.Size(), "bytes")

}
