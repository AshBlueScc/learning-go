package utils

import (
	"github.com/kataras/iris/context"
	"net/http"
	"os"
	"strconv"
	"strings"
	types2 "xorm/千锋/projects/example0_login_project/types"
)


func GetEnvStr(key string, defaultValue string) string {
	e := os.Getenv(key)	//获取系统key的环境变量，如果没有环境变量就返回空
	if len(e) <= 0 {
		return defaultValue
	}
	return os.ExpandEnv(e)	//根据环境变量的值替换当前值
}

func GetEnvStrArray(key string) []string {
	e := os.Getenv(key)
	if len(e) <= 0 {
		return nil
	}
	e = os.ExpandEnv(e)
	if len(e) <= 0 {
		return nil
	}
	return strings.Split(e, ",")
}

func GetEnvBool(key string, defaultValue bool) bool {
	e := GetEnvStr(key, "")
	if len(e) <= 0 {
		return defaultValue
	}
	e = strings.ToLower(e)
	switch e {
	case "false", "0", "off", "no":
		return false
	case "true", "1", "on", "yes":
		return true
	}
	return false
}

func GetEnvInt(key string, defaultValue int) int {
	return StrToInt(GetEnvStr(key, ""), defaultValue)
}

func StrToInt(val string, defaultValue int) int {
	v, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return v
}

func StrToInt64(val string, defaultValue int64) int64 {
	v, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

func SendMessage(ctx context.Context, statusCode int, message string) {
	ctx.StatusCode(statusCode)
	ctx.JSON(types2.ErrorMsg{
		Message: message,
	})
}

func SendError(ctx context.Context, err error) {
	SendMessage(ctx, http.StatusInternalServerError, err.Error())
}