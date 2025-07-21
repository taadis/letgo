package yaml

import (
	"gopkg.in/yaml.v3"

	"github.com/taadis/letgo/codec"
)

func init() {
	codec.Register(&yamlCodec{})
}

// yamlCodec is a Codec implementation with yaml.
type yamlCodec struct{}

func (*yamlCodec) Marshal(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

func (*yamlCodec) Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}

func (*yamlCodec) Name() string {
	return "yaml"
}
