package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/taadis/letgo/codec"
)

func init() {
	codec.Register(gobCodec{})
}

type gobCodec struct{}

func (gobCodec) Name() string {
	return "gob"
}

func (gobCodec) Marshal(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (gobCodec) Unmarshal(data []byte, v any) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(v)
}
