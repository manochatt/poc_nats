package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manochatt/line-noti/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lc *LineTemplateController) Create(c *gin.Context) {
	var lineTemplateDTO domain.CreateLineTemplateDTO

	err := c.ShouldBind(&lineTemplateDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	lineTemplate := domain.LineTemplate{
		ID:        primitive.NewObjectID(),
		ProjectID: lineTemplateDTO.ProjectID,
		Messages:  lineTemplateDTO.Messages,
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
