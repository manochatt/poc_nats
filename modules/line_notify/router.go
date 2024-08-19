package line_notify

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/modules/line_notify/controller"
	"github.com/manochatt/line-noti/modules/line_notify/repository"
	"github.com/manochatt/line-noti/modules/line_notify/usecase"
)

func NewLineNotifyRouter(timeout time.Duration, group *gin.RouterGroup) {
	lnr := repository.NewLineNotifyRepository()
	lnc := &controller.LineNotifyController{
		LineNotifyUsecase: usecase.NewLineNotifyUsecase(lnr, timeout),
	}

	group.POST("/line-notify", lnc.SendNotify)
}
