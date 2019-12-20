package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

/*
Session和Cookie的区别：
同：
1.session和cookie两者都是用来存储客户的状态信息的手段。在登陆，注册等动作后，可以存储相关账户的状态信息，方便程序后续跟踪及使用。
异：
1.存储位置：Cookie是存储在客户端浏览器上，方便客户端请求时使用;session存储的相关信息存储在服务器端，用于存储客户端连接的状态信息
2.数据类型：Cookie仅仅支持存储字符串string一种数据类型，Session支持int, string, bool等多种数据类型，Session支持的数据类型更全更多
 */

func main() {
	app := iris.New()
	sessionID := "mySession"
	//session的创建
	sess := sessions.New(sessions.Config{
		Cookie: sessionID,
	})
	app.Logger().SetLevel("debug")




	USERNAME := ""
	ISLOGIN := ""

	app.Post("/login", func(context context.Context) {
		path := context.Path()
		app.Logger().SetLevel("debug")
		app.Logger().Info("请求path: ", path)
		//PostValue用于表单提交，html里面的input
		//userName := context.PostValue("name")
		//passwd := context.PostValue("pwd")

		userName := context.FormValue("name")
		passwd := context.FormValue("pwd")

		if userName == "ash" && passwd == "pwd123" {
			session := sess.Start(context)
			//用户名
			session.Set(USERNAME, userName)
			//登陆状态
			session.Set(ISLOGIN, true)

			context.WriteString("账户登录成功！")
		} else {
			session := sess.Start(context)
			session.Set(ISLOGIN, false)
			context.WriteString("账户登陆失败，请重新尝试！")
		}
	})

	app.Get("/logout", func(context context.Context) {
		path := context.Path()
		app.Logger().SetLevel("debug")
		app.Logger().Info("退出登录Path: ", path)
		session := sess.Start(context)
		//删除session
		session.Delete(ISLOGIN)
		session.Delete(USERNAME)
		context.WriteString("退出登录成功！")
	})

	app.Get("/query", func(context context.Context) {
		path := context.Path()
		app.Logger().SetLevel("debug")
		app.Logger().Info("查询信息 path: ", path)
		session := sess.Start(context)

		isLogin, err := session.GetBoolean(ISLOGIN)
		if err != nil {
			context.WriteString("用户未登录，请先登录！")
			return
		}

		if isLogin {
			app.Logger().Info("账户已登录")
			context.WriteString("账户已登录")
		} else {
			app.Logger().Info("账户未登录")
			context.WriteString("账户未登录")
		}
	})

	app.Run(iris.Addr(":8000"))
}










