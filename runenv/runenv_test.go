package runenv

import "testing"

func TestIs(t *testing.T)  {
	isDevelopment := Is(Development)
	if !isDevelopment {
		t.Fail()
	}
}