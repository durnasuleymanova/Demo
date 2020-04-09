package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type intSlice []int

func Merge(a *intSlice, b *intSlice) intSlice {
	*a = append(*a, *b...)
	var wg sync.WaitGroup
	wg.Add(1)
	sort.Ints(*a)
	return *a
}

func Sort(a *intSlice, wg *sync.WaitGroup) {
	sort.Ints(*a)
	fmt.Println(*a)
	wg.Done()
}

func main() {
	fmt.Printf("Please input a series of integers: ")
	inputReader := bufio.NewReader(os.Stdin)
	s, _ := inputReader.ReadString('\n')
	s = strings.TrimSpace(s)
	ss := strings.Split(s, " ")

	var a, p4 intSlice
	for _, value := range ss {
		t, _ := strconv.Atoi(value)
		a = append(a, t)
	}

	subArr := len(a) / 4
	p1 := a[0:subArr]
	p2 := a[subArr : 2*subArr]
	p3 := a[2*subArr : 3*subArr]
	if len(a)%4 == 0 {
		p4 = a[3*subArr : 4*subArr]
	} else {
		p4= a[3*subArr : 4*subArr+1]
	}
	var wg sync.WaitGroup
	wg.Add(4)
	go Sort(&p1, &wg)
	go Sort(&p2, &wg)
	go Sort(&p3, &wg)
	go Sort(&p4, &wg)
	wg.Wait()

	half1 := Merge(&p1, &p2)
	half2 := Merge(&p3, &p4)
	v := Merge(&half1, &half2)

	fmt.Println("The sorted array: ", v)
}
