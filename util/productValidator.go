package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/PB-Digital/ms-retail-products-info/model"
	log "github.com/sirupsen/logrus"
	"strings"
	"unicode/utf8"
)

type IProductValidator interface {
	ValidateForClientQuestionRequest(ctx context.Context, dto *model.ClientQuestionDto) error
	ValidateForClientRequest(ctx context.Context, dto *model.ClientDto) error
}

type ProductValidator struct {
}

func (v *ProductValidator) ValidateForClientQuestionRequest(ctx context.Context, dto *model.ClientQuestionDto) error {
	_, e := v.ValidateNameAndSurname(ctx, &dto.Name, &dto.Surname)
	if e != nil {
		return e
	}

	_, e = v.ValidatePhoneNumber(ctx, dto.Phone)
	if e != nil {
		return e
	}

	_, e = v.ValidateProductName(ctx, dto.ProductName)
	if e != nil {
		return e
	}

	return nil
}

func (v *ProductValidator) ValidateForClientRequest(ctx context.Context, dto *model.ClientDto) error {

	_, e := v.ValidateNameAndSurname(ctx, dto.Name, dto.Surname)
	if e != nil {
		return e
	}

	_, e = v.ValidatePhoneNumber(ctx, dto.Phone)
	if e != nil {
		return e
	}

	_, e = v.ValidateProductName(ctx, dto.ProductName)
	if e != nil {
		return e
	}

	_, e = v.ValidatePinCode(ctx, dto.PinCode)
	if e != nil {
		return e
	}

	if dto.DateOfBirth != nil {
		_, e = v.ValidateDateOfBirth(ctx, dto.DateOfBirth)
		if e != nil {
			return e
		}
	}

	return nil
}

func (v *ProductValidator) ValidateNameAndSurname(ctx context.Context, name *string, surname *string) (bool, error) {
	logger := ctx.Value(model.ContextLogger).(*log.Entry)

	if name == nil && surname == nil {
		return true, nil
	}

	if utf8.RuneCountInString(*name) == 0 || utf8.RuneCountInString(*surname) == 0 {
		logger.Errorf("ActionLog.ValidateName.error: name or surname is empty")
		return false, errors.New(fmt.Sprintf("%s.empty-name-or-surname", model.Exception))
	}

	if utf8.RuneCountInString(*name) > 20 || utf8.RuneCountInString(*surname) > 20 {
		logger.Errorf("ActionLog.ValidateName.error: name or surname is more than 20 symbols")
		return false, errors.New(fmt.Sprintf("%s.long-name-or-surname", model.Exception))
	}

	return true, nil
}

func (v *ProductValidator) ValidatePhoneNumber(ctx context.Context, phoneNumber string) (bool, error) {
	logger := ctx.Value(model.ContextLogger).(*log.Entry)
	prefix := phoneNumber[0:4]

	if utf8.RuneCountInString(phoneNumber) == 0 {
		logger.Errorf("ActionLog.ValidatePhoneNumber.error: phone is empty")
		return false, errors.New(fmt.Sprintf("%s.empty-phone-number", model.Exception))
	}

	if utf8.RuneCountInString(phoneNumber) > 20 {
		logger.Errorf("ActionLog.ValidatePhoneNumber.error: phone number length is too long")
		return false, errors.New(fmt.Sprintf("%s.phone-length-is-too-long", model.Exception))
	}

	if prefix != "+994" {
		logger.Errorf("ActionLog.ValidatePhoneNumber.error: phone number prefix is not start with +994")
		return false, errors.New(fmt.Sprintf("%s.wrong-phone-prefix", model.Exception))
	}

	return true, nil
}

func (v *ProductValidator) ValidateProductName(ctx context.Context, productName model.ProductName) (bool, error) {
	logger := ctx.Value(model.ContextLogger).(*log.Entry)

	productNameList := []string{
		"ExtractiveDomesticMortgage",
		"UnsecuredDomesticMortgage",
		"MidaMortgage",
		"StateMortgage",
		"PreferentialStateMortgage",
		"PSQ",
	}

	if !contains(productNameList, string(productName)) {
		logger.Errorf("ActionLog.ValidateProductName.error: product name is wrong")
		return false, errors.New(fmt.Sprintf("%s.wrong-product-name", model.Exception))
	}

	return true, nil
}

func (v *ProductValidator) ValidatePinCode(ctx context.Context, pinCode string) (bool, error) {
	logger := ctx.Value(model.ContextLogger).(*log.Entry)
	logger.Infof("ActionLog.ValidatePinCode.start: validating pincode %s", pinCode)

	if len(pinCode) != 7 {
		logger.Errorf("ActionLog.ValidatePinCode.error: pincode is not equal to 7 symbols")
		return false, errors.New(fmt.Sprintf("%s.wrong-pincode", model.Exception))
	}

	partsPinCode := strings.Fields(pinCode)
	if len(partsPinCode) > 1 {
		logger.Errorf("ActionLog.ValidatePinCode.error: pincode should consist of one word maximum")
		return false, errors.New(fmt.Sprintf("%s.max-word-in-pincode", model.Exception))
	}

	logger.Infof("ActionLog.ValidatePinCode.success: validating pincode %s", pinCode)

	return true, nil
}

func (v *ProductValidator) ValidateDateOfBirth(ctx context.Context, dateOfBirth *string) (bool, error) {
	logger := ctx.Value(model.ContextLogger).(*log.Entry)
	logger.Infof("ActionLog.ValidateDateOfBirth.start: validating date of birth %v", dateOfBirth)

	if len(*dateOfBirth) > 15 {
		logger.Errorf("ActionLog.ValidateDateOfBirth.error: date of birth is more than 15 symbols")
		return false, errors.New(fmt.Sprintf("%s.long-date-of-birth", model.Exception))
	}

	partsPinCode := strings.Fields(*dateOfBirth)
	if len(partsPinCode) > 1 {
		logger.Errorf("ActionLog.ValidateDateOfBirth.error: date of birth should consist of one word maximum")
		return false, errors.New(fmt.Sprintf("%s.max-word-in-date-of-birth", model.Exception))
	}

	logger.Infof("ActionLog.ValidateDateOfBirth.success: validating pincode %v", dateOfBirth)

	return true, nil
}

func contains(s []string, searchTerm string) bool {
	for _, name := range s {
		if name == searchTerm {
			return true
		}
	}
	return false
}
