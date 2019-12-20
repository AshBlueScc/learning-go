package main

import (
	"github.com/kataras/iris"
	admin2 "xorm/千锋/projects/example0_login_project/login/admin"
)

func initRouter(app *iris.Application) {
	v1 := app.Party("/api/v1")

	admin2.InitRouter(v1)
}