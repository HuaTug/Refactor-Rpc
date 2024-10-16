package serialization

import "fmt"

// SerializationException 是自定义的序列化异常类型
type SerializationException struct {
	Message string
	Err     error
}

func (e *SerializationException) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewSerializationException 创建新的 SerializationException
func NewSerializationException(message string, err error) *SerializationException {
	return &SerializationException{
		Message: message,
		Err:     err,
	}
}
