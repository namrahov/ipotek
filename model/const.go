package model

const (
	HeaderKeyCustomerID = "DP-Customer-ID"
	HeaderKeyUserID     = "DP-User-ID"
	HeaderKeyUserAgent  = "User-Agent"
	HeaderKeyUserIP     = "X-Forwarded-For"
	HeaderKeyRequestID  = "requestid"
)
const (
	LoggerKeyRequestID  = "REQUEST_ID"
	LoggerKeyOperation  = "OPERATION"
	LoggerKeyCustomerID = "CUSTOMER_ID"
	LoggerKeyUserID     = "USER_ID"
	LoggerKeyUserIP     = "USER_IP"
	LoggerKeyUserAgent  = "USER_AGENT"
	ContextLogger       = "contextLogger"
	ContextHeader       = "contextHeader"
)

const (
	Exception = "error.retail-products-info"
)
