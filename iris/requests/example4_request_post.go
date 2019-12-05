package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type Person struct {
	Name string
	Age int8
}

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	//type：POST请求
	//携带数据：name,pwd命名得请求数据
	//url: http://localhost:8000/postLogin?name=czwqewq&pwd=123eqweqw
	app.Post("/postLogin", func(context context.Context) {
		//获取请求path
		path := context.Path()
		//日志
		app.Logger().Info(path)
		//获取请求字段

		//这个拿不到params里面的参数，暂时不知道这个方法是干啥用的
		//name := context.PostValue("name")

		name := context.FormValue("name")
		pwd := context.FormValue("pwd")
		app.Logger().Info(name, "|", pwd)
		//返回
		//多种返回格式，在返回的body里面, 可以用postman测试
		context.HTML("HTML:"+name)
		context.WriteString("WRITESTRING:"+name + " successfully login!!")
		context.JSON(iris.Map{"message": "hello word", "requestCode": 200})
		context.XML(Person{Name: "Davie", Age: 18})
	})

	//url：http://localhost:8000/user/info
	//type：POST请求，Handle方法第一个参数为POST，表明是Post请求
	app.Handle("POST", "/user/info", func(context context.Context) {
		context.WriteString(" User Info is Post Request , Deal is in handle func ")
	})

	app.Run(iris.Addr(":8000"))
}