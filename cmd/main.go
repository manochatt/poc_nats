package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/bootstrap"
	"github.com/manochatt/line-noti/modules"
	"github.com/manochatt/line-noti/utils/nats_utils"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	go nats_utils.Consumer()

	timeOut := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	modules.SetupRoute(env, timeOut, db, gin)

	gin.Run(env.ServerAddress)
}
