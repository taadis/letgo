package yaml

import (
	"gopkg.in/yaml.v3"

	"github.com/taadis/letgo/codec"
)

func init() {
	codec.Register(&ymlCodec{})
}

// ymlCodec is a Codec implementation with yaml.
type ymlCodec struct{}

func (*ymlCodec) Marshal(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

func (*ymlCodec) Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}

func (*ymlCodec) Name() string {
	return "yml"
}
