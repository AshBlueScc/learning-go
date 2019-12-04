package main

import "fmt"

//探索切片长度
func main() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	//numbers = append(numbers, 2)
	//printSlice(numbers)


	numbers = append(numbers, 2, 3, 4, 5, 6  )
	printSlice(numbers)
	//numbers = append(numbers, 3)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 4)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 5)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 6)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 7)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 8)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 9, 10)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 11, 12)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 13, 14)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 15, 16)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 17, 18)
	//printSlice(numbers)
	//
	//numbers = append(numbers, 19, 20)
	//printSlice(numbers)





	/* 创建切片 numbers1 是之前切片的两倍容量*/
	//numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	//copy(numbers1,numbers)
	//printSlice(numbers1)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}