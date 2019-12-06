package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}