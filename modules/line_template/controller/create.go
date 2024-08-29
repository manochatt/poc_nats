package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/domain"
	"github.com/manochatt/line-noti/domain/models"
	"github.com/manochatt/line-noti/domain/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lc *LineTemplateController) Create(c *gin.Context) {
	var lineTemplateRequest requests.CreateLineTemplateRequest

	err := c.ShouldBind(&lineTemplateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	lineTemplate := models.LineTemplate{
		ID:        primitive.NewObjectID(),
		ProjectID: lineTemplateRequest.ProjectID,
		Messages:  lineTemplateRequest.Messages,
	}

	err = lc.LineTemplateUsecase.Create(c, &lineTemplate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Line Template created successfully",
	})

}
