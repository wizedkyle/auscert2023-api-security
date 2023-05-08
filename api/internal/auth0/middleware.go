// A lot of this code has been taken from https://github.com/auth0/go-jwt-middleware and refactored to work natively with
// the Go Gin framework.

package auth0

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ContextKey is the key used in the request
// context where the information from a
// validated JWT will be stored.
type ContextKey struct{}

type JWTMiddleware struct {
	validateToken       ValidateToken
	tokenExtractor      TokenExtractor
	credentialsOptional bool
	validateOnOptions   bool
}

// ValidateToken takes in a string JWT and makes sure it is valid and
// returns the valid token. If it is not valid it will return nil and
// an error message describing why validation failed.
// Inside ValidateToken things like key and alg checking can happen.
// In the default implementation we can add safe defaults for those.
type ValidateToken func(context.Context, string) (interface{}, error)

// NewMiddleware constructs a new JWTMiddleware instance with the supplied options.
// It requires a ValidateToken function to be passed in, so it can
// properly validate tokens.
func NewMiddleware(validateToken ValidateToken) *JWTMiddleware {
	m := &JWTMiddleware{
		validateToken:       validateToken,
		credentialsOptional: false,
		tokenExtractor:      AuthHeaderTokenExtractor,
		validateOnOptions:   true,
	}
	return m
}

// CheckJWT is the main JWTMiddleware function which performs the main logic. It
// is passed a http.Handler which will be called if the JWT passes validation.
func (m *JWTMiddleware) CheckJWT(c *gin.Context) {
	// If we don't validate on OPTIONS and this is OPTIONS
	// then continue onto next without validating.
	if !m.validateOnOptions && c.Request.Method == http.MethodOptions {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	token, err := m.tokenExtractor(c.Request)
	if err != nil {
		// This is not ErrJWTMissing because an error here means that the
		// tokenExtractor had an error and _not_ that the token was missing.
		c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("error extracting token: %w", err))
		return
	}

	if token == "" {
		// If credentials are optional continue
		// onto next without validating.
		if m.credentialsOptional {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "unauthorized")
			return
		}
	}

	// Validate the token using the token validator.
	validToken, err := m.validateToken(c.Request.Context(), token)
	if err != nil {
		// TODO: Add error handling
		c.AbortWithStatusJSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	// No err means we have a valid token, so set
	// it into the context and continue onto next.
	c.Set("token", validToken)
	c.Next()
}
