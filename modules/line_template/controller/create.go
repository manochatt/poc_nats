package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lc *LineTemplateController) Create(c *gin.Context) {
	var lineTemplate domain.LineTemplate

	err := c.ShouldBind(&lineTemplate)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	lineTemplate.ID = primitive.NewObjectID()

	err = lc.LineTemplateUsecase.Create(c, &lineTemplate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Line Template created successfully",
	})

}
