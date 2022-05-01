module github.com/PB-Digital/ms-retail-products-info

go 1.16

require (
	github.com/alexflint/go-arg v1.4.2
	github.com/coreos/go-etcd v2.0.0+incompatible // indirect
	github.com/cpuguy83/go-md2man v1.0.10 // indirect
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/go-pg/pg v8.0.7+incompatible
	github.com/gobuffalo/envy v1.7.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/instana/go-sensor v1.35.0
	github.com/instana/go-sensor/instrumentation/instamux v1.0.0
	github.com/jessevdk/go-flags v1.5.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.10.4
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/prometheus/client_golang v1.9.0
	github.com/prometheus/common v0.26.0 // indirect
	github.com/rubenv/sql-migrate v1.1.0
	github.com/sirupsen/logrus v1.8.1
	github.com/streadway/amqp v1.0.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8 // indirect
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	gopkg.in/gorp.v1 v1.7.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	mellium.im/sasl v0.2.1 // indirect
)

replace go.etcd.io/etcd => github.com/coreos/etcd v3.3.25+incompatible

replace (
	github.com/apache/thrift v0.12.0 => github.com/apache/thrift v0.13.0
	github.com/aws/aws-sdk-go v1.27.0 => github.com/aws/aws-sdk-go v1.34.2
	github.com/coreos/etcd => go.etcd.io/etcd/v3 v3.5.0-alpha.0
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket v0.0.0-20170926233335-4201258b820c => github.com/gorilla/websocket v1.4.1
	github.com/miekg/dns v1.0.14 => github.com/miekg/dns v1.1.25
	github.com/nats-io/nats-server/v2 => github.com/nats-io/nats-server/v2 v2.2.0
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20201203163018-be400aefbc4c
	golang.org/x/text => golang.org/x/text v0.3.3
	gopkg.in/yaml.v2 v2.2.5 => gopkg.in/yaml.v2 v2.2.8
)
