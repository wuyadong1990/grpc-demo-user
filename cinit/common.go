package cinit

import (
	"log"

	"github.com/jinzhu/configor"
)

const (
	ReqParam        = "req_param"     //  请求参数绑定
	JWTName         = "Authorization" //  JWT请求头名称
	JWTMsg          = "JWT-MSG"       //  JWT自定义的消息
	FloatComputeBit = 2               //  浮点计算位数
)

// 服务名称
const (
	MySQL    = "MySQL"
	Trace    = "Trace"
	Mongo    = "Mongo"
	Redis    = "Redis"
	Kafka    = "Kafka"
	Metrics  = "Metrics"
	Postgres = "Postgres"
)

// Config 公共配置
var Config = struct {
	Service struct {
		Name      string `default:"srv-user"` //  服务名称
		Version   string `default:"v1.0"`     //  服务版本号
		RateTime  int    `default:"1024"`     //  限制请求
		AppKey    string `default:"admin"`
		AppSecret string `default:"admin"`
	}
	//  tracing
	Trace struct { //  链路跟踪
		Address       string  `default:"http://simplest-collector:14268/api/traces?format=jaeger.thrift"` //  http://jaeger:14268/api/traces?format=jaeger.thrift
		ZipkinURL     string  `default:""`                                                                //  http://zipkin:9411/api/v1/spans
		SamplingRate  float64 `default:"1"`                                                               //  采样率 0.01-1范围
		LogTraceSpans bool    `default:"false"`                                                           //  日志
	}
	//  log config
	Log struct { // 日志
		Path         string `default:"tmp"`   //  日志保存路径
		IsStdOut     string `default:"yes"`   //  是否输出日志到标准输出 yes:输出 no:不输出
		MaxAge       int    `default:"7"`     //  日志最大的保存时间，单位天
		RotationTime int    `default:"1"`     //  日志分割的时间，单位天
		MaxSize      int    `default:"100"`   //  日志分割的尺寸，单位MB
		LogLevel     string `default:"debug"` //  日志级别
	}
	//  mysql config
	Mysql struct {
		DbName   string `default:"grpcdemo"`      // 数据库名称
		Addr     string `default:"192.168.0.180"` // 地址
		User     string `default:"root"`
		Password string `default:"Wyd@123456"`
		Port     int    `default:"3306"` // required:"true" env:"DB_PROT"
		IDleConn int    `default:"4"`    // 空闲连接
		MaxConn  int    `default:"20"`   // 最大连接
	}
	// mysql config
	Postgres struct {
		DbName   string `default:"test"`      // 数据库名称
		Addr     string `default:"127.0.0.1"` // 地址
		User     string `default:"postgres"`
		Password string `default:"postgres"`
		Port     int    `default:"5432"` // required:"true"
		IDleConn int    `default:"4"`    // 空闲连接
		MaxConn  int    `default:"20"`   // 最大连接
	}
	// redis config
	Redis struct {
		Addr     string `default:"192.168.0.190:32629"` // 地址
		Password string `default:""`
		Db       int    `default:"0"`
	}
	// mongo config
	Mongo struct {
		Hosts     string `default:"127.0.0.1:27017"` // 数据库地址，可以多个，用逗号分割
		DbName    string `default:"test"`            // 数据库名称
		User      string `default:"root"`
		Password  string `default:"root"`
		PoolLimit int    `default:"4096"` // 连接池限制
	}
	Kafka struct {
		Addrs string `default:"127.0.0.1:9092"` // 数据库地址，可以多个，用逗号分割
	}
	// metrics config
	Metrics struct {
		Enable   string `default:"yes"` // 是否启用:yes 启用 no 停用
		Duration int    `default:"5"`   // 单位秒
		URL      string `default:"http://influxdb:8086"`
		Database string `default:"test01"`
		UserName string `default:""`
		Password string `default:""`
	}

	// userservice
	SrvUser struct {
		Port              string `default:":5001"`          // 定义的端口
		Address           string `default:"127.0.0.1:5001"` // 访问地址
		GateWayAddr       string `default:":9999"`          // 网关端口
		GateWaySwaggerDir string `default:"/swagger"`       //  swagger目录
	}
	// accountservice
	SrvAccount struct {
		Port              string `default:":5003"`          // 定义的端口
		Address           string `default:"127.0.0.1:5003"` // 访问地址
		GateWayAddr       string `default:":9997"`          // 网关端口
		GateWaySwaggerDir string `default:"/swagger"`       //  swagger目录
	}
	// api backend
	APIBackend struct {
		Port    string `default:":8888"`
		Address string `default:"127.0.0.1:8888"`
	}
	// api backend
	APIFrontend struct {
		Port    string `default:":8889"`
		Address string `default:"127.0.0.1:8889"`
	}
	// gamesocket
	SrvSocket struct {
		Port    string `default:":5002"`          // 定义的端口
		Address string `default:"127.0.0.1:5002"` // 访问地址
	}
}{}

// 初始化配置文件// 配置加载顺序1.是否设置了变量conf，设置了第一个加载，如果文件不存在，加载默认配置文件// 如果设置了环境变量 CONFIGOR_ENV = test等，那么加载config_test.yml的配置文件// 最后加载环境变量,是否设置环境变量前缀,如果设置了CONFIGOR_ENV_PREFIX=WEB,设置环境变量为WEB_DB_NAME=root,否则为DB_NAME=root
func configInit(sn string) {
	err := configor.Load(&Config, "config.yml")
	if err != nil {
		log.Printf("load config error:%+v", err)
	}
	if Config.Service.Name == "" {
		Config.Service.Name = sn // 使用传入的名称
	}
	log.Printf("config: %+v\n", Config)
}

// 保存需要关闭的选项
var closeArgs []string

// 初始化选项// log:日志(必须) trace:链路跟踪 mysql:mysql数据库 mongo:MongoDB postgres:postgres数据库
func InitOption(sn string, args ...string) {
	// 开启pprof
	//go pprof.Run()
	// 保存需要关闭的参数
	closeArgs = args
	// 1.初始化配置参数
	configInit(sn)
	// 2.初始化日志
	//logInit()
	// 3.其他服务
	for _, o := range args {
		switch o {
		case Trace:
			traceInit()
		case MySQL:
			mysqlInit()
		case Mongo:
		case Redis:
			redisInit()
			//case Kafka:
			//KafkaInit()
			//case Metrics:
			//metricsInit(sn)
			//case Postgres:
			//pgInit()
		}
	}
}

// 关闭打开的服务
func Close() {
	for _, o := range closeArgs {
		switch o {
		//case Trace:
		// 关闭链路跟踪
		//	tracerClose()
		case MySQL:
			// 关闭mysql
			mysqlClose()
		case Mongo:
		case Redis:
			redisClose()
			//case Kafka:
			//KafkaClose()
			//case Metrics:
			//case Postgres:
			//pgClose()
		}
	}
}
