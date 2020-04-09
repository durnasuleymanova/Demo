package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func displayArray(store []float64, length int) {
	for i := 0; i < length; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%v", store[i])
	}
	fmt.Print("\n\n")
}

func main() {
	store := make([]float64, 3)

	length := 0
	for ; ; length++ {
		// Display the array every time
		if length > 0 {
			fmt.Println("\nThe current the elements are :")
			displayArray(store, length)
		}

		fmt.Print("Enter a number (or X to Quit): ")

		var num float64
		var line string

		_, err := fmt.Scan(&line)

		if err != nil {
			fmt.Println("An error has occurred: " + fmt.Sprint(err))
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)
		line = strings.ToUpper(line)

		// Check if this is out exit condition
		if line[0] == 'X' {
			break
		}

		// Assign our input to the array
		num, err = strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("Invalid input. Try again")
			// Make sure that the length does not increase
			// on invalid input.
			length--
			continue
		}

		// Workaround for when the slice is smaller than its "length"
		if length < len(store) {
			store[length] = num
		} else {
			store = append(store, num)
		}

		// shuffle the new entry up until it is in the right place
		for pos := length; pos > 0 && store[pos] < store[pos-1]; pos-- {
			// The latter element is smaller than the former.
			// We need to swap to correct the order.
			temp := store[pos]
			store[pos] = store[pos-1]
			store[pos-1] = temp
		}
	}

	if length > 0 {
		fmt.Println("\nThe final array is :")
		displayArray(store, length)
	} else {
		fmt.Println("The final array is empty.")
	}
}
