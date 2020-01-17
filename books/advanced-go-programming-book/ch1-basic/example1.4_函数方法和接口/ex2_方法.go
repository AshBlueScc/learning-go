package main

import "sync"

//方法是由函数演变而来，只是将函数的第一个对象参数移动到了函数名前面了而已。
//通过嵌入匿名的成员，我们不仅可以继承匿名成员的内部成员，而且可以继承匿名成员类型所对应的方法。我们一般会将Point看作基类，
//把ColoredPoint看作是它的继承类或子类。不过这种方式继承的方法并不能实现C++中虚函数的多态特性。所有继承来的方法的接收者参数依然是那个匿名成员本身，而不是当前的变量。

type Cache struct {
	m map[string]string
	sync.Mutex
}

func (p *Cache) Lookup(key string) string {
	p.Lock()
	defer p.Unlock()

	return p.m[key]
}

//Cache结构体类型通过嵌入一个匿名的sync.Mutex来继承它的Lock和Unlock方法. 但是在调用p.Lock()和p.Unlock()时,
// p并不是Lock和Unlock方法的真正接收者, 而是会将它们展开为p.Mutex.Lock()和p.Mutex.Unlock()调用. 这种展开是编译期完成的, 并没有运行时代价.

