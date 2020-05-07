package Config

const (

	// 随机
	Random = "Random"
	// 加权随机
	WeightedRandom = "WeightedRandom"

	// 轮询
	RoundRobin = "RoundRobin"
	// 加权轮询
	WeightedRoundRobin = "WeightedRoundRobin"
	// 平滑加权轮询
	SmoothlyWeightedRoundRobin = "SmoothlyWeightedRoundRobin"

	// 最少连接
	LeastConnection = "LeastConnection"
	// 加权最少连接
	WeightedLeastConnection = "WeightedLeastConnection"

	// 注册中心配置
	Etcd      = "etcd"
	Nacos     = "nacos"
	Consul    = "consul"
	Zookeeper = "zookeeper"
	Eureka    = "eureka"
)
