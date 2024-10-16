package main

import (
	"bytes"
	"fmt"
	"log"

	"HuaTug.com/rpc-core/codec"
	"HuaTug.com/rpc-core/common"
	"HuaTug.com/rpc-core/protocol"
)

func test() {
	request := common.RpcRequest{
		ServiceMethod:  "ExampleService.Method",
		Method:         "MethodName",
		ParameterTypes: []string{"string"},
		Parameters:     []interface{}{"example"},
	}

	header := protocol.MessageHeader{
		Magic:         protocol.ProtocolConstants.MAGIC,
		Version:       1,
		Serialization: 1,
		MsgType:       byte(protocol.REQUEST),
		Status:        0,
		RequestID:     "12345678",
	}

	message := &protocol.MessageProtocol[common.RpcRequest]{
		Header: &header,
		Body:   request,
	}

	encoder := codec.RpcEncoder{
		Writer: new(bytes.Buffer),
	}
	if err := encoder.WriteBody(&header, request); err != nil {
		log.Fatalf("Failed to encode message: %v", err)
	}

	encodedData := encoder.Writer.Bytes()

	log.Printf("Encoded Data: %v", encodedData)

	decoder := codec.RpcDecoder{}

	decodedMessage, err := decoder.Decode(bytes.NewBuffer(encodedData))
	if err != nil {
		log.Fatalf("Failed to decode message: %v", err)
	}

	if decodedMessage.Header.RequestID != message.Header.RequestID {
		log.Fatal("Request ID does not match")
	}
	if req, ok := decodedMessage.Body.(*common.RpcRequest); ok {
		if req.ServiceMethod != request.ServiceMethod {
			log.Fatalf("Expected ServiceMethod %s, got %s", request.ServiceMethod, req.ServiceMethod)
		}
		fmt.Println("Encoding and decoding successful!")
		fmt.Printf("Decoded Request: %+v\n", req)
	} else {
		log.Fatal("Decoded body is not of type RpcRequest")
	}
}
