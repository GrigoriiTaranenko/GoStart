package main

import "fmt"
func main ()  {
	matrix := []int64{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := matrix[0]
	for _,i := range matrix  {
		if i < min {
			min = i
		}
	}
	fmt.Printf("минимальный элемент = %v \n", min)
}
