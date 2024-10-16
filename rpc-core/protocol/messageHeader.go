package protocol

import (
	"fmt"
	"math/rand"
	"time"
)

// MessageHeader 消息头
type MessageHeader struct {
	Magic         uint16 // 魔数 (2byte)
	Version       byte   // 协议版本号 (1byte)
	Serialization byte   // 序列化算法 (1byte)
	MsgType       byte   // 报文类型 (1byte)
	Status        byte   // 状态 (1byte)
	RequestID     string // 消息 ID (32byte)
	MsgLen        int32  // 数据长度 (4byte)
}

// BuildMessageHeader 构建消息头
func BuildMessageHeader(serialization byte) *MessageHeader {
	return &MessageHeader{
		Magic:         ProtocolConstants.MAGIC,
		Version:       ProtocolConstants.VERSION,
		Serialization: serialization,
		MsgType:       byte(REQUEST),
		RequestID:     generateUUID(),
	}
}

// 生成唯一ID (伪UUID)
func generateUUID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%032x", rand.Uint64())
}
