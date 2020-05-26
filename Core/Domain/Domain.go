package Domain

type Message struct {
	Code int
	Msg  string
	Data map[string]interface{}
}

func NewMessage(c int, m string, d map[string]interface{}) Message {

	return Message{
		Code: c,
		Msg:  m,
		Data: d,
	}
}

type RpcRequest struct {
	Request map[string]interface{} `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func NewRpcRequest() RpcRequest {
	reqMap := make(map[string]interface{})

	return RpcRequest{
		Request: reqMap,
	}
}

type RpcResponse struct {
	Response map[string]interface{} `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func NewRpcResponse() RpcResponse {
	respMap := make(map[string]interface{})

	return RpcResponse{
		Response: respMap,
	}
}
