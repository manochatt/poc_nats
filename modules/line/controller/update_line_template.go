package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/domain"
	"github.com/manochatt/line-noti/domain/requests"
)

func (lc *LineController) UpdateLineTemplate(c *gin.Context) {
	id := c.Param("line-template-id")
	var updateLineTemplateRequest requests.UpdateLineTemplateRequest

	err := c.ShouldBind(&updateLineTemplateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = lc.LineUsecase.UpdateLineTemplate(c, id, &updateLineTemplateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Line Template updated successfully",
	})
}
