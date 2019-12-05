package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	//url: http://localhost:8000/userpath
	//type: GET请求，用GET方法处理
	app.Get("/userpath", func(context context.Context){

		//获取Path
		path := context.Path()
		//日志输出
		app.Logger().Info(path)
		//写入返回数据: string类型
		context.WriteString("请求路径：" + path)
	})

	//url: http://localhost:8000/hello
	//type: GET请求，Handle方法第一个参数为GET,表明是GET请求方式
	app.Handle("GET", "/hello", func(context context.Context) {
		context.HTML("<h1> Hello world. </h1>")
	})

	app.Run(iris.Addr(":8000"))
}