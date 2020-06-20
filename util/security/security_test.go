package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithMd5(t *testing.T) {
	plaintext := "123456"
	ciphertext := WithMd5(plaintext)

	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", ciphertext)
}
