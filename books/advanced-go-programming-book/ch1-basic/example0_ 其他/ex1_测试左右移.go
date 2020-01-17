package main

import "fmt"

func LRmove()(){
	ints := [3 << 2]int{1, 2, 3}
	fmt.Printf("%#v", ints)
	//这个表示3转化成2进制以后左移两位，后面补0.右移的话删除左边的位上的值

	//左移
	var a int = 1
	a = a << 10
	fmt.Print("左移后的结果为:")
	fmt.Println(a)

	//右移
	var b int = 1024
	b = b >> 10
	fmt.Print("右移后的结果为:")
	fmt.Println(b)
}