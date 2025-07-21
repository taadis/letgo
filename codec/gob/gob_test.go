package gob

import (
	"reflect"
	"testing"

	"github.com/taadis/letgo/codec"
)

type testStruct struct {
	Name  string
	Value int
}

func TestGobCodec_Name(t *testing.T) {
	c := gobCodec{}
	if name := c.Name(); name != "gob" {
		t.Errorf("expected Name() to return 'gob', got '%s'", name)
	}
}

func TestGobCodec_MarshalUnmarshal(t *testing.T) {
	// 注册 codec
	codec.Register(gobCodec{})

	// 获取 codec
	c := codec.GetCodec("gob")
	if c == nil {
		t.Fatal("GetCodec('gob') returned nil")
	}

	// 测试用例
	tests := []struct {
		name  string
		input testStruct
	}{
		{
			name:  "basic struct",
			input: testStruct{Name: "test", Value: 42},
		},
		{
			name:  "empty string",
			input: testStruct{Name: "", Value: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试 Marshal
			data, err := c.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}
			if len(data) == 0 {
				t.Fatal("Marshal returned empty data")
			}

			// 测试 Unmarshal
			var result testStruct
			err = c.Unmarshal(data, &result)
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			// 验证结果
			if !reflect.DeepEqual(tt.input, result) {
				t.Errorf("Unmarshaled value does not match input: got %+v, want %+v", result, tt.input)
			}
		})
	}
}

func TestGobCodec_UnmarshalInvalidData(t *testing.T) {
	c := gobCodec{}
	var result testStruct
	err := c.Unmarshal([]byte("invalid gob data"), &result)
	if err == nil {
		t.Error("Unmarshal should fail on invalid data")
	}
}
