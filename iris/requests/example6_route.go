package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	var err error
	app := iris.New()

	app.Logger().SetLevel("debug")

	//url：http://localhost:8000/api/users/false
	app.Get("/api/users/{isLogin:bool}", func(context context.Context){
		path := context.Path()
		app.Logger().Info("请求url:" + path)

		isLogin, err := context.Params().GetBool("isLogin")
		if err != nil {
			context.StatusCode(iris.StatusNonAuthoritativeInfo)
			return
		}
		if isLogin {
			context.WriteString("已登录")
		} else {
			context.WriteString("未登录")
		}

	})

	err = app.Run(iris.Addr(":8000"), iris.WithCharset("utf-8"))
	app.Logger().Error(err)
}
