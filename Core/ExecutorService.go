package Core

import (
	"apiGateway/Config"
	"apiGateway/DBModels"
	"apiGateway/Utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"io/ioutil"
	"net/http"
	"time"
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

func (p *HttpInvoker) Execute(ginCtx *gin.Context) {
	// 先获取该服务在注册中心的集群数量 通过负载均衡选在一个server
	p.discovery()
	// 接收返回数据的chan，用在go程返回
	doneChan := make(chan Message)
	// 使用模板调用相关服务
	// MethodExecute(p,ginCtx,doneChan)

	// 获取到已经wrapper超时的上下文
	ctx := ginCtx.Request.Context()
	// 使用httpClient调用相关服务
	switch p.ApiMethod {
	case Config.Post:
		go func() {
			resp, err := http.PostForm(handleProtocol(p.ProtocolType, p.host, p.BackendUrl), ginCtx.Request.Form)
			if err != nil {
				Utils.RuntimeLog().Info("do Get request error .", err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				Utils.RuntimeLog().Info("read body error .", err)
				return
			}
			var msg Message
			err = json.Unmarshal(body, &msg)
			if err != nil {
				Utils.RuntimeLog().Info("json transfer error .", err)
				return
			}
			doneChan <- msg
		}()
		// 监听通道是否超时或者完成
		select {
		// 如果上下文超时，或者被cancel则不返回数据
		case <-ctx.Done():
			return
		// 如果调用完成则写数据
		case res := <-doneChan:
			ginCtx.JSON(res.Code, res)
		}
	case Config.Get:
		go func() {
			resp, err := http.Get(handleProtocol(p.ProtocolType, p.host, p.BackendUrl))
			if err != nil {
				Utils.RuntimeLog().Info("do Get request error .", err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				Utils.RuntimeLog().Info("read body error .", err)
				return
			}
			var msg Message
			err = json.Unmarshal(body, &msg)
			if err != nil {
				Utils.RuntimeLog().Info("json transfer error .", err)
				return
			}
			doneChan <- msg
		}()
		// 监听通道是否超时或者完成
		select {
		// 如果上下文超时，或者被cancel则不返回数据
		case <-ctx.Done():
			return
		// 如果调用完成则写数据
		case res := <-doneChan:
			ginCtx.JSON(res.Code, res)
		}
	default:
		client := &http.Client{
			Timeout: time.Duration(p.ApiTimeout),
		}
		req, err := http.NewRequest(p.ApiMethod, handleProtocol(p.ProtocolType, p.host, p.BackendUrl), ginCtx.Request.Body)
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
		ginCtx.JSON(http.StatusAccepted, string(body))
	}
}

// 模板方法处理
func MethodExecute(p *HttpInvoker, ginCtx *gin.Context, doneChan chan Message) {
	// 获取到已经wrapper超时的上下文
	ctx := ginCtx.Request.Context()
	go func() {
		var err error
		var resp *http.Response
		switch p.ApiMethod {
		case Config.Post:
			resp, err = http.PostForm(handleProtocol(p.ProtocolType, p.host, p.BackendUrl), ginCtx.Request.Form)
			break
		case Config.Get:
			resp, err = http.Get(handleProtocol(p.ProtocolType, p.host, p.BackendUrl))
			break
		default:
			client := &http.Client{
				Timeout: time.Duration(p.ApiTimeout) * time.Millisecond,
			}
			req, err := http.NewRequest(p.ApiMethod, handleProtocol(p.ProtocolType, p.host, p.BackendUrl), ginCtx.Request.Body)
			if err != nil {
				Utils.RuntimeLog().Info("send request error .", err)
				return
			}
			resp, err = client.Do(req)
		}
		// 错误处理
		if err != nil {
			Utils.RuntimeLog().Info("send request error .", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Utils.RuntimeLog().Info("read body error .", err)
			return
		}
		var msg Message
		err = json.Unmarshal(body, &msg)
		if err != nil {
			Utils.RuntimeLog().Info("json transfer error .", err)
			return
		}
		doneChan <- msg
	}()
	// 监听通道是否超时或者完成
	select {
	// 如果上下文超时，或者被cancel则不返回数据
	case <-ctx.Done():
		return
	// 如果调用完成则写数据
	case res := <-doneChan:
		ginCtx.JSON(res.Code, res)
	}
}

// 处理协议
func handleProtocol(protocolType, host, uri string) string {
	var addr string

	// 处理协议
	addr = protocolType + "://" + host + uri
	return addr
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
