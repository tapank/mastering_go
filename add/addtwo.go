package add

import (
	"fmt"

	"github.com/mactsouk/go/simpleGitHub"
)

func AddTwo(a, b int) {
	fmt.Printf("AddTwo(%d, %d): %d\n", a, b, simpleGitHub.AddTwo(a, b))
}
