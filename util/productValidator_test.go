package util

import (
	"context"
	"github.com/PB-Digital/ms-retail-products-info/model"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	validator = ProductValidator{}
)

func mockContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, model.ContextLogger, log.WithContext(ctx))
	return ctx
}

func TestValidateForClientQuestionRequest_Valid(t *testing.T) {
	dto := model.ClientQuestionDto{
		Name:        "John",
		Surname:     "Lark",
		Phone:       "+9944620606",
		ProductName: "ExtractiveDomesticMortgage",
	}

	//when:
	err := validator.ValidateForClientQuestionRequest(mockContext(), &dto)

	//then:
	assert.Nil(t, err)
}

func TestValidateForClientRequest_Valid(t *testing.T) {
	dto := model.ClientDto{
		Phone:       "+9944620606",
		ProductName: "ExtractiveDomesticMortgage",
		PinCode:     "1234t56",
	}

	//when:
	err := validator.ValidateForClientRequest(mockContext(), &dto)

	//then:
	assert.Nil(t, err)
}
