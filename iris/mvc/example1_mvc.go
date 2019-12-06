package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

//自定义的控制器
type CustomController struct {}

type UserController struct{}

func main() {
	//注册自定义控制器处理请求
	app := iris.New()

	//Handie或者下面的Configure用一个就行
	mvc.New(app).Handle(new(CustomController))

	mvc.Configure(app.Party("/user"), func(mvc *mvc.Application){
		mvc.Handle(new(UserController))
	})

	app.Run(iris.Addr(":8000"))

}

//自动处理基础的http请求
//url: http://localhost:8000
//Type: GET请求
func (cc *CustomController) Post() mvc.Result {
	//todo
	return mvc.Response{}
}

/**
* url：http://localhost:8000/user/info
* type：GET请求
**/
func (cc *UserController) GetInfo() string{
	//todo
	return "hello000"
}

/**
* url：http://localhost:8000/user/info
* type：GET请求
**/
func (cc *CustomController) GetUserInfo() string{
	//todo
	return "hello111"
}

/**
* url：http://localhost:8000/login
* type：POST
**/
func (cc *CustomController) PostLogin() mvc.Result{
	//todo
	return mvc.Response{}
}

/**
 * url：http://localhost:8000/users/info
 * type：GET请求
 **/
func (m *CustomController) BeforeActivation(a mvc.BeforeActivation){
	a.Handle("GET", "/users/info", "QueryInfo")
}

func (m *CustomController) QueryInfo() string {
	//todo
	//return mvc.Response{}
	return "says Hey"
}

//handle和configure若有相同的url则以先处理的返回结果为准
//其实两种方式只用一种就可以了，这里只是为了测试