package utils

import (
	"github.com/kataras/iris/context"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

// HTTPLogger ...
func HTTPLogger(ctx context.Context) {
	w := ctx.Recorder()
	ip := ctx.RemoteAddr()
	method := ctx.Method()
	path := ctx.Request().URL.RequestURI()
	body, _ := ctx.GetBody()

	ctx.Next()

	status := strconv.Itoa(ctx.GetStatusCode())

	resp := strings.Replace(string(w.Body()), "\n", "\\n", -1)
	logrus.Infof("%s: %s [%s %s] -> [%s] [%s]", ip, method, path, string(body), status, resp)
}

