#!/bin/bash

if [ -f /vault/secrets/secrets ]
then
  echo "Loading values from secrets file"
  source /vault/secrets/secrets
  export DB_RETAIL_PRODUCTS_INFO_HOST
  export DB_RETAIL_PRODUCTS_INFO_PORT
  export DB_RETAIL_PRODUCTS_INFO_NAME
  export DB_RETAIL_PRODUCTS_INFO_USER
  export DB_RETAIL_PRODUCTS_INFO_PASS
  export RABBITMQ_HOST
  export RABBITMQ_PORT
  export RABBITMQ_USER
  export RABBITMQ_PASS
  sudo rm /vault/secrets/secrets
fi

sleep 5

exec ./ms-retail-products-info
