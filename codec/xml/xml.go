package xml

import (
	"encoding/xml"

	"github.com/taadis/letgo/codec"
)

func init() {
	codec.Register(xmlCodec{})
}

type xmlCodec struct{}

func (xmlCodec) Name() string {
	return "xml"
}

func (xmlCodec) Marshal(v any) ([]byte, error) {
	return xml.Marshal(v)
}

func (xmlCodec) Unmarshal(data []byte, v any) error {
	return xml.Unmarshal(data, v)
}
