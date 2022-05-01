package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/PB-Digital/ms-retail-products-info/model"
	queueMock "github.com/PB-Digital/ms-retail-products-info/queue/mock"
	repoMock "github.com/PB-Digital/ms-retail-products-info/repo/mock"
	validatorMock "github.com/PB-Digital/ms-retail-products-info/util/mock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func mockContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, model.ContextLogger, log.WithContext(ctx))
	return ctx
}

var mockRepo = repoMock.RetailProductRepoMock{}
var mockValidator = validatorMock.ProductValidatorMock{}
var messageSenderMock = queueMock.MessageSenderMock{}

var s = RetailProductService{
	RetailProductRepo: &mockRepo,
	Validator:         &mockValidator,
	MessageSender:     &messageSenderMock,
}

func TestSaveClientQuestionRequest_Success(t *testing.T) {
	dto := model.ClientQuestionDto{
		Name:        "John",
		Surname:     "Lark",
		Phone:       "055",
		ProductName: "ExtractiveDomesticMortgage",
	}

	product := model.Product{
		Id:          int64(1),
		Name:        "John",
		NameAz:      "ExtractiveDomesticMortgage",
		Status:      "ACTIVE",
		Description: "description",
	}

	clientQuestionRequestWithZeroId := model.ClientQuestionRequest{
		Id:      int64(0),
		Name:    dto.Name,
		Surname: dto.Surname,
		Phone:   dto.Phone,
	}

	mockValidator.On("ValidateForClientQuestionRequest", &dto).Once().Return(nil, nil)
	mockRepo.On("GetClientQuestionRequestByNameAndSurnameAndPhoneAndCreatedInLast1Hour", &dto).
		Once().Return(&clientQuestionRequestWithZeroId, nil)
	mockRepo.On("GetProductByName", string(dto.ProductName)).Once().Return(&product, nil)
	mockRepo.On("SaveClientQuestionRequest", mock.Anything).Once().Return(nil, nil)
	messageSenderMock.On("SendMessage", mock.Anything, mock.Anything).Once().Return(nil, nil)

	//when:
	err := s.SaveClientQuestionRequest(&dto, mockContext())

	//then:
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSaveClientQuestionRequest_CanNotGetProduct(t *testing.T) {
	dto := model.ClientQuestionDto{
		Name:        "John",
		Surname:     "Lark",
		Phone:       "055",
		ProductName: "ExtractiveDomesticMortgage",
	}

	clientQuestionRequestWithZeroId := model.ClientQuestionRequest{
		Id:      int64(0),
		Name:    dto.Name,
		Surname: dto.Surname,
		Phone:   dto.Phone,
	}

	mockValidator.On("ValidateForClientQuestionRequest", &dto).Once().Return(nil, nil)
	mockRepo.On("GetClientQuestionRequestByNameAndSurnameAndPhoneAndCreatedInLast1Hour", &dto).
		Once().Return(&clientQuestionRequestWithZeroId, nil)
	mockRepo.On("GetProductByName", string(dto.ProductName)).Once().
		Return(nil, errors.New(fmt.Sprintf("%s.could-not-get-product", model.Exception)))

	//when:
	err := s.SaveClientQuestionRequest(&dto, mockContext())

	//then:
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fmt.Sprintf("%s.could-not-get-product", model.Exception))
	mockRepo.AssertExpectations(t)
}

func TestSaveClientQuestionRequest_CanNotSaveClientQuestionRequest(t *testing.T) {
	dto := model.ClientQuestionDto{
		Name:        "John",
		Surname:     "Lark",
		Phone:       "055",
		ProductName: "ExtractiveDomesticMortgage",
	}

	product := model.Product{
		Id:          int64(1),
		Name:        "John",
		NameAz:      "ExtractiveDomesticMortgage",
		Status:      "ACTIVE",
		Description: "description",
	}

	clientQuestionRequestWithZeroId := model.ClientQuestionRequest{
		Id:      int64(0),
		Name:    dto.Name,
		Surname: dto.Surname,
		Phone:   dto.Phone,
	}

	mockValidator.On("ValidateForClientQuestionRequest", &dto).Once().Return(nil, nil)
	mockRepo.On("GetClientQuestionRequestByNameAndSurnameAndPhoneAndCreatedInLast1Hour", &dto).
		Once().Return(&clientQuestionRequestWithZeroId, nil)
	mockRepo.On("GetProductByName", string(dto.ProductName)).Once().Return(&product, nil)
	mockRepo.On("SaveClientQuestionRequest", mock.Anything).Once().
		Return(nil, errors.New(fmt.Sprintf("%s.could-not-get-product", model.Exception)))

	//when:
	err := s.SaveClientQuestionRequest(&dto, mockContext())

	//then:
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fmt.Sprintf("%s.could-not-save-client-question-request", model.Exception))
	mockRepo.AssertExpectations(t)
}

