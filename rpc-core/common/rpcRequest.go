package common

type RpcRequest struct {
	ServiceMethod  string
	Method         string
	ParameterTypes []string
	Parameters     []interface{}
}

func (r *RpcRequest) GetType() string {
    return "RpcRequest"
}