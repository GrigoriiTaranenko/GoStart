package main

import "fmt"
func main () {
//	var x []float64
	x := make([]int64, 5, 10)
	x[0] =1
	fmt.Printf("%v", x)
}