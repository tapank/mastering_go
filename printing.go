package main

import (
	"fmt"
	"os"
)

func printing() {
	v1 := 123
	v2 := "123"
	v3 := "have a nice day"
	v4 := "abc"

	fmt.Fprintf(os.Stdout, "Writing to stdout -- v1:%d, v2:%s, v3:%s, v4%s\n", v1, v2, v3, v4)
	fmt.Fprintf(os.Stderr, "Writing to stderr -- v1:%d, v2:%s, v3:%s, v4%s\n", v1, v2, v3, v4)
}
