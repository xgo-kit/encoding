package encoding

var (
	registered = make(map[string]Codec)
)

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Name() string
}

func Register(codec Codec) {
	registered[codec.Name()] = codec
}

func Get(name string) Codec {
	return registered[name]
}
