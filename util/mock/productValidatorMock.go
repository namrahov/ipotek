package mock

import (
	"context"
	"github.com/PB-Digital/ms-retail-products-info/model"
	"github.com/stretchr/testify/mock"
)

type ProductValidatorMock struct {
	mock.Mock
}

func (p *ProductValidatorMock) ValidateForClientQuestionRequest(ctx context.Context, dto *model.ClientQuestionDto) error {
	args := p.Called(dto)
	return args.Error(0)
}

func (p *ProductValidatorMock) ValidateForClientRequest(ctx context.Context, dto *model.ClientDto) error {
	args := p.Called(dto)
	return args.Error(0)
}
