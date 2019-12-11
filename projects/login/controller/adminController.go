package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"xorm/projects/login/service"
)

type AdminController struct {
	//iris框架自动为每个请求都绑定上下文对象
	Ctx iris.Context

	//admin功能实体
	Service service.AdminService

	//session对象
	Session *sessions.Session

}


//接口:	/admin/login
//请求:	Post
func (ac *AdminController) PostLogin(context iris.Context) mvc.Result {

}