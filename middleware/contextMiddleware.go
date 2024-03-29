package middleware

import (
	"context"
	"github.com/PB-Digital/ms-retail-products-info/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var headers = []string{
	"x-request-id",
	"x-b3-traceid",
	"x-b3-spanid",
	"x-b3-parentspanid",
	"x-b3-sampled",
	"x-b3-flags",
	"x-ot-span-context",
	"DP-Customer-ID",
	"DP-User-ID",
	"User-Agent",
	"X-Forwarded-For",
	"requestid",
}

// RequestParamsMiddleware is middleware function for context time logger and header transport
func RequestParamsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Headers", "*")

		ctx := r.Context()

		requestID := r.Header.Get(model.HeaderKeyRequestID)
		operation := r.RequestURI
		customerID := r.Header.Get(model.HeaderKeyCustomerID)
		userID := r.Header.Get(model.HeaderKeyUserID)
		userAgent := r.Header.Get(model.HeaderKeyUserAgent)
		userIP := r.Header.Get(model.HeaderKeyUserIP)

		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}
		fields := log.Fields{}
		addLoggerParam(fields, model.LoggerKeyRequestID, requestID)
		addLoggerParam(fields, model.LoggerKeyCustomerID, customerID)
		addLoggerParam(fields, model.LoggerKeyOperation, operation)
		addLoggerParam(fields, model.LoggerKeyUserAgent, userAgent)
		addLoggerParam(fields, model.LoggerKeyUserID, userID)
		addLoggerParam(fields, model.LoggerKeyUserIP, userIP)

		logger := log.WithFields(fields)
		header := http.Header{}

		for _, v := range headers {
			header.Add(v, r.Header.Get(v))
		}

		ctx = context.WithValue(ctx, model.ContextLogger, logger)
		ctx = context.WithValue(ctx, model.ContextHeader, header)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func addLoggerParam(fields log.Fields, field string, value string) {
	if len(value) > 0 {
		fields[field] = value
	}
}
