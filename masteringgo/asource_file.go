package masteringgo

import (
	"fmt"

	"github.com/mactsouk/go/simpleGitHub"
)

func addtwo() {
	fmt.Println("This is a sample Go program!")
	fmt.Println(simpleGitHub.AddTwo(3, 4))
}
