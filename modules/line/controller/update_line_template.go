package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/domain"
	line_requests "github.com/manochatt/line-noti/domain/line/requests"
)

func (lc *LineController) UpdateLineTemplate(c *gin.Context) {
	id := c.Param("line-template-id")
	var updateLineTemplateRequest line_requests.UpdateLineTemplateRequest

	err := c.ShouldBind(&updateLineTemplateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.BadRequestWith(err.Error()))
		return
	}

	err = lc.LineUsecase.UpdateLineTemplate(c, id, &updateLineTemplateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.OkApiResponse[any](nil, "Line Template updated successfully", nil))
}
