package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

//路由组
func main () {
	app := iris.New()
	baseUrl := "/admin"

	route1 := app.Party(baseUrl+"/admin1", func(context context.Context) {
		//to do 处理这个路由组的方法
	})

	fmt.Print(route1)
}
