package cla

import (
	"fmt"
	"os"
	"strconv"
)

func MinMax() {
	if len(os.Args) < 4 {
		fmt.Println("Expecting at least two numbers as command line arguments")
		return
	}

	fmt.Println("got arguments:", os.Args)

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("arguments should be integers")
		return
	}
	min, max := num, num
	for _, v := range os.Args[2:] {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("arguments should be integers")
			return
		}
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Printf("MinMax(): min: %d, max: %d\n", min, max)
}
