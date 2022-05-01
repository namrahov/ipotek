package properties

import (
	"github.com/alexflint/go-arg"
)

// RootPath is project root path
const RootPath = "/v1/retail/products-info"

type args struct {
	LogLevel        string `arg:"env:LOG_LEVEL"`
	Port            int    `arg:"env:PORT"`
	DbHost          string `arg:"env:DB_RETAIL_PRODUCTS_INFO_HOST"`
	DbPort          string `arg:"env:DB_RETAIL_PRODUCTS_INFO_PORT"`
	DbName          string `arg:"env:DB_RETAIL_PRODUCTS_INFO_NAME"`
	DbUser          string `arg:"env:DB_RETAIL_PRODUCTS_INFO_USER"`
	DbPass          string `arg:"env:DB_RETAIL_PRODUCTS_INFO_PASS"`
	MailSenderQueue string `arg:"env:MAIL_SENDER_QUEUE"`
	RabbitMqHost    string `arg:"env:RABBITMQ_HOST"`
	RabbitMqPort    string `arg:"env:RABBITMQ_PORT"`
	RabbitMqUser    string `arg:"env:RABBITMQ_USER"`
	RabbitMqPass    string `arg:"env:RABBITMQ_PASS"`
}

// Props is for storing environment properties
var Props args

// LoadConfig loads service configuration into environment
func LoadConfig() {
	err := arg.Parse(&Props)
	if err != nil {
		return
	}
}
