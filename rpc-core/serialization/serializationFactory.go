package serialization

import (
	"errors"
)


// SerializationFactory 根据类型返回对应的序列化实现
func SerializationFactory(typeEnum SerializationTypeEnum) (RpcSerialization, error) {
	switch typeEnum {
	case JSON:
		return &JsonSerialization{}, nil
	default:
		return nil, errors.New("unsupported serialization type")
	}
}
