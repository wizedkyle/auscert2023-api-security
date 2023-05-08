package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateExpirationNeverExpire(t *testing.T) {
	expiration := GenerateExpiration(0)
	assert.Equal(t, 2099, expiration.Year())
}

func TestGenerateExpiration(t *testing.T) {
	expiration := GenerateExpiration(30)
	expectedTime := time.Now().AddDate(0, 0, 30)
	assert.Equal(t, expectedTime.Year(), expiration.Year())
	assert.Equal(t, expectedTime.Month(), expiration.Month())
	assert.Equal(t, expectedTime.Day(), expiration.Day())
}