func TestSaveClientQuestionRequest_NameOrSurnameIsEmpty(t *testing.T) {
	dto := model.ClientQuestionDto{
		Name:        "",
		Surname:     "Lark",
		Phone:       "055",
		ProductName: "ExtractiveDomesticMortgage",
	}

	mockValidator.On("ValidateForClientQuestionRequest", &dto).Once().
		Return(errors.New(fmt.Sprintf("%s.empty-name-or-surname", model.Exception)))

	//when:
	err := s.SaveClientQuestionRequest(&dto, mockContext())

	//then:
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fmt.Sprintf("%s.empty-name-or-surname", model.Exception))
	mockRepo.AssertExpectations(t)
}

func TestSaveClientQuestionRequest_LongNameOrSurname(t *testing.T) {
	dto := model.ClientQuestionDto{
		Name:        "Abccvcvcvcvcvcvcvcvcvcvc",
		Surname:     "Lark",
		Phone:       "055",
		ProductName: "ExtractiveDomesticMortgage",
	}

	mockValidator.On("ValidateForClientQuestionRequest", &dto).Once().
		Return(errors.New(fmt.Sprintf("%s.long-name-or-surname", model.Exception)))

	//when:
	err := s.SaveClientQuestionRequest(&dto, mockContext())

	//then:
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fmt.Sprintf("%s.long-name-or-surname", model.Exception))
	mockRepo.AssertExpectations(t)
}

func TestSaveClientQuestionRequest_PhoneNumberIsEmpty(t *testing.T) {
	dto := model.ClientQuestionDto{
		Name:        "John1",
		Surname:     "Lark",
		Phone:       "",
		ProductName: "ExtractiveDomesticMortgage",
	}

	mockValidator.On("ValidateForClientQuestionRequest", &dto).Once().
		Return(errors.New(fmt.Sprintf("%s.empty-phone-number", model.Exception)))

	//when:
	err := s.SaveClientQuestionRequest(&dto, mockContext())

	//then:
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fmt.Sprintf("%s.empty-phone-number", model.Exception))
	mockRepo.AssertExpectations(t)
}

func TestSaveClientQuestionRequest_PhoneNumberIsTooLong(t *testing.T) {
	dto := model.ClientQuestionDto{
		Name:        "John1",
		Surname:     "Lark",
		Phone:       "+9942424234536214234",
		ProductName: "ExtractiveDomesticMortgage",
	}

	mockValidator.On("ValidateForClientQuestionRequest", &dto).Once().
		Return(errors.New(fmt.Sprintf("%s.phone-length-is-too-long", model.Exception)))

	//when:
	err := s.SaveClientQuestionRequest(&dto, mockContext())

	//then:
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fmt.Sprintf("%s.phone-length-is-too-long", model.Exception))
	mockRepo.AssertExpectations(t)
}

func TestSaveClientQuestionRequest_PhoneNumberWithWrongPrefix(t *testing.T) {
	dto := model.ClientQuestionDto{
		Name:        "John1",
		Surname:     "Lark",
		Phone:       "050-462-06-06",
		ProductName: "ExtractiveDomesticMortgage",
	}

	mockValidator.On("ValidateForClientQuestionRequest", &dto).Once().
		Return(errors.New(fmt.Sprintf("%s.phone-length-is-too-long", model.Exception)))

	//when:
	err := s.SaveClientQuestionRequest(&dto, mockContext())

	//then:
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fmt.Sprintf("%s.phone-length-is-too-long", model.Exception))
	mockRepo.AssertExpectations(t)
}

func TestSaveClientRequest_Success(t *testing.T) {
	dateOfBirth := new(string)
	*dateOfBirth = "12.11.2021"

	dto := model.ClientDto{
		Phone:       "055",
		PinCode:     "wew33",
		DateOfBirth: dateOfBirth,
		ClientAgreementDto: &model.ClientAgreementDto{
			AgreementText: "test",
			IsSigned:      true,
			SignDate:      "11.12.2021",
		},
		ProductName: "ExtractiveDomesticMortgage",
	}

	product := model.Product{
		Id:          int64(1),
		Name:        "John",
		NameAz:      "ExtractiveDomesticMortgage",
		Status:      "ACTIVE",
		Description: "description",
	}

	clientRequestWithZeroId := model.ClientRequest{
		Id:          int64(0),
		Phone:       dto.Phone,
		PinCode:     dto.PinCode,
		DateOfBirth: *dto.DateOfBirth,
	}

	mockValidator.On("ValidateForClientRequest", &dto).Once().Return(nil, nil)
	mockRepo.On("GetClientRequestCreatedInLast1Hour", &dto).
		Once().Return(&clientRequestWithZeroId, nil)
	mockRepo.On("GetProductByName", string(dto.ProductName)).Once().Return(&product, nil)
	mockRepo.On("SaveClientRequest", mock.Anything).Once().Return(nil, nil)
	mockRepo.On("SaveClientAgreement", mock.Anything).Once().Return(nil, nil)

	//when:
	err := s.SaveClientRequest(&dto, mockContext())

	//then:
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSaveClientRequest_CanNotGetProduct(t *testing.T) {
	dateOfBirth := new(string)
	*dateOfBirth = "12.11.2021"

	dto := model.ClientDto{
		Phone:       "055",
		PinCode:     "wew33",
		DateOfBirth: dateOfBirth,
		ClientAgreementDto: &model.ClientAgreementDto{
			AgreementText: "test",
			IsSigned:      true,
			SignDate:      "11.12.2021",
		},
		ProductName: "ExtractiveDomesticMortgage",
	}

	clientRequestWithZeroId := model.ClientRequest{
		Id:    int64(0),
		Phone: dto.Phone,
	}

	mockValidator.On("ValidateForClientRequest", &dto).Once().Return(nil, nil)
	mockRepo.On("GetClientRequestCreatedInLast1Hour", &dto).
		Once().Return(&clientRequestWithZeroId, nil)
	mockRepo.On("GetProductByName", string(dto.ProductName)).Once().
		Return(nil, errors.New(fmt.Sprintf("%s.could-not-get-product", model.Exception)))

	//when:
	err := s.SaveClientRequest(&dto, mockContext())

	//then:
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fmt.Sprintf("%s.could-not-get-product", model.Exception))
	mockRepo.AssertExpectations(t)
}

