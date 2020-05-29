package Core

import (
	"apiGateway/Config"
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
	"github.com/micro/go-micro/registry"
)

type Selector struct {
	host string
	LoadBalance
}

type LoadBalance struct {
	LBType   string
	Strategy string
	Addr     string
}

// 选择服务
func (p *Invoker) selectService(servicesList []*registry.Service) string {
	return GetAddr(servicesList, p.Strategy)
}

// 获取负载均衡数据
func (p *Invoker) loadBalance() {
	lb := DBModels.LoadBalance{}
	err := lb.GetLoadBalanceByServiceName(p.ApiName)
	if err != nil {
		ComponentUtil.RuntimeLog().Info("get load balance info error .", err)
	}
	reg := DBModels.Registry{}
	reg.Name = lb.RegistryName
	err = reg.GetRegistry()
	if err != nil {
		ComponentUtil.RuntimeLog().Info("get registry info error .", err)
	}
	p.LBType = reg.RegistryType
	p.Addr = reg.Addr
	p.Strategy = lb.Strategy
}

// 获取注册中心数据
func (p *Invoker) discovery() {
	p.loadBalance()
	var servicesList []*registry.Service
	switch p.LBType {

	case Config.Consul:
		servicesList = ConsulInitService(p.Addr, p.ApiName)
		break
	case Config.Etcd:
		servicesList = EtcdInitService(p.Addr, p.ApiName)
		break
	case Config.Zookeeper:
		servicesList = ZookeeperInitService(p.Addr, p.ApiName)
		break
	case Config.Eureka:
		servicesList = EurekaInitService(p.Addr, p.ApiName)
		break
	default:
		servicesList = ConsulInitService(p.Addr, p.ApiName)
	}
	// 获取服务地址
	serviceAddr := p.selectService(servicesList)
	if serviceAddr == "" {
		ComponentUtil.RuntimeLog().Info("can not fetch service address .")
		return
	}
	p.host = serviceAddr
}
