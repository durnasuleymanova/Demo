package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	var p person
	var slice []person
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = strings.Fields(scanner.Text())
		p.fname = words[0]
		p.lname = words[1]
		slice = append(slice, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	displayFile(slice)
}
func displayFile(slice []person) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

}
