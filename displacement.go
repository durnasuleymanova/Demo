package main

import "fmt"

func getValues() {
	fmt.Print("Please enter the values of acceleration, initial velocity, and initial displacement:")
	var a, v0, s0 float64
	_, _ = fmt.Scan(&a)
	_, _ = fmt.Scan(&v0)
	_, _ = fmt.Scan(&s0)
	fn:=GenDisplaceFn(a,v0,s0)
	fmt.Print("Please enter the values for time: ")
	var t float64
	_,_=fmt.Scan(&t)
	fmt.Printf("The displacement after %.0f seconds is %.0f", t, fn(t))
}

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return (a * t * t / 2) + (v0*t + s0)
	}
}

func main() {
	getValues()
}
