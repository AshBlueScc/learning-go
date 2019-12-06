package main

import "fmt"

func main() {
	//_, numb, str := numbers()
	//fmt.Println(numb, str)
	show_func := show(1, 2)
	fmt.Println(show_func(3, 4))
}

func numbers()(int, int, string){
	a, b, c := 1, 2, "str"
	return a, b, c
}

//闭包使用方法
func show(x1, x2 int) func(x3 int, x4 int)(int ,int, int, int){
	return func(x3 int, x4 int) (int ,int, int, int) {
		return x1, x2, x3, x4
	}
}