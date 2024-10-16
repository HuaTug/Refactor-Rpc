package codec

import (
	"bytes"
	"encoding/binary"

	"HuaTug.com/rpc-core/protocol"
	"HuaTug.com/rpc-core/serialization"
)

type RpcEncoder struct {
	Writer *bytes.Buffer // 假设这里使用 bytes.Buffer 作为写入目标
}

func (e *RpcEncoder) WriteBody(header *protocol.MessageHeader, body interface{}) error {
	if body == nil {
		return nil
	}

	// 创建一个字节缓冲区来写入数据
	byteBuf := new(bytes.Buffer)

	// 写入魔数
	if err := binary.Write(byteBuf, binary.BigEndian, header.Magic); err != nil {
		return err
	}

	// 写入协议版本号、序列化算法、报文类型、状态
	byteBuf.WriteByte(header.Version)
	byteBuf.WriteByte(header.Serialization)
	byteBuf.WriteByte(header.MsgType)
	byteBuf.WriteByte(header.Status)

	// 写入消息 ID（假设长度为8字节）
	byteBuf.WriteString(header.RequestID)

	// 序列化消息体
	jsonSer, err := serialization.SerializationFactory(serialization.JSON)
	if err != nil {
		return err
	}

	data, err := jsonSer.Serialize(body)
	if err != nil {
		return err
	}

	// 数据长度
	if err := binary.Write(byteBuf, binary.BigEndian, int32(len(data))); err != nil {
		return err
	}

	// 写入数据内容
	byteBuf.Write(data)

	// 将最终的数据写入 writer（这里假设 writer 是一个 io.Writer）
	_, err = e.Writer.Write(byteBuf.Bytes())

	return err
}
