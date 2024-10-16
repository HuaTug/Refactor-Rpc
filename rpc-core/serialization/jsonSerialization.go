package serialization

import (
	"encoding/json"
)

// JsonSerialization 实现 JSON 序列化
type JsonSerialization struct{}

func (j *JsonSerialization) Serialize(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j *JsonSerialization) Deserialize(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
