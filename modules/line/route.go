package line

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/bootstrap"
	line_models "github.com/manochatt/line-noti/domain/line/models"
	"github.com/manochatt/line-noti/modules/line/controller"
	"github.com/manochatt/line-noti/modules/line/repository"
	"github.com/manochatt/line-noti/modules/line/usecase"
	"github.com/manochatt/line-noti/mongo"
)

func NewLineRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	lr := repository.NewLineRepository(db, line_models.CollectionLineTemplate)
	lc := &controller.LineController{
		LineUsecase: usecase.NewLineUsecase(lr, timeout),
	}

	// line_template
	group.POST("/line-template", lc.CreateLineTemplate)
	group.GET("/line-template/:line-template-id", lc.FindLineTemplateById)
	group.PATCH("/line-template/:line-template-id", lc.UpdateLineTemplate)

	// lin_notify
	group.POST("/line-notify", lc.SendDirectNotify)
}