func TestSaveClientRequest_CanNotSaveClientRequest(t *testing.T) {
	dateOfBirth := new(string)
	*dateOfBirth = "12.11.2021"

	dto := model.ClientDto{
		Phone:       "055",
		PinCode:     "wew33",
		DateOfBirth: dateOfBirth,
		ClientAgreementDto: &model.ClientAgreementDto{
			AgreementText: "test",
			IsSigned:      true,
			SignDate:      "11.12.2021",
		},
		ProductName: "ExtractiveDomesticMortgage",
	}

	product := model.Product{
		Id:          int64(1),
		Name:        "John",
		NameAz:      "ExtractiveDomesticMortgage",
		Status:      "ACTIVE",
		Description: "description",
	}

	clientRequestWithZeroId := model.ClientRequest{
		Id:          int64(0),
		Phone:       dto.Phone,
		PinCode:     dto.PinCode,
		DateOfBirth: *dto.DateOfBirth,
	}

	mockValidator.On("ValidateForClientRequest", &dto).Once().Return(nil, nil)
	mockRepo.On("GetClientRequestCreatedInLast1Hour", &dto).
		Once().Return(&clientRequestWithZeroId, nil)
	mockRepo.On("GetProductByName", string(dto.ProductName)).Once().Return(&product, nil)
	mockRepo.On("SaveClientRequest", mock.Anything).Once().
		Return(nil, errors.New(fmt.Sprintf("%s.could-not-save-client-request", model.Exception)))

	//when:
	err := s.SaveClientRequest(&dto, mockContext())

	//then:
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fmt.Sprintf("%s.could-not-save-client-request", model.Exception))
	mockRepo.AssertExpectations(t)
}

func TestSaveClientRequest_ValidatePinCode(t *testing.T) {
	dateOfBirth := new(string)
	*dateOfBirth = "12.11.2021"

	dto := model.ClientDto{
		Phone:       "055",
		PinCode:     "lo35dgfw45dgd45dgddf",
		DateOfBirth: dateOfBirth,
		ClientAgreementDto: &model.ClientAgreementDto{
			AgreementText: "test",
			IsSigned:      true,
			SignDate:      "11.12.2021",
		},
		ProductName: "ExtractiveDomesticMortgage",
	}

	mockValidator.On("ValidateForClientRequest", &dto).Once().
		Return(errors.New(fmt.Sprintf("%s.long-pincode", model.Exception)))

	//when:
	err := s.SaveClientRequest(&dto, mockContext())

	//then:
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fmt.Sprintf("%s.long-pincode", model.Exception))
	mockRepo.AssertExpectations(t)
}

func TestSaveClientRequest_ValidateDateOfBirth(t *testing.T) {
	dateOfBirth := new(string)
	*dateOfBirth = "12.11.2021.1111.11"

	dto := model.ClientDto{
		Phone:       "055",
		PinCode:     "lo35dgfw45dgd45dgddf",
		DateOfBirth: dateOfBirth,
		ClientAgreementDto: &model.ClientAgreementDto{
			AgreementText: "test",
			IsSigned:      true,
			SignDate:      "11.12.2021",
		},
		ProductName: "ExtractiveDomesticMortgage",
	}

	mockValidator.On("ValidateForClientRequest", &dto).Once().
		Return(errors.New(fmt.Sprintf("%s.max-word-in-date-of-birth", model.Exception)))

	//when:
	err := s.SaveClientRequest(&dto, mockContext())

	//then:
	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fmt.Sprintf("%s.max-word-in-date-of-birth", model.Exception))
	mockRepo.AssertExpectations(t)
}
