package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

func main() {
	var a [3]int                 // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}    // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2} // 定义长度为3的int型数组, 元素为 0, 2, 3
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

	for i, v := range b { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

	//2.切片添加元素
	//2.1切片尾部添加元素
	var a2 []int
	a2 = append(a2, 1)               // 追加1个元素
	a2 = append(a2, 1, 2, 3)         // 追加多个元素, 手写解包方式
	a2 = append(a2, []int{1,2,3}...) // 追加一个切片, 切片需要解包
	//2.2切片头部添加元素,只能用切片的方式加
	var a3 = []int{1, 2, 3}
	a3 = append([]int{0}, a3...)
	a3 = append([]int{-3, -2, -1}, a3...)
	//在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。因此，从切片的开头添加元素的性能一般要比从尾部追加元素的性能差很多。
	//2.3切片指定中间为添加元素
	var a4 []int
	var i = 3
	var x = 4
	a4 = append(a4[:i], append([]int{x}, a[i:]...)...)     // 在第i个位置插入x
	a4 = append(a4[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片
	//可以用copy和append组合可以避免创建中间的临时切片，同样是完成添加元素的操作：(copy的两个参数都必须为切片)
	var a5 = []int{1, 2, 3}
	a5 = append(a5, 0)     // 切片扩展1个空间
	copy(a5[i+1:], a5[i:]) // a[i:]向后移动1个位置
	a5[i] = x             // 设置新添加的元素
	//用copy和append组合也可以实现在中间位置插入多个元素(也就是插入一个切片):
	var a6 = []int{1}
	a6 = append(a6, []int{0, 0, 0}...)       // 为x切片扩展足够的空间,...表示解包取一个一个单个的值
	copy(a[i+len(a5):], a5[i:]) // a[i:]向后移动len(x)个位置
	copy(a[i:], a5)            // 复制新添加的切片

	//3.删除切片元素
	//3.1从尾部删除元素
	var a7 = []int{1, 2, 3}
	var N = 2
	a7 = a7[:len(a7)-1]	//尾部删除一个元素
	a7 = a7[:len(a7) - N]	//尾部删除N个元素
	//比较简单，看书就行

	//4.切片内存技巧
	//切片高效操作的要点是要降低内存分配的次数，尽量保证append操作不会超出cap的容量，降低触发内存分配的次数和每次分配内存大小。

	//5.避免切片内存泄漏
	//切片操作并不会复制底层的数据。底层的数组会被保存在内存中，直到它不再被引用。但是有时候可能会因为一个小的内存引用而导致底层整个数组处于被使用的状态，这会延迟自动内存回收器对底层数组的回收。

}


/**
1.在对切片本身赋值或参数传递时，和数组指针的操作方式类似，只是复制切片头信息（reflect.SliceHeader），并不会复制底层的数据。
2.在容量不足的情况下，append的操作会导致重新分配内存，可能导致巨大的内存分配和复制数据代价。即使容量足够，依然需要用append函数的返回值来更新切片本身，因为新切片的长度已经发生了变化。
 */

//6.切片类型强制转换
func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以int方式给float64排序
	sort.Ints(b)
}

func SortFloat64FastV2(a []float64) {
	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	// 以int方式给float64排序
	sort.Ints(c)
}