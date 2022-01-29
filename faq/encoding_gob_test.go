package faq

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"
	"unsafe"
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

// TestEncodingGob_Sizeof gob编码后字节大小大2.5倍左右,何解?
// 横向对比其他序列化方式大小,如:json/xml等
func TestEncodingGob_Sizeof(t *testing.T) {
	user := &User{
		Name:     "李四",
		Birthday: time.Now(),
		Height:   178,
	}

	userSizeof := unsafe.Sizeof(user.Name) + unsafe.Sizeof(user.Birthday) + unsafe.Sizeof(user.Height)
	t.Logf("user sizeof:%v", userSizeof) // output: 48

	// gob encode
	bs, err := Encode(user)
	if err != nil {
		t.Fatalf("encode error:%v", err)
	}
	t.Logf("gob encode bytes length:%d", len(bs)) // output:103

	// json encode
	bs, err = json.Marshal(user)
	if err != nil {
		t.Fatalf("json encode error:%v", err)
	}
	t.Logf("json encode bytes length:%d", len(bs)) // output:76

	// xml encode
	bs, err = xml.Marshal(user)
	if err != nil {
		t.Fatalf("xml encode error:%v", err)
	}
	t.Logf("xml encode bytes length:%d", len(bs)) // output:105
}
