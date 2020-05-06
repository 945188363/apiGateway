package Config

const (
	// 存储周期
	Day  = "day"
	Hour = "hour"

	// 日志类型
	AccessLog  = 1
	RuntimeLog = 2

	// 监控类型
	ELK        = "ELK"
	Prometheus = "Prometheus"

	// 插件类型
	API      = "API"
	Strategy = "策略"

	// MQ配置
	MqUrl          = ""
	MqUsername     = ""
	MqPassword     = ""
	MqExchange     = ""
	MqExchangeType = ""
	MqVirtualHost  = ""
	MqQueueName    = ""
	MqRoutingKey   = ""

	// ES配置
	EsUrl   = ""
	EsIndex = ""
	EsHost  = ""
)
