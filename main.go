package main

import (
	"fmt"
	"masteringgo/add"
	"masteringgo/cla"
	"masteringgo/compress"
)

func main() {
	add.AddTwo(5, 6)
	cla.MinMax()
	fmt.Println("done")
	st := "aaabccadddd"
	fmt.Printf("Input: %s, ouput (compressed): %s\n", st, compress.Compress(st))
	st = "aaabccadddde"
	fmt.Printf("Input: %s, ouput (compressed): %s\n", st, compress.Compress(st))
}
