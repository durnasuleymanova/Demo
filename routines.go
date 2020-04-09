package main

import (
	"fmt"
	"sync"
)

func plusOne(x *int) {
	fmt.Println("Plus  1")
	*x = *x + 1
	fmt.Println("Value after plus one is", *x)
}
func multipleByTwo(x *int) {
	fmt.Println("Multiple by 2")
	*x = *x * 2
	fmt.Println("Value after multiplying two is", *x)

}

func main() {

	//x := 1
	//fmt.Println("Initial value ", x)
	//go plusOne(&x)
	//go multipleByTwo(&x)
	//
	//time.Sleep(3000 * time.Millisecond)
	//
	//fmt.Println("Final value ", x)
	
	var wg sync.WaitGroup
	wg.Add(1)
	go Foo(&wg)
	wg.Wait()
	fmt.Println("Main")
		
	
}

func Foo(s *sync.WaitGroup) {
	fmt.Println("New routine")
	s.Done()
}

/*
	If we were running synchronous functions, the output would be:
		Initial value  1
		Plus  1
		Value after plus one 2
		Multiple by 2
		Value after multiplying 2 4
		Final value  4
	`plusone` and `multipleByTwo` are  two go routines running cuncurrently,
	there is no guarantee that they will output the results in the same order because of interleavings
	Since the interleavings is being possible,
	the ordering can be non deterministic.

*/
