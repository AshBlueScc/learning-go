package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"os"
	"time"
	utils2 "xorm/千锋/projects/example0_login_project/utils"
)

func main() {
	app := iris.New()
	app.Logger().
		SetLevel(logLevel).
		SetOutput(os.Stdout).
		SetTimeFormat(time.RFC3339)

	app.Use(recover.New())
	app.Use(utils2.HTTPLogger)

	initRouter(app)

	_ = app.Run(iris.Addr(fmt.Sprintf(":%d", serverPort)), iris.WithConfiguration(iris.Configuration{
		DisableBodyConsumptionOnUnmarshal: true,
		FireMethodNotAllowed:              true,
		PostMaxMemory:                     1024 * 1024 * 1,
	}))
}
