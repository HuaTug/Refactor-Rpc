package codec

import (
	"encoding/binary"
	"errors"
	"io"
	"log"

	"HuaTug.com/rpc-core/common"
	"HuaTug.com/rpc-core/protocol"
	"HuaTug.com/rpc-core/serialization"
)

type RpcDecoder struct{}

type Type interface {
	GetType() string
}

// ToDo:使用一个 Type 接口可以完成不同类型的处理，这是因为接口提供了一种抽象机制，使得不同的具体类型可以被统一处理。
func (d *RpcDecoder) Decode(reader io.Reader) (*protocol.MessageProtocol[Type], error) {
	var header protocol.MessageHeader

	// 读取魔数
	if err := binary.Read(reader, binary.BigEndian, &header.Magic); err != nil {
		return nil, err
	}
	if header.Magic != protocol.ProtocolConstants.MAGIC {
		return nil, errors.New("magic number is illegal")
	}

	// 读取其他字段
	if err := binary.Read(reader, binary.BigEndian, &header.Version); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &header.Serialization); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &header.MsgType); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &header.Status); err != nil {
		return nil, err
	}

	// 读取消息 ID（假设长度为8字节）
	requestIDBytes := make([]byte, 8)
	if _, err := reader.Read(requestIDBytes); err != nil {
		return nil, err
	}
	header.RequestID = string(requestIDBytes)

	// 读取数据长度
	var dataLength int32
	if err := binary.Read(reader, binary.BigEndian, &dataLength); err != nil {
		return nil, err
	}

	data := make([]byte, dataLength)
	if _, err := reader.Read(data); err != nil {
		return nil, err
	}
	msgTypeEnum := protocol.MsgType(header.MsgType)
	var body Type
	jsonSer, err := serialization.SerializationFactory(serialization.JSON)
	if err != nil {
		return nil, err
	}

	switch msgTypeEnum {
	case protocol.REQUEST:
		var request common.RpcRequest
		if err := jsonSer.Deserialize(data, &request); err != nil {
			return nil, err
		}
		body = &request
	case protocol.RESPONSE:
		var response common.RpcResponse
		if err := jsonSer.Deserialize(data, &response); err != nil {
			return nil, err
		}
		body = &response
	default:
		return nil, errors.New("unknown message type")
	}
	log.Println(body)

	// 检查 body 是否为 Type 类型
	if body, ok := body.(Type); ok {
		return &protocol.MessageProtocol[Type]{
			Header: &header,
			Body:   body,
		}, nil
	}
	return nil, errors.New("failed to decode message")
}
