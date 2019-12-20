package main

import (
	utils2 "xorm/千锋/projects/example0_login_project/utils"
)

var (
	logLevel      = utils2.GetEnvStr("LOG_LEVEL", "debug")
	serverPort    = utils2.GetEnvInt("SERVER_PORT", 8080)
)

