package echo

import (
	"fmt"
	"os"
)

func main() {
	var result string
	for i := 1; i < len(os.Args); i++ {
		result += os.Args[i] + ","
	}
	fmt.Printf("args is : %s", result)
}
