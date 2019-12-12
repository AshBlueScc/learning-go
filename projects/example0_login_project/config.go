package main

import "xorm/projects/example0_login_project/utils"

var (
	logLevel      = utils.GetEnvStr("LOG_LEVEL", "debug")
	serverPort    = utils.GetEnvInt("SERVER_PORT", 8080)
)

