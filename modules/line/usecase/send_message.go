package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	line_models "github.com/manochatt/line-noti/domain/line/models"
	line_requests "github.com/manochatt/line-noti/domain/line/requests"
)

func (lu *lineUsecase) SendMessage(c context.Context, lineMessageRequest line_requests.LineMessageRequest) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	lineTemplates, err := lu.lineRepository.FindLineTemplateByProjectId(ctx, lineMessageRequest.ProjectID)
	if err != nil {
		return err
	}
	if len(lineTemplates) == 0 {
		return errors.New("0 line template found")
	}
	fmt.Println(lineTemplates, lineMessageRequest.ProjectID)
	payloadData := line_models.LineMessage{
		To:       lineMessageRequest.ToID,
		Messages: lineTemplates[0].Messages,
	}
	payload, err := json.Marshal(payloadData)
	if err != nil {
		return err
	}

	replacer := strings.NewReplacer(
		"${Title}", lineMessageRequest.MessageValue.Title,
		"${Place}", lineMessageRequest.MessageValue.Place,
		"${Time}", lineMessageRequest.MessageValue.StartDateTime,
	)

	message := replacer.Replace(string(payload))

	return lu.lineRepository.SendNotify(ctx, bytes.NewBuffer([]byte(message)))
}
