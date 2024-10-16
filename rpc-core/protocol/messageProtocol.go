package protocol

// MessageProtocol 消息协议
type MessageProtocol[T any] struct {
	Header *MessageHeader // 消息头
	Body   T    // 消息体
}
