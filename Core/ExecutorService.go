package Core

import (
	"apiGateway/Config"
	"apiGateway/DBModels"
	"apiGateway/Utils"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"io/ioutil"
	"net/http"
)

type Executor interface {
	execute(ginCtx *gin.Context)
}

type Selector struct {
	host string
	LoadBalance
}

type LoadBalance struct {
	LBType   string
	Strategy string
	Addr     string
}

type HttpInvoker struct {
	Selector
	DBModels.Api
}

func (p *HttpInvoker) execute(ginCtx *gin.Context) {
	// 先获取该服务在注册中心的集群数量 通过负载均衡选在一个server
	p.discovery()
	// 使用httpClient调用相关服务
	switch p.ApiMethod {
	case Config.Post:
		resp, err := http.PostForm(p.host+p.BackendUrl, ginCtx.Request.Form)
		if err != nil {
			Utils.RuntimeLog().Info("send Post request error .", err)
			return
		}
		ginCtx.JSON(http.StatusAccepted, resp)
	case Config.Get:
		resp, err := http.Get(p.host + p.BackendUrl)
		if err != nil {
			Utils.RuntimeLog().Info("do Get request error .", err)
			return
		}
		ginCtx.JSON(http.StatusAccepted, resp)
	default:
		client := &http.Client{}
		req, err := http.NewRequest(p.ApiMethod, p.host+p.BackendUrl, ginCtx.Request.Body)
		if err != nil {
			Utils.RuntimeLog().Info("send request error .", err)
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(req)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Utils.RuntimeLog().Info("read body error .", err)
			return
		}
		ginCtx.JSON(http.StatusAccepted, body)
	}
}

// 选择服务
func (p *HttpInvoker) selectService(servicesList []*registry.Service) string {
	return GetAddr(servicesList, p.Strategy)
}

// 获取负载均衡数据
func (p *HttpInvoker) loadBalance() {
	lb := DBModels.LoadBalance{}
	err := lb.GetLoadBalanceByServiceName(p.ApiName)
	if err != nil {
		Utils.RuntimeLog().Info("get load balance info error .", err)
	}
	reg := DBModels.Registry{}
	reg.Name = lb.RegistryName
	err = reg.GetRegistry()
	if err != nil {
		Utils.RuntimeLog().Info("get registry info error .", err)
	}
	p.LBType = reg.RegistryType
	p.Addr = reg.Addr
	p.Strategy = lb.Strategy
}

// 获取注册中心数据
func (p *HttpInvoker) discovery() {
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
		Utils.RuntimeLog().Info("can not fetch service address .")
		return
	}
	p.host = serviceAddr
}
