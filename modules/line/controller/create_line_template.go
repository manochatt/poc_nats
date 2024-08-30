package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/domain"
	"github.com/manochatt/line-noti/domain/models"
	"github.com/manochatt/line-noti/domain/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lc *LineController) CreateLineTemplate(c *gin.Context) {
	var createLineTemplateRequest requests.CreateLineTemplateRequest

	err := c.ShouldBind(&createLineTemplateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	lineTemplate := models.LineTemplate{
		ID:        primitive.NewObjectID(),
		ProjectID: createLineTemplateRequest.ProjectID,
		Messages:  createLineTemplateRequest.Messages,
	}

	err = lc.LineUsecase.CreateLineTemplate(c, &lineTemplate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Line Template created successfully",
	})

}
