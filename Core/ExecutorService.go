package Core

import (
	"apiGateway/Config"
	"apiGateway/Constant/Code"
	"apiGateway/Constant/Message"
	"apiGateway/Core/Domain"
	"apiGateway/DBModels"
	"apiGateway/Utils"
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
}

type RpcInvoker struct {
	DBModels.Api
	c client.Client
}

func (c *RpcInvoker) InvokeMethod(ctx context.Context, in *Domain.RpcRequest, opts ...client.CallOption) (*Domain.RpcResponse, error) {
	req := c.c.NewRequest(c.ApiName, c.ApiName+"."+c.BackendUrl, in)
	out := new(Domain.RpcResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// gRpc协议的RPC调用
func (p *RpcInvoker) Execute(ginCtx *gin.Context) {

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
		p.c = serviceClient.Client()

		req := Domain.NewRpcRequest()
		req.Request["size"] = 2
		resp, err := p.InvokeMethod(ctx, &req)
		if err != nil {
			fmt.Println(err.Error())
			// Utils.RuntimeLog().Warn("Invoke Rpc Service error,error:", err.Error())
			return
		}
		message := Domain.Message{
			Code: Code.RPC_SUCCESS,
			Msg:  Message.RPC_SUCCESS,
			Data: resp.Response,
		}
		doneChan <- message
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

// http协议的RPC调用
func (p *HttpInvoker) Execute(ginCtx *gin.Context) {
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
				Utils.RuntimeLog().Info("do Get request error .", err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				Utils.RuntimeLog().Info("read body error .", err)
				return
			}
			var msg Domain.Message
			err = json.Unmarshal(body, &msg)
			if err != nil {
				Utils.RuntimeLog().Info("json transfer error .", err)
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
				Utils.RuntimeLog().Info("do Get request error .", err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				Utils.RuntimeLog().Info("read body error .", err)
				return
			}
			var msg Domain.Message
			err = json.Unmarshal(body, &msg)
			if err != nil {
				Utils.RuntimeLog().Info("json transfer error .", err)
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
	addr = protocolType + "://" + host + uri
	return addr
}

func handlePreResponse(resp *http.Response) Domain.Message {
	defer resp.Body.Close()
	var msg Domain.Message
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Utils.RuntimeLog().Info("read body error .", err)
		return msg
	}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		Utils.RuntimeLog().Info("json transfer error .", err)
		return msg
	}
	return msg
}

func handleResponseReturn(res Domain.Message, returnType string, ginCtx *gin.Context) {
	switch returnType {
	case Config.Raw:
		resBytes, _ := json.Marshal(res)
		ginCtx.JSON(res.Code, string(resBytes))
		break
	case Config.Yaml:
		ginCtx.YAML(res.Code, res)
		break
	case Config.Json:
		ginCtx.JSON(res.Code, res)
		break
	case Config.Xml:
		ginCtx.XML(res.Code, res)
		break
	default:
		ginCtx.JSON(res.Code, res)
	}
}
