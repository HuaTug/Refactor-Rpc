package serialization

// RpcSerialization 定义序列化接口
type RpcSerialization interface {
	Serialize(obj interface{}) ([]byte, error)
	Deserialize(data []byte, out interface{}) error
}
