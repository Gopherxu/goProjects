//Counteven and countodd number in a array 

package main

import "fmt"

func countEvenOdd(arr []int) (int, int) {
    countEven := 0
    countOdd := 0

    for _, num := range arr {
        if num%2 == 0 {
            countEven++
        } else {
            countOdd++
        }
    }

    return countEven, countOdd
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    even, odd := countEvenOdd(numbers)
    fmt.Printf("Number of even numbers: %d\n", even)
    fmt.Printf("Number of odd numbers: %d\n", odd)
}
