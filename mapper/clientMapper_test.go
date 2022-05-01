package mapper

import (
	"github.com/PB-Digital/ms-retail-products-info/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildClientQuestionRequest(t *testing.T) {
	clientQuestionDto := &model.ClientQuestionDto{
		Name:    "John",
		Surname: "Lark",
		Phone:   "055",
	}

	clientQuestionRequest := BuildClientQuestionRequest(clientQuestionDto, int64(1))

	assert.Equal(t, clientQuestionRequest.Name, "John")
	assert.Equal(t, clientQuestionRequest.Surname, "Lark")
	assert.Equal(t, clientQuestionRequest.Phone, "055")
	assert.Equal(t, clientQuestionRequest.ProductId, int64(1))
}
