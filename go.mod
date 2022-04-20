module github.com/wuyadong1990/grpc-demo-user

go 1.17

//replace github.com/wuyadong1990/grpc-demo-proto => ../grpc-demo-proto

//replace github.com/wuyadong1990/grpc-demo-user => ../grpc-demo-user

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/jinzhu/configor v1.2.1
	github.com/jinzhu/copier v0.3.5
	github.com/jmoiron/sqlx v1.3.5
	github.com/mitchellh/mapstructure v1.4.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/wuyadong1990/grpc-demo-proto v0.0.0-00010101000000-000000000000
	github.com/xiaomeng79/go-log v2.0.4+incompatible
	github.com/xiaomeng79/go-utils v0.0.0-20181017074258-e4925b865baa
	google.golang.org/grpc v1.45.0
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.19.0 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220317150908-0efb43f6373e // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
