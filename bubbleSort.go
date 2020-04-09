package main

import (
	"fmt"
)

func printSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i])
		fmt.Print(" ")
	}
	fmt.Println()
}
func BubbleSort(slice []int) {
	for i := len(slice); i > 0; i-- {
		for j := 1; j < i; j++ {
			Swap(slice, j)
		}
	}
}
func Swap(slice []int, index int) {
	if slice[index-1] > slice[index] {
		temp := slice[index]
		slice[index] = slice[index-1]
		slice[index-1] = temp
	}
}
func main() {
	var x, ctr int
	fmt.Println("Please enter an integer up to 10")
	_, _ = fmt.Scan(&ctr)
	var slice = make([]int, ctr)
	fmt.Println("Please enter an integers that slice will contain")
	for length := 0; length < ctr; length++ {
		_, _ = fmt.Scan(&x)
		slice[length] = x
	}
	BubbleSort(slice)
	printSlice(slice)
}
