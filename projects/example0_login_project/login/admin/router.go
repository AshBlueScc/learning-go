package admin

import "github.com/kataras/iris"

func InitRouter(p iris.Party){
	admin :=p.Party("/admin")

	admin.Post("/login", PostLogin)
}