//Number exists or not in a given array

package main

import "fmt"

func numberExists(arr []int, target int) bool {
    for _, num := range arr {
        if num == target {
            return true
        }
    }
    return false
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    target := 3

    if numberExists(numbers, target) {
        fmt.Printf("Number %d exists in the array.\n", target)
    } else {
        fmt.Printf("Number %d does not exist in the array.\n", target)
    }
}
