package add

import (
	"fmt"

	"github.com/mactsouk/go/simpleGitHub"
)

func AddTwo(a, b int) {
	fmt.Println("This is a sample Go program!")
	fmt.Println(simpleGitHub.AddTwo(a, b))
}
