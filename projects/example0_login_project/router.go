package main

import (
	"github.com/kataras/iris"
	"xorm/projects/example0_login_project/login/admin"
)

func initRouter(app *iris.Application) {
	v1 := app.Party("/api/v1")

	admin.InitRouter(v1)
}