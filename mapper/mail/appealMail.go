package mail

import (
	"fmt"
	"github.com/PB-Digital/ms-retail-products-info/model"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func PrepareAppealMail(dto *model.ClientDto, email string) *Mail {
	receivers := []ReceiverDto{
		{Alias: false, Address: email},
	}

	subject := fmt.Sprintf("Musteri muracieti")

	file, err := ioutil.ReadFile(filepath.Clean("templates/client_request.html"))
	if err != nil {
		log.Errorf("ActionLog.PrepareAppealMail.error : Error can't read appeal mail template ")
	}

	template := string(file)
	template = strings.Replace(template, "{{Name}}", *dto.Name, 1)
	template = strings.Replace(template, "{{Surname}}", *dto.Surname, 1)
	template = strings.Replace(template, "{{Phone}}", dto.Phone, 1)
	template = strings.Replace(template, "{{ProductName}}", string(dto.ProductName), 1)
	template = strings.Replace(template, "{{PinCode}}", dto.PinCode, 1)
	template = strings.Replace(template, "{{DateOfBirth}}", *dto.DateOfBirth, 1)

	result := &Mail{
		From:    "sdfsfd",
		To:      receivers,
		Subject: subject,
		Message: template,
	}

	return result
}
