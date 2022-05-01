package mock

import (
	"github.com/PB-Digital/ms-retail-products-info/model"
	"github.com/stretchr/testify/mock"
)

type RetailProductRepoMock struct {
	mock.Mock
}

func (r *RetailProductRepoMock) SaveClientQuestionRequest(clientQuestionRequest *model.ClientQuestionRequest) error {
	args := r.Called(clientQuestionRequest)
	return checkArgumentsClientQuestionRequest(args)
}

func (r *RetailProductRepoMock) SaveClientRequest(clientRequest *model.ClientRequest) (int64, error) {
	args := r.Called(clientRequest)
	return checkArgumentsClientRequest(args)
}

func (r *RetailProductRepoMock) SaveClientAgreement(clientAgreement *model.ClientAgreement) error {
	args := r.Called(clientAgreement)
	return checkArgumentsClientQuestionRequest(args)
}

func checkArgumentsClientQuestionRequest(args mock.Arguments) error {
	firstArg := args.Get(0)
	if firstArg != nil {
		return args.Error(1)
	}
	return args.Error(1)
}

func checkArgumentsClientRequest(args mock.Arguments) (int64, error) {
	firstArg := args.Get(0)
	if firstArg != nil {
		return 1, args.Error(1)
	}
	return 1, args.Error(1)
}

func (r *RetailProductRepoMock) GetProductByName(name string) (*model.Product, error) {
	args := r.Called(name)
	return checkArgumentsProduct(args)
}

func checkArgumentsProduct(args mock.Arguments) (*model.Product, error) {
	firstArg := args.Get(0)
	if firstArg != nil {
		return firstArg.(*model.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (r *RetailProductRepoMock) GetClientQuestionRequestByNameAndSurnameAndPhoneAndCreatedInLast1Hour(dto *model.ClientQuestionDto) (*model.ClientQuestionRequest, error) {
	args := r.Called(dto)
	return checkArgumentsForGetClientQuestionRequest(args)
}

func (r *RetailProductRepoMock) GetClientRequestCreatedInLast1Hour(dto *model.ClientDto) (*model.ClientRequest, error) {
	args := r.Called(dto)
	return checkArgumentsForGetClientRequest(args)
}

func checkArgumentsForGetClientQuestionRequest(args mock.Arguments) (*model.ClientQuestionRequest, error) {
	firstArg := args.Get(0)
	if firstArg != nil {
		return firstArg.(*model.ClientQuestionRequest), args.Error(1)
	}
	return nil, args.Error(1)
}

func checkArgumentsForGetClientRequest(args mock.Arguments) (*model.ClientRequest, error) {
	firstArg := args.Get(0)
	if firstArg != nil {
		return firstArg.(*model.ClientRequest), args.Error(1)
	}
	return nil, args.Error(1)
}
