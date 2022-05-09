package masteringgo

import (
	"bufio"
	"fmt"
	"os"
)

func mimic() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
