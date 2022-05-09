package masteringgo

import (
	"fmt"
	"os"
)

func renamefile() {
	var fileName, newFileName string
	if len(os.Args) != 3 {
		fmt.Println("usage: go run osutils <fileName> <newFileName>")
		return
	}
	fileName, newFileName = os.Args[1], os.Args[2]
	fmt.Println("files:", fileName, newFileName)
	if _, err := os.Stat(fileName); err != nil {
		fmt.Println(err)
		return
	}
	if err := os.Rename(fileName, newFileName); err != nil {
		fmt.Println(err)
	}
}
