package main

import "fmt"

//func fibonacci(n int)(result int){
//	if n < 2 {
//		return n
//	}
//	return fibonacci(n-2)+fibonacci(n-1)
//}

func fibonacci2(n int) (int,int) {
	if n < 2 {
		return 0,n
	}
	a,b := fibonacci2(n-1)
	return b,a+b
}


func fibonacci(n int) int {
	_, b := fibonacci2(n)
	return b
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d  ", fibonacci(i))
	}

}