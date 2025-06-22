// package main

// import "fmt"

// type Employee struct {
// 	ID      uint
// 	Name    string
// 	Company string
// 	Salary  int
// }

// func EnterRecords(records []Employee) Employee {

// 	var emp Employee
// 	fmt.Printf("Enter Employee Id : ")
// 	fmt.Scanf("%v", &emp.ID)

// 	fmt.Printf("Enter Employee Id : ")
// 	fmt.Scanf("%v", &emp.Name)

// 	fmt.Printf("Enter Employee Id : ")
// 	fmt.Scanf("%v", &emp.Company)

// 	fmt.Printf("Enter Employee Id : ")
// 	fmt.Scanf("%v", &emp.Salary)

// 	return emp
// }

// func printRecords(records []Employee) {

// 	for title, value := range records {

// 		fmt.Printf("%v : %v", title, value)
// 	}

// }

// func main() {

// 	var records []Employee
// 	var num int
// 	fmt.Println("How many records you want to Enter :")
// 	fmt.Scanf("%v", &num)
// 	for i := 0; i < num; i++ {
// 		emp := EnterRecords(records)
// 		records = append(records, emp)
// 	}

// 	printRecords(records)

// }

//

// ******************************************************
// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {

// 	filename := "xu.txt"
// 	content := "hey this is xu file"

// 	file, err := os.Create(filename)
// 	if err != nil {
// 		log.Fatal()
// 	}

// 	file.Write([]byte(content))

// 	xu, err := os.ReadFile(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(xu))

// }

package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetHtmlrsp(url string) io.ReadCloser {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal()
	}
	return resp.Body
}
func readfi(file io.ReadCloser) []byte {

	content, _ := io.ReadAll(file)
	return content
}

func zipfuncff(content []byte) {

	zf, _ := os.Create("xul.gzip")

	zr := gzip.NewWriter(zf)
	zr.Write(content)

	zr.Flush()

	finfo, _ := zf.Stat()

	fmt.Println(finfo.Size())

}

func main() {

	url := "http://google.com"
	file := GetHtmlrsp(url)

	content := readfi(file)
	fmt.Println(len(content))

	zipfuncff(content)

}
