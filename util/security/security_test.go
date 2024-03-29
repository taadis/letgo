package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestWithMd5
func TestWithMd5(t *testing.T) {
	plaintext := "123456"
	ciphertext := WithMd5(plaintext)

	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", ciphertext)
}

// TestCheckPassword
func TestCheckPassword(t *testing.T) {
	result := CheckPassword("123", "456")
	assert.Equal(t, false, result)

	result = CheckPassword("123", "123")
	assert.Equal(t, true, result)
}
