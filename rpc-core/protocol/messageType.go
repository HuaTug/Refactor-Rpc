package protocol

// MsgType 消息类型
type MsgType byte

const (
	REQUEST  MsgType = 1
	RESPONSE MsgType = 2
)

// FindByType 根据类型查找消息类型
func FindByType(t byte) MsgType {
	switch t {
	case byte(REQUEST):
		return REQUEST
	case byte(RESPONSE):
		return RESPONSE
	default:
		return 0
	}
}
