package main

import "fmt"

func main(){
	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
	//:表示的意思是，2: 3表示第三个位置放3，1: 2表示第二个位置放2，前面表示index，后面表示值。且按照顺序放，不够补0
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6

	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)

	//Go语言中数组是值语义.
	//1.当一个数组变量被赋值或者被传递的时候，实际上会复制整个数组。如果数组较大的话，数组的赋值也会有较大的开销。
	//为了避免复制数组带来的开销，可以传递一个指向数组的指针，但是数组指针并不是数组。
	var a1 = [...]int{1, 2, 3}
	var b1 = &a1

	fmt.Printf("%d %d\n", a1[0], a1[1])
	fmt.Printf("%d %d\n", b1[0], b1[1])

	for i, v := range b {     // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}
}
/**


 */