package yaml

import (
	"github.com/xgo-kit/encoding"
	yamlv3 "gopkg.in/yaml.v3"
)

const (
	Name = `yaml`
)

type yamlCodec struct {
}

func (y yamlCodec) Name() string {
	return Name
}

func (y yamlCodec) Marshal(v interface{}) ([]byte, error) {
	return yamlv3.Marshal(v)
}

func (y yamlCodec) Unmarshal(data []byte, v interface{}) error {
	return yamlv3.Unmarshal(data, v)
}

func init() {
	encoding.Register(yamlCodec{})
}
