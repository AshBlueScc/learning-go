package main

import "fmt"

func main() {
	var s *int
	var a int = 12

	*s = a
	//s = &a
	fmt.Println(s)
}
