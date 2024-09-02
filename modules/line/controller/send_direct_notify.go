package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/domain"
	line_requests "github.com/manochatt/line-noti/domain/line/requests"
)

func (lc *LineController) SendDirectNotify(c *gin.Context) {
	var lineNotifyDTO line_requests.LineNotifyRequest

	err := c.ShouldBind(&lineNotifyDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	payload, err := json.Marshal(lineNotifyDTO)
	if err != nil {
		log.Fatal("failed to marshal payload:", err)
	}

	replacer := strings.NewReplacer(
		"${Title}", "Banana Cafe",
		"${Place}", "Flex Tower, 7-7-4 Midori-ku, Tokyo",
		"${Time}", "10.00-23.00",
	)
	updatedPayload := replacer.Replace(string(payload))

	err = lc.LineUsecase.SendDirectNotify(c, bytes.NewBuffer([]byte(updatedPayload)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.OkApiResponse[any](nil, "Line Notify send successfully", nil))
}
