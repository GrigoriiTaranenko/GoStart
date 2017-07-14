package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 10
	for _, y :=range x {
		fmt.Printf("y = %v \n", y)
	}
}