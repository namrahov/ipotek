package mapper

import "github.com/PB-Digital/ms-retail-products-info/model"

func BuildClientQuestionRequest(dto *model.ClientQuestionDto, productId int64) model.ClientQuestionRequest {
	var request model.ClientQuestionRequest
	request.ProductId = productId
	request.Name = dto.Name
	request.Surname = dto.Surname
	request.Phone = dto.Phone

	return request
}

func BuildClientRequest(dto *model.ClientDto, productId int64) model.ClientRequest {
	var request model.ClientRequest
	request.ProductId = productId

	if dto.Name != nil {
		request.Name = *dto.Name
	}

	if dto.Surname != nil {
		request.Surname = *dto.Surname
	}

	request.Phone = dto.Phone

	if dto.DateOfBirth != nil {
		request.DateOfBirth = *dto.DateOfBirth
	}

	request.PinCode = dto.PinCode

	return request
}

func BuildClientAgreement(dto *model.ClientAgreementDto, clientId int64) model.ClientAgreement {
	var clientAgreement model.ClientAgreement
	clientAgreement.ClientId = clientId
	clientAgreement.AgreementText = dto.AgreementText
	clientAgreement.IsSigned = dto.IsSigned
	clientAgreement.SignDate = dto.SignDate

	return clientAgreement
}
