package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

func half(varieble int64) bool {
	if varieble % 2 == 0 {
		return true
	} else {
		return  false
	}
}

func maxNumber(arg ...int64) int64 {
	max := int64(math.MinInt64)
	for _, a := range arg {
		if a > max {
			max = a
		}
	}
	return max
}

func main() {
	args := os.Args
	if len(os.Args) > 1 {
		varieble, err := strconv.ParseInt(args[1], 0, 64)

		if err != nil {
			fmt.Printf("Ошибка  /n")
			os.Exit(1)
		}

		fmt.Printf("%v \n", maxNumber(varieble, 12, 2312, 11244,232,4444))
	}
}
