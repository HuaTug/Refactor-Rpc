package serialization

import (
	"errors"
	"strings"
)

type SerializationTypeEnum int

const (
	JSON SerializationTypeEnum = 1
)

func (s SerializationTypeEnum) GetType() byte {
	return byte(s)
}

func ParseByName(typeName string) (SerializationTypeEnum, error) {
	switch strings.ToUpper(typeName) {
	case "JSON":
		return JSON, nil
	default:
		return JSON, errors.New("unsupported serialization type")
	}
}

func ParseByType(typeByte byte) (SerializationTypeEnum, error) {
	switch SerializationTypeEnum(typeByte) {
	case JSON:
		return JSON, nil
	default:
		return JSON, errors.New("unsupported serialization type")
	}
}
