package line_template

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/bootstrap"
	"github.com/manochatt/line-noti/domain/models"
	"github.com/manochatt/line-noti/modules/line_template/controller"
	"github.com/manochatt/line-noti/modules/line_template/repository"
	"github.com/manochatt/line-noti/modules/line_template/usecase"
	"github.com/manochatt/line-noti/mongo"
)

func NewLineTemplateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	lr := repository.NewLineTemplateRepository(db, models.CollectionLineTemplate)
	lc := &controller.LineTemplateController{
		LineTemplateUsecase: usecase.NewLineTemplateUsecase(lr, timeout),
	}

	group.POST("/line-template", lc.Create)
	group.GET("/line-template/:line-template-id", lc.Fetch)
}
