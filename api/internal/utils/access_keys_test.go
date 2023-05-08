package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testKey        = "4a53bc6b8419"
	testSHA256Hash = "8c159ddc47b109a1807318a6a7b23151d80a4c2e803542804f1698cdf9ade7d2"
	validScope     = "read:users"
	invalidScope   = "read:abcdefg"
)

func TestGenerateHash(t *testing.T) {
	hash := GenerateHash(testKey)
	assert.Equal(t, testSHA256Hash, hash)
}

func TestValidateAvailableScopesSuccess(t *testing.T) {
	result := ValidateAvailableScopes(validScope)
	assert.True(t, result)
}

func TestValidateAvailableScopesFailure(t *testing.T) {
	result := ValidateAvailableScopes(invalidScope)
	assert.False(t, result)
}
