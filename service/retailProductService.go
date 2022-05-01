package service

import (
	"context"
	"fmt"
	"github.com/PB-Digital/ms-retail-products-info/mapper"
	"github.com/PB-Digital/ms-retail-products-info/mapper/mail"
	"github.com/PB-Digital/ms-retail-products-info/model"
	"github.com/PB-Digital/ms-retail-products-info/properties"
	"github.com/PB-Digital/ms-retail-products-info/queue"
	"github.com/PB-Digital/ms-retail-products-info/repo"
	"github.com/PB-Digital/ms-retail-products-info/util"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
)

type IRetailProductService interface {
	SaveClientQuestionRequest(dto *model.ClientQuestionDto, ctx context.Context) *model.ErrorResponse
	SaveClientRequest(dto *model.ClientDto, ctx context.Context) *model.ErrorResponse
}

type RetailProductService struct {
	RetailProductRepo repo.IRetailProductRepo
	Validator         util.IProductValidator
	MessageSender     queue.IMessageSender
}

func (s *RetailProductService) SaveClientQuestionRequest(dto *model.ClientQuestionDto, ctx context.Context) *model.ErrorResponse {
	logger := ctx.Value(model.ContextLogger).(*log.Entry)
	logger.Infof("ActionLog.SaveClientQuestionRequest.start")

	e := s.Validator.ValidateForClientQuestionRequest(ctx, dto)
	if e != nil {
		logger.Errorf("ActionLog.SaveClientQuestionRequest.error :validation error client question request dto = %v", dto)
		return &model.ErrorResponse{
			Code:   e.Error(),
			Status: http.StatusBadRequest,
		}
	}

	existingClientQuestionRequest, err :=
		s.RetailProductRepo.GetClientQuestionRequestByNameAndSurnameAndPhoneAndCreatedInLast1Hour(dto)
	if existingClientQuestionRequest.Id != 0 {
		logger.Errorf("ActionLog.SaveClientQuestionRequest.error: client question request already exist %v", dto)
		return &model.ErrorResponse{
			Code:   fmt.Sprintf("%s.client-question-request-already-exist", model.Exception),
			Status: http.StatusBadRequest,
		}
	}

	product, err := s.RetailProductRepo.GetProductByName(string(dto.ProductName))
	if err != nil {
		logger.Errorf("ActionLog.SaveClientQuestionRequest.error: cannot get product for %v, with error %v", dto, err)
		return &model.ErrorResponse{
			Code:   fmt.Sprintf("%s.could-not-get-product", model.Exception),
			Status: http.StatusBadRequest,
		}
	}

	clientQuestionRequest := mapper.BuildClientQuestionRequest(dto, product.Id)

	err = s.RetailProductRepo.SaveClientQuestionRequest(&clientQuestionRequest)
	if err != nil {
		logger.Errorf("ActionLog.SaveClientQuestionRequest.error : can't save "+
			" client question request for %v card %v", clientQuestionRequest, err)
		return &model.ErrorResponse{
			Code:   fmt.Sprintf("%s.could-not-save-client-question-request", model.Exception),
			Status: http.StatusInternalServerError,
		}
	}

	if product.Email != nil {
		mailDto := mail.PrepareMail(dto, *product.Email)
		err = s.MessageSender.SendMessage(mailDto, properties.Props.MailSenderQueue)
		if err != nil {
			log.Errorf("ActionLog.SaveClientQuestionRequest.error : Error can't sent message to outgoingMail queue: %v,\n%s",
				err, string(debug.Stack()))
			return &model.ErrorResponse{
				Code:   fmt.Sprintf("%s.could-not-send-client-question-request-email", model.Exception),
				Status: http.StatusInternalServerError,
			}
		}
	}

	logger.Info("ActionLog.SaveClientQuestionRequest.end")
	return nil
}

func (s *RetailProductService) SaveClientRequest(dto *model.ClientDto, ctx context.Context) *model.ErrorResponse {
	logger := ctx.Value(model.ContextLogger).(*log.Entry)
	logger.Info("ActionLog.SaveClientRequest.start")

	e := s.Validator.ValidateForClientRequest(ctx, dto)
	if e != nil {
		logger.Errorf("ActionLog.SaveClientRequest.error :validation error client request dto = %v", dto)
		return &model.ErrorResponse{
			Code:   e.Error(),
			Status: http.StatusBadRequest,
		}
	}

	existingClientRequest, err :=
		s.RetailProductRepo.GetClientRequestCreatedInLast1Hour(dto)
	if existingClientRequest.Id != 0 {
		logger.Errorf("ActionLog.SaveClientRequest.error: client request already exist %v", dto)
		return &model.ErrorResponse{
			Code:   fmt.Sprintf("%s.client-request-already-exist", model.Exception),
			Status: http.StatusBadRequest,
		}
	}

	product, err := s.RetailProductRepo.GetProductByName(string(dto.ProductName))
	if err != nil {
		logger.Errorf("ActionLog.SaveClientRequest.error: cannot get product for %v, with error %v", dto, err)
		return &model.ErrorResponse{
			Code:   fmt.Sprintf("%s.could-not-get-product", model.Exception),
			Status: http.StatusBadRequest,
		}
	}

	clientRequest := mapper.BuildClientRequest(dto, product.Id)

	clientId, err := s.RetailProductRepo.SaveClientRequest(&clientRequest)
	if err != nil {
		logger.Errorf("ActionLog.SaveClientRequest.error: can't save client request for %v with error %v", clientRequest, err)
		return &model.ErrorResponse{
			Code:   fmt.Sprintf("%s.could-not-save-client-request", model.Exception),
			Status: http.StatusInternalServerError,
		}
	}

	if dto.ClientAgreementDto != nil {
		clientAgreement := mapper.BuildClientAgreement(dto.ClientAgreementDto, clientId)

		err = s.RetailProductRepo.SaveClientAgreement(&clientAgreement)
		if err != nil {
			logger.Errorf("ActionLog.SaveClientRequest.error : can't save client agreement for %v with error %v", clientAgreement, err)
			return &model.ErrorResponse{
				Code:   fmt.Sprintf("%s.could-not-save-client-agreement", model.Exception),
				Status: http.StatusInternalServerError,
			}
		}
	}

	if product.Email != nil {
		var mailDto *mail.Mail
		if dto.ProductName != model.PSQ {
			mailDto = mail.PrepareAppealMail(dto, *product.Email)
		} else {
			mailDto = mail.PrepareAppealMailForPashaGroup(dto, *product.Email)
		}

		err = s.MessageSender.SendMessage(mailDto, properties.Props.MailSenderQueue)
		if err != nil {
			log.Errorf("ActionLog.SaveClientRequest.error : Error can't sent message to outgoingMail queue: %v,\n%s",
				err, string(debug.Stack()))
			return &model.ErrorResponse{
				Code:   fmt.Sprintf("%s.could-not-send-client-appeal-request-email", model.Exception),
				Status: http.StatusInternalServerError,
			}
		}
	}

	logger.Info("ActionLog.SaveClientRequest.end")
	return nil
}
