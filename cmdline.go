// cmd line args input
// go run cmdline.go foo bar
package main

import (
	"fmt"
	"os"
)

func main() {
	// sep is just a space between args
	var s, sep string
	// for _, arg := range os.Args[1:] 
	// another way loop above
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " " 
	}
	fmt.Println(s)
}
