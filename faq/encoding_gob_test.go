package faq

import (
	"bytes"
	"encoding/gob"
	"testing"
	"time"
)

type User struct {
	Name     string    // 姓名
	Birthday time.Time // 生日
	Height   int       // 身高(cm)
}

// Encode 用gob对数据结构进行编码
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decode 用gob把字节数组解码至指定的数据结构
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	return decoder.Decode(to)
}

func TestEncodingGob_EncodeDecode(t *testing.T) {
	user := &User{
		Name:     "张三",
		Birthday: time.Now(),
		Height:   177,
	}

	// 对user进行编码
	bs, err := Encode(user)
	if err != nil {
		t.Fatalf("encode error:%+v", err)
	}

	// 对user进行解码
	var out User
	err = Decode(bs, &out)
	if err != nil {
		t.Fatalf("decode error:%+v", err)
	}

	t.Logf("out user:%+v", out)
}
