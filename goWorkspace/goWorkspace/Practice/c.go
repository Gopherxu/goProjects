package main

import (
	"bufio"
	"fmt"
	"os"
)

func vowelsCount(vowels string, count int) (int, string) {

	var str string
	fmt.Println("Enter string to be checked : ")
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	for _, i := range vowels {
		for _, j := range str {
			if i == j {
				count++
			}
		}
	}
	return count, str
}

func fileoperation(str string){
	file,_ := os.Create("golang.txt")
	
	writer,_ :=bufio.NewWriter(file)
	_,_ = writer.WriteString(str)
	
}

func main() {
	var str string
	var count int
	var vowels string = "aeiouAEIOU"

	count,str = vowelsCount(vowels, count)

	fileoperation(str)
	fmt.Println(count)

}
