package mail

import (
	"fmt"
	"github.com/PB-Digital/ms-retail-products-info/model"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func PrepareAppealMailForPashaGroup(dto *model.ClientDto, email string) *Mail {
	receivers := []ReceiverDto{
		{Alias: true, Address: email},
	}

	subject := fmt.Sprintf("Musteri muracieti")

	file, err := ioutil.ReadFile(filepath.Clean("templates/client_request_for_pasha_group.html"))
	if err != nil {
		log.Errorf("ActionLog.PrepareAppealMailForPashaGroup.error : Error can't read appeal mail for pasha group template ")
	}

	template := string(file)
	template = strings.Replace(template, "{{Phone}}", dto.Phone, 1)
	template = strings.Replace(template, "{{ProductName}}", string(dto.ProductName), 1)
	template = strings.Replace(template, "{{PinCode}}", dto.PinCode, 1)

	result := &Mail{
		From:    "FROM_IB",
		To:      receivers,
		Subject: subject,
		Message: template,
	}

	return result
}
