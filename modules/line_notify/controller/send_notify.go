package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/domain"
)

func (lnc *LineNotifyController) SendNotify(c *gin.Context) {
	var lineNotifyDTO domain.LineNotifyDTO

	err := c.ShouldBind(&lineNotifyDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	payload, err := json.Marshal(lineNotifyDTO)
	if err != nil {
		log.Fatal("failed to marshal payload:", err)
	}

	err = lnc.LineNotifyUsecase.SendNotify(c, bytes.NewBuffer(payload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Line Notify send successfully",
	})
}
