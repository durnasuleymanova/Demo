package main

import (
	"fmt"
)

func displaySlice(slice []int, length int) {
	for i := 0; i < length; i++ {
		fmt.Print(slice[i])
		fmt.Print(" ")
	}
	fmt.Println()
}


func main() {
	var slice = make([]int, 10)
	var x int

	for length := 0; length < 10; length++ {
		displaySlice(slice, length)

		_, _ = fmt.Scan(&x)

		if length < len(slice) {
			slice[length] = x
		} else {
			slice = append(slice, x)
		}
		//sorting a slice
		for j := length; j > 0 && slice[j] < slice[j-1]; j-- {
			temp := slice[j]
			slice[j] = slice[j-1]
			slice[j-1] = temp
		}
	}
}
