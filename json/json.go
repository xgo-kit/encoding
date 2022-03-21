package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/xgo-kit/encoding"
)

const (
	Name = `json`
)

var (
	api = jsoniter.Config{
		EscapeHTML: true,
	}.Froze()
)

type jsonCodec struct {
}

func (j jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return api.Marshal(v)
}

func (j jsonCodec) Unmarshal(data []byte, v interface{}) error {
	return api.Unmarshal(data, v)
}

func (j jsonCodec) Name() string {
	return Name
}

func init() {
	encoding.Register(jsonCodec{})
}
