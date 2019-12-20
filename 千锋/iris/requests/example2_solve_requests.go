package main

import (

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	//url:http://localhost:8000/getRequest
	//type:GET请求
	app.Get("/getRequest", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
	})

	//url:http://localhost:8000/user/info
	app.Handle("POST", "/user/info", func(context context.Context) {
		context.WriteString("User Info is Post Request , Deal is in handle func ")
	})

	//启动端口监听服务
	app.Run(iris.Addr(":8000"))
}