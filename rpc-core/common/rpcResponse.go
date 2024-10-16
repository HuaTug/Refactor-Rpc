package common

type RpcResponse struct {
	Data    interface{}
	Message string
}

func (r *RpcResponse) GetType() string {
    return "RpcRequest"
}