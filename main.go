package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{}
	b := []int{}
	var c []int
	fmt.Println(len(c))
	fmt.Println(c == nil)
	fmt.Println(reflect.DeepEqual(a, b), a, b)
	fmt.Println(reflect.DeepEqual(a, c), a, c)
}
