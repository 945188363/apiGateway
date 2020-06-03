module apiGateway

go 1.14

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/didip/tollbooth v4.0.2+incompatible
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/franela/goblin v0.0.0-20200409142057-1def193310bb // indirect
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/gin v1.6.2
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/lestrrat-go/file-rotatelogs v2.3.0+incompatible
	github.com/lestrrat-go/strftime v1.0.1 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.2.1
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/shirou/gopsutil v2.20.5+incompatible
	github.com/sirupsen/logrus v1.6.0
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71 // indirect
	github.com/tebeka/strftime v0.1.4 // indirect
	github.com/vladoatanasov/logrus_amqp v0.0.0-20181023103017-b21faf6f8ae3
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2 // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	google.golang.org/genproto v0.0.0-20200527145253-8367513e4ece // indirect
	google.golang.org/grpc v1.29.1 // indirect
	google.golang.org/protobuf v1.24.0
	gopkg.in/olivere/elastic.v5 v5.0.85
	gopkg.in/sohlich/elogrus.v2 v2.0.2
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
