package repo

import (
	"github.com/PB-Digital/ms-retail-products-info/model"
	log "github.com/sirupsen/logrus"
)

type IRetailProductRepo interface {
	SaveClientQuestionRequest(clientQuestionRequest *model.ClientQuestionRequest) error
	SaveClientRequest(clientRequest *model.ClientRequest) (int64, error)
	SaveClientAgreement(clientAgreement *model.ClientAgreement) error
	GetProductByName(name string) (*model.Product, error)
	GetClientQuestionRequestByNameAndSurnameAndPhoneAndCreatedInLast1Hour(dto *model.ClientQuestionDto) (*model.ClientQuestionRequest, error)
	GetClientRequestCreatedInLast1Hour(dto *model.ClientDto) (*model.ClientRequest, error)
}

type RetailProductRepo struct{}

func (r *RetailProductRepo) SaveClientQuestionRequest(clientQuestionRequest *model.ClientQuestionRequest) error {

	_, err := Db.Model(clientQuestionRequest).Insert()
	if err != nil {
		return err
	}

	return err
}

func (r *RetailProductRepo) SaveClientRequest(clientRequest *model.ClientRequest) (int64, error) {

	_, err := Db.Model(clientRequest).Returning("id").Insert()
	if err != nil {
		return 0, err
	}

	return clientRequest.Id, nil
}

func (r *RetailProductRepo) SaveClientAgreement(clientAgreement *model.ClientAgreement) error {

	_, err := Db.Model(clientAgreement).Insert()
	if err != nil {
		return err
	}

	return err
}

func (r *RetailProductRepo) GetProductByName(name string) (*model.Product, error) {
	var product model.Product
	err := Db.Model(&product).
		Where("name = ?", name).
		Where("status = ?", "ACTIVE").
		Select()

	if err != nil {
		log.Fatal(err)
	}

	return &product, nil
}

func (r *RetailProductRepo) GetClientQuestionRequestByNameAndSurnameAndPhoneAndCreatedInLast1Hour(dto *model.ClientQuestionDto) (*model.ClientQuestionRequest, error) {
	var clientQuestionRequest model.ClientQuestionRequest
	err := Db.Model(&clientQuestionRequest).
		Where("name = ?", dto.Name).
		Where("surname = ?", dto.Surname).
		Where("phone = ?", dto.Phone).
		Where("created_at > now() AT TIME ZONE 'Asia/Baku' - interval '1 hours'").
		Select()

	if err != nil {
		log.Info(err)
	}
	return &clientQuestionRequest, nil
}

func (r *RetailProductRepo) GetClientRequestCreatedInLast1Hour(dto *model.ClientDto) (*model.ClientRequest, error) {
	var clientRequest model.ClientRequest
	err := Db.Model(&clientRequest).
		Where("phone = ?", dto.Phone).
		Where("pin_code = ?", dto.PinCode).
		Where("created_at > now() AT TIME ZONE 'Asia/Baku' - interval '1 hours'").
		Select()

	if err != nil {
		log.Info(err)
	}
	return &clientRequest, nil
}
