package main

//函数对应操作序列，是程序的基本组成元素。Go语言中的函数有具名和匿名之分：具名函数一般对应于包级的函数，是匿名函数的一种特例，
// 当匿名函数引用了外部作用域中的变量时就成了闭包函数，闭包函数是函数式编程语言的核心。

//函数导入顺序：引入包（包里面含有其他包则再次执行引入包）到main --> 创建和初始化常量和变量 -->  调用包里面的init --> 执行main.main函数
//注意：要注意的是，在main.main函数执行之前所有代码都运行在同一个goroutine，也就是程序的主系统线程中。
// 因此，如果某个init函数内部用go关键字启动了新的goroutine的话，新的goroutine只有在进入main.main函数之后才可能被执行到。

//具名函数
func Add(a int, b int) int {
	return a + b
}

//匿名函数
var add = func(a int, b int) int {
	return a + b
}

func Swap(a int, b int) (int, int) {
	return b, a
}

//如果返回值命名了，可以通过名字来修改返回值，也可以通过defer语句在return语句之后修改返回值：
func Inc() (v int) {
	defer func() { v++ }()
	return 42
}

//输出：43
//因为这个匿名函数捕获了外部函数的局部变量v，这种函数我们一般叫闭包,闭包对捕获的外部变量并不是传值方式访问，而是以引用的方式访问。

func Inc1() () {
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
}

// Output:
// 3
// 3
// 3
//因为是闭包，在for迭代语句中，每个defer语句延迟执行的函数引用的都是同一个i迭代变量，在循环结束后这个变量的值为3，因此最终输出的都是3。

func Inc2() () {
	for i := 0; i < 3; i++ {
		i := i // 定义一个循环体内局部变量i
		defer func() { println(i) }()
	}
}

//output:
//2
//1
//0

func Inc3() () {
	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) { println(i) }(i)
	}
	//最后一个小括号表示调用，参考:https://segmentfault.com/q/1010000005050337
}

//output:
//2
//1
//0

// 切片传值还是传引用
func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

type IntSliceHeader struct {
	Data []int
	Len int
	Cap int
}

func twice2(x IntSliceHeader){
	for i := 0; i < x.Len; i++ {
		x.Data[i] *= 2
	}
}



func f(x int) *int {
	return &x
}

func g() int {
	var x = new(int)
	return *x
}