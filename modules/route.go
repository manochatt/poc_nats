package modules

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/bootstrap"
	"github.com/manochatt/line-noti/modules/line_template"
	"github.com/manochatt/line-noti/mongo"
)

func SetupRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicGroup := gin.Group("")

	line_template.NewLineTemplateRouter(env, timeout, db, publicGroup)
}
