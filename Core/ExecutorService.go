package Core

import (
	"apiGateway/Config"
	"apiGateway/Constant/Code"
	"apiGateway/Constant/Message"
	"apiGateway/Core/Domain"
	"apiGateway/Core/rpcService"
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
	"apiGateway/Utils/DataUtil"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"net/http"
	"time"
)

type Executor interface {
	Execute(ginCtx *gin.Context)
}

type Invoker struct {
	Selector
	DBModels.Api
}

type HttpInvoker struct {
	Invoker
}

type RpcService interface {
	InvokeMethod(ctx context.Context, in *Domain.RpcRequest, opts ...client.CallOption) (*Domain.RpcResponse, error)
	InvokeMethod2(ctx context.Context, in *rpcService.RpcRequest, opts ...client.CallOption) (*rpcService.RpcResponse, error)
}

type RpcInvoker struct {
	DBModels.Api
	c client.Client
}

func (c *RpcInvoker) InvokeMethod(ctx context.Context, in *Domain.RpcRequest, opts ...client.CallOption) (*Domain.RpcResponse, error) {
	req := c.c.NewRequest(c.ApiName, c.BackendUrl, in)
	out := new(Domain.RpcResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *RpcInvoker) InvokeMethod2(ctx context.Context, in *rpcService.RpcRequest, opts ...client.CallOption) (*rpcService.RpcResponse, error) {
	req := c.c.NewRequest(c.ApiName, c.BackendUrl, in)
	out := new(rpcService.RpcResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// gRpc协议的RPC调用
func (p *RpcInvoker) Execute(ginCtx *gin.Context) {
	value, exists := ginCtx.Get("ApiInfo")
	if exists {
		p.Api = value.(DBModels.Api)
	}
	doneChan := make(chan Domain.Message)
	go func(ctx context.Context) {
		// 获取consul注册地址
		consulReg := consul.NewRegistry(
			registry.Addrs("localhost:8500"),
		)
		// 负载均衡选择器
		mySelector := selector.NewSelector(
			selector.Registry(consulReg),
			selector.SetStrategy(selector.Random),
		)
		// 创建客户端
		serviceClient := micro.NewService(
			micro.Name(p.ApiName+".Client"),
			micro.Selector(mySelector),
		)
		rpcServ := NewRpcService(p.Api, serviceClient.Client())
		// p.c = serviceClient.Client()

		// req := Domain.NewRpcRequest()
		// req.Request["size"] = 2
		// reqMessage := Domain.Message{}
		// reqMessage.Data = make(map[string]interface{})
		// reqMessage.Data["Size"] = 8
		paramMap := make(map[string]interface{})
		paramMap["Size"] = 22
		req := rpcService.RpcRequest{
			// Request:Utils.MessageToBytes(reqMessage),
			Request: DataUtil.IntToBytes(8),
			// Request:Utils.MapToBytes(paramMap),
		}

		// 通过包装callOptions来设置重试、超时等
		// resp, err := rpcServ.InvokeMethod(ctx, &req,
		// 	client.WithDialTimeout(time.Duration(30000) * time.Millisecond),
		// 	client.WithRequestTimeout(time.Duration(30000) * time.Millisecond),
		resp, err := rpcServ.InvokeMethod2(ctx, &req,
			client.WithDialTimeout(time.Duration(30000)*time.Millisecond),
			client.WithRequestTimeout(time.Duration(30000)*time.Millisecond),
		)
		// resp, err := rpcServ.InvokeMethod(ctx, &req)
		// resp, err := p.InvokeMethod(ctx, &req)
		if err != nil {
			fmt.Println(err.Error())
			// Utils.RuntimeLog().Warn("Invoke Rpc Service error,error:", err.Error())
			return
		}
		var msg Domain.Message
		msg.Code = Code.RPC_SUCCESS
		msg.Msg = Message.RPC_SUCCESS
		msg.Data = make(map[string]interface{})
		msg.Data["Num"] = DataUtil.BytesToInt(resp.Response)
		doneChan <- msg
	}(context.Background())

	// 获取到已经wrapper超时的上下文
	ctx := ginCtx.Request.Context()
	select {
	// 如果上下文超时，或者被cancel则不返回数据
	case <-ctx.Done():
		return
	// 如果调用完成则写数据
	case res := <-doneChan:
		handleResponseReturn(res, p.ApiReturnType, ginCtx)
	}
}

func NewRpcService(api DBModels.Api, c client.Client) RpcService {
	if c == nil {
		c = client.NewClient()
	}
	// 客户端重试和超时时间
	// _ = c.Init(
	// 	client.Retries(api.ApiRetry),
	// 	client.DialTimeout(6*time.Second),
	// 	client.RequestTimeout(6*time.Second),
	// 	client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
	// 		// Utils.RuntimeLog().Warn(req.Method(), retryCount, " client retry")
	// 		fmt.Println(req.Method(), retryCount, " client retry")
	// 		return true, nil
	// 	}),
	// )
	return &RpcInvoker{
		Api: api,
		c:   c,
	}
}

// 协议的RPC调用
func (p *HttpInvoker) Execute(ginCtx *gin.Context) {
	// 获取api信息
	value, exists := ginCtx.Get("ApiInfo")
	if exists {
		p.Api = value.(DBModels.Api)
	}

	// 先获取该服务在注册中心的集群数量 通过负载均衡选在一个server
	p.discovery()

	// 接收返回数据的chan，用在go程返回
	doneChan := make(chan Domain.Message)

	// 使用模板调用相关服务
	// MethodExecute(p,ginCtx,doneChan)

	// 获取到已经wrapper超时的上下文
	ctx := ginCtx.Request.Context()
	// 使用httpClient调用相关服务
	switch p.ApiMethod {
	case Config.Post:
		go func(ctx context.Context) {
			resp, err := http.PostForm(handleProtocol(p.ProtocolType, p.host, p.BackendUrl), ginCtx.Request.Form)
			if err != nil {
				ComponentUtil.RuntimeLog().Info("do Get request error .", err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				ComponentUtil.RuntimeLog().Info("read body error .", err)
				return
			}
			var msg Domain.Message
			err = json.Unmarshal(body, &msg)
			if err != nil {
				ComponentUtil.RuntimeLog().Info("json transfer error .", err)
				return
			}
			doneChan <- msg
		}(ctx)
		// 监听通道是否超时或者完成
		select {
		// 如果上下文超时，或者被cancel则不返回数据
		case <-ctx.Done():
			return
		// 如果调用完成则写数据
		case res := <-doneChan:
			handleResponseReturn(res, p.ApiReturnType, ginCtx)
		}
	case Config.Get:
		go func(ctx context.Context) {
			resp, err := http.Get(handleProtocol(p.ProtocolType, p.host, p.BackendUrl))
			if err != nil {
				ComponentUtil.RuntimeLog().Info("do Get request error .", err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				ComponentUtil.RuntimeLog().Info("read body error .", err)
				return
			}
			var msg Domain.Message
			err = json.Unmarshal(body, &msg)
			if err != nil {
				ComponentUtil.RuntimeLog().Info("json transfer error .", err)
				return
			}
			doneChan <- msg
		}(ctx)
		// 监听通道是否超时或者完成
		select {
		// 如果上下文超时，或者被cancel则不返回数据
		case <-ctx.Done():
			return
		// 如果调用完成则写数据
		case res := <-doneChan:
			handleResponseReturn(res, p.ApiReturnType, ginCtx)
		}
	default:
		client := &http.Client{
			Timeout: time.Duration(p.ApiTimeout),
		}
		req, err := http.NewRequest(p.ApiMethod, handleProtocol(p.ProtocolType, p.host, p.BackendUrl), ginCtx.Request.Body)
		if err != nil {
			ComponentUtil.RuntimeLog().Info("send request error .", err)
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(req)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ComponentUtil.RuntimeLog().Info("read body error .", err)
			return
		}
		ginCtx.JSON(http.StatusAccepted, string(body))
	}
}

// 模板方法处理
func MethodExecute(p *HttpInvoker, ginCtx *gin.Context, doneChan chan Domain.Message) {
	// 获取到已经wrapper超时的上下文
	ctx := ginCtx.Request.Context()
	go func(ctx context.Context) {
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
				ComponentUtil.RuntimeLog().Info("send request error .", err)
				return
			}
			resp, err = client.Do(req)
		}
		// 错误处理
		if err != nil {
			ComponentUtil.RuntimeLog().Info("send request error .", err)
			return
		}
		msg := handlePreResponse(resp)
		doneChan <- msg
	}(ctx)
	// 监听通道是否超时或者完成
	select {
	// 如果上下文超时，或者被cancel则不返回数据
	case <-ctx.Done():
		return
	// 如果调用完成则写数据
	case res := <-doneChan:
		handleResponseReturn(res, p.ApiReturnType, ginCtx)
	}
}

// 处理协议
func handleProtocol(protocolType, host, uri string) string {
	var addr string

	// 处理协议
	if protocolType == Config.Http {
		addr = protocolType + "://" + host + uri
	}
	return addr
}

func handlePreResponse(resp *http.Response) Domain.Message {
	defer resp.Body.Close()
	var msg Domain.Message
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ComponentUtil.RuntimeLog().Info("read body error .", err)
		return msg
	}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		ComponentUtil.RuntimeLog().Info("json transfer error .", err)
		return msg
	}
	return msg
}

func handleResponseReturn(res Domain.Message, returnType string, ginCtx *gin.Context) {
	switch returnType {
	case Config.Raw:
		rawBytes, _ := json.Marshal(res)
		ginCtx.JSON(res.Code, string(rawBytes))
		break
	case Config.Yaml:
		ginCtx.YAML(res.Code, res)
		break
	case Config.Json:
		ginCtx.JSON(res.Code, res)
		break
	case Config.Xml:
		rawBytes, _ := json.Marshal(res)
		dataStr := string(rawBytes)
		ginCtx.XML(res.Code, Domain.NewMessageXml(res.Code, res.Msg, dataStr))
		break
	default:
		ginCtx.JSON(res.Code, res)
	}
}
