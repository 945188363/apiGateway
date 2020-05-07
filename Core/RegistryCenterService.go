package Core

import (
	"apiGateway/Utils"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/go-plugins/registry/etcd"
	"github.com/micro/go-plugins/registry/eureka"
	"github.com/micro/go-plugins/registry/zookeeper"
)

func ConsulInitService(addr string, apiName string) []*registry.Service {
	consulReg := consul.NewRegistry(
		registry.Addrs(addr),
	)
	return serviceInit(consulReg, apiName)
}

func EtcdInitService(addr string, apiName string) []*registry.Service {
	etcdReg := etcd.NewRegistry(
		registry.Addrs(addr),
	)
	return serviceInit(etcdReg, apiName)
}

func EurekaInitService(addr string, apiName string) []*registry.Service {
	eurekaReg := eureka.NewRegistry(
		registry.Addrs(addr),
	)
	return serviceInit(eurekaReg, apiName)
}

func ZookeeperInitService(addr string, apiName string) []*registry.Service {
	zookeeperReg := zookeeper.NewRegistry(
		registry.Addrs(addr),
	)
	return serviceInit(zookeeperReg, apiName)
}

func serviceInit(reg registry.Registry, apiName string) []*registry.Service {
	servicesList, err := reg.GetService(apiName)
	if err != nil {
		Utils.RuntimeLog().Info("init registry center error .", err)
		return nil
	}
	return servicesList
}
