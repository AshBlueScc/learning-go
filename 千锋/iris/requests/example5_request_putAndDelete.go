package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	//type:PUT类型请求
	app.Put("/putinfo", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("请求url:" + path)
	})

	//type:DELETE类型请求
	app.Delete("/deleteuser", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("请求url:" + path)
	})
}