package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/auditmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/authmodel"
)

const (
	chars               = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	forgeResponsePrefix = "forge_"
)

// GenerateAccessKey
// Generates a new API access key and returns the prefix, plain text key and computed hash.
func GenerateAccessKey() (string, string, string, error) {
	prefix, err := RandomString(7)
	if err != nil {
		return "", "", "", err
	}
	key, err := RandomString(44)
	if err != nil {
		return "", "", "", err
	}
	fullKey := forgeResponsePrefix + prefix + "." + key
	return forgeResponsePrefix + prefix, fullKey, GenerateHash(fullKey), nil
}

// GenerateHash
// Creates a SHA256 hash of the provided string.
func GenerateHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}

// ValidateAvailableScopes
// Validates the supplied scopes against the scopes available in the application.
func ValidateAvailableScopes(providedScope string) bool {
	for _, scope := range authmodel.Scopes {
		if scope == providedScope {
			return true
		}
	}
	return false
}

// ValidateEventScopes
// Validates the supplied event against the events available in the application.
func ValidateEventScopes(providedScope string) bool {
	for _, scope := range auditmodel.AuditScopes {
		if scope == providedScope {
			return true
		}
	}
	return false
}

// RandomString
// Generates a cryptographically secure string of the length provided.
func RandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}
	return string(bytes), nil
}
