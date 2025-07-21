package codec

import (
	"encoding/xml"
	"runtime/debug"
	"testing"
)

type codec struct{}

func (c codec) Marshal(_ any) ([]byte, error) {
	panic("implement me")
}

func (c codec) Unmarshal(_ []byte, _ any) error {
	panic("implement me")
}

func (c codec) Name() string {
	return ""
}

// codec2 is a Codec implementation with xml.
type codec2 struct{}

func (codec2) Marshal(v any) ([]byte, error) {
	return xml.Marshal(v)
}

func (codec2) Unmarshal(data []byte, v any) error {
	return xml.Unmarshal(data, v)
}

func (codec2) Name() string {
	return "xml"
}

func TestRegisterCodec(t *testing.T) {
	f := func() { Register(nil) }
	funcDidPanic, panicValue, _ := didPanic(f)
	if !funcDidPanic {
		t.Fatalf("func should panic\n\tPanic value:\t%#v", panicValue)
	}
	if panicValue != "cannot register a nil Codec" {
		t.Fatalf("panic error got %s want cannot register a nil Codec", panicValue)
	}
	f = func() {
		Register(codec{})
	}
	funcDidPanic, panicValue, _ = didPanic(f)
	if !funcDidPanic {
		t.Fatalf("func should panic\n\tPanic value:\t%#v", panicValue)
	}
	if panicValue != "cannot register Codec with empty Name()" {
		t.Fatalf("panic error got %s want cannot register Codec with empty Name()", panicValue)
	}
	codec := codec2{}
	Register(codec)
	got := GetCodec("xml")
	if got != codec {
		t.Fatalf("RegisterCodec(%v) want %v got %v", codec, codec, got)
	}
}

// PanicTestFunc defines a func that should be passed to assert.Panics and assert.NotPanics
// methods, and represents a simple func that takes no arguments, and returns nothing.
type PanicTestFunc func()

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
func didPanic(f PanicTestFunc) (bool, any, string) {
	didPanic := false
	var message any
	var stack string
	func() {
		defer func() {
			if message = recover(); message != nil {
				didPanic = true
				stack = string(debug.Stack())
			}
		}()

		// call the target function
		f()
	}()

	return didPanic, message, stack
}
