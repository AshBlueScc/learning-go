package admin

import (
	"github.com/kataras/iris/context"
	"log"
	"os"
	"xorm/projects/example0_login_project/database"
)


func PostLogin(context context.Context){
	service := database.GetService()

	name := context.FormValue("name")
	pwd := context.FormValue("pwd")


	var user *database.User
	var isAdmin bool
	user, isAdmin = service.GetByAdminNameAndPassword(name, pwd)

	if isAdmin {
		context.HTML("Hi，" + user.UserName + "登陆成功，欢迎您！")

		logger := log.New(os.Stdout, "[PostLogin]", log.Lshortfile | log.Ldate | log.Ltime)
		logger.Printf("管理员登录，信息：%#v", user)
	}
}
