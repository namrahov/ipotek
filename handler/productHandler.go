package handler

import (
	"encoding/json"
	"github.com/PB-Digital/ms-retail-products-info/middleware"
	"github.com/PB-Digital/ms-retail-products-info/model"
	"github.com/PB-Digital/ms-retail-products-info/properties"
	"github.com/PB-Digital/ms-retail-products-info/queue"
	"github.com/PB-Digital/ms-retail-products-info/repo"
	"github.com/PB-Digital/ms-retail-products-info/service"
	"github.com/PB-Digital/ms-retail-products-info/util"
	mid "github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type productHandler struct {
	service service.IRetailProductService
}

var productService = service.RetailProductService{
	RetailProductRepo: &repo.RetailProductRepo{},
	Validator:         &util.ProductValidator{},
	MessageSender:     &queue.MessageSender{},
}

func ProductHandler(router *mux.Router) *mux.Router {
	router.Use(mid.Recoverer)
	router.Use(middleware.RequestParamsMiddleware)
	router.Use(middleware.Throttle)

	h := &productHandler{service: &productService}

	router.HandleFunc(properties.RootPath+"/client-question-request", h.saveClientQuestionRequest).Methods("POST")
	router.HandleFunc(properties.RootPath+"/client-request", h.saveClientRequest).Methods("POST")

	return router
}

func (h *productHandler) saveClientQuestionRequest(w http.ResponseWriter, r *http.Request) {
	clientQuestionDto, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("failed to parse the model", clientQuestionDto)
		return
	}

	dto := new(model.ClientQuestionDto)
	err = json.Unmarshal(clientQuestionDto, dto)
	if err != nil {
		log.Error("failed to parse the model", clientQuestionDto)
		return
	}

	errorResponse := h.service.SaveClientQuestionRequest(dto, r.Context())
	if errorResponse != nil {
		w.WriteHeader(errorResponse.Status)
		e := json.NewEncoder(w).Encode(errorResponse)
		if e != nil {
			return
		}
		return
	}

	w.WriteHeader(201)
}

func (h *productHandler) saveClientRequest(w http.ResponseWriter, r *http.Request) {
	clientDto, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("failed to parse the model", clientDto)
		return
	}

	dto := new(model.ClientDto)
	err = json.Unmarshal(clientDto, dto)
	if err != nil {
		log.Error("failed to parse the model", clientDto)
		return
	}

	errorResponse := h.service.SaveClientRequest(dto, r.Context())
	if errorResponse != nil {
		w.WriteHeader(errorResponse.Status)
		e := json.NewEncoder(w).Encode(errorResponse)
		if e != nil {
			return
		}
		return
	}

	w.WriteHeader(201)
}
