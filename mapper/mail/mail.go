package mail

import (
	"fmt"
	"github.com/PB-Digital/ms-retail-products-info/model"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Mail struct {
	From        string          `json:"from"`
	To          []ReceiverDto   `json:"to"`
	Subject     string          `json:"subject"`
	Message     string          `json:"message"`
	Attachments []AttachmentDto `json:"attachments"`
}

type AttachmentDto struct {
	FileName         string `json:"fileName"`
	FileId           string `json:"fileId"`
	AttachmentSource string `json:"attachmentSource"`
}

type ReceiverDto struct {
	Alias   bool   `json:"alias"`
	Address string `json:"address"`
}

func PrepareMail(dto *model.ClientQuestionDto, email string) *Mail {
	receivers := []ReceiverDto{
		{Alias: false, Address: email},
	}

	subject := fmt.Sprintf("Mene zeng edin")

	file, err := ioutil.ReadFile(filepath.Clean("templates/client_question_request.html"))
	if err != nil {
		log.Errorf("ActionLog.PrepareMail.error : Error can't read mail template ")
	}

	template := string(file)
	template = strings.Replace(template, "{{Name}}", dto.Name, 1)
	template = strings.Replace(template, "{{Surname}}", dto.Surname, 1)
	template = strings.Replace(template, "{{Phone}}", dto.Phone, 1)
	template = strings.Replace(template, "{{ProductName}}", string(dto.ProductName), 1)

	result := &Mail{
		From:    "FROM_IB",
		To:      receivers,
		Subject: subject,
		Message: template,
	}

	return result
}
