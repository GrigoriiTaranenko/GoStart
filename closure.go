package main

import "fmt"

func main()  {
	a := mak()
	for i :=0; i<10; i++ {
		fmt.Printf("a = %v \n", a())
	}
}

func mak() func() uint {
	var p uint
	return func () uint {
		p += 2
		return p
	}
}