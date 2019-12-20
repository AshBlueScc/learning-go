package main

import "fmt"

func Factorial(n uint64)(result uint64){
	if n > 0 {
		return n*Factorial(n-1)
	}
	return 1
}

func main() {
	fmt.Printf("15的阶乘是 %d", Factorial(15))
}