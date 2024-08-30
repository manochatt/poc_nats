package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/domain"
)

func (lc *LineController) FindLineTemplateById(c *gin.Context) {
	id := c.Param("line-template-id")

	lineTemplates, err := lc.LineUsecase.FindLineTemplateById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.OkApiResponse(&lineTemplates, "", nil))
}
