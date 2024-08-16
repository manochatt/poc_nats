package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/bootstrap"
	"github.com/manochatt/line-noti/modules"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	// go kafka_utils.TopicSubscribe(env, db)

	timeOut := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	modules.SetupRoute(env, timeOut, db, gin)

	gin.Run(env.ServerAddress)
}
