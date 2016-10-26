package client

import "encoding/json"

type EntityMapper interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type JsonMapper struct {

}

func(jm JsonMapper) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func(jm JsonMapper) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

