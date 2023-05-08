package middleware

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/auth0"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/database"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type CustomClaims struct {
	Permissions []string `json:"permissions"`
	TenantId    string   `json:"https://www.forgeresponse.com/tenant_id"`
}

const secretKey = "abc1234"

// ValidateAuth
// Takes the authorization header and determines if it is a bearer token or API token and applies the appropriate
// authentication mechanism. For API key authentication the associated scopes in the database are set in the gin context
// for use in authorization to reduce a database query.
func ValidateAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, "no authorization header provided")
			return
		}
		token := strings.Split(auth, " ")
		if token[0] == "Bearer" {
			// Performs auth based on JWT
			audience := os.Getenv("audience")
			issuerUrl, err := url.Parse(os.Getenv("issuerUrl"))
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}
			provider := auth0.NewCachingProvider(issuerUrl, time.Duration(5*time.Minute))
			jwtValidator, err := auth0.New(provider.KeyFunc,
				auth0.RS256,
				issuerUrl.String(),
				[]string{audience},
				auth0.WithCustomClaims(
					func() auth0.CustomClaims {
						return &CustomClaims{}
					},
				),
				auth0.WithAllowedClockSkew(time.Minute),
			)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}
			middleware := auth0.NewMiddleware(jwtValidator.ValidateToken)
			middleware.CheckJWT(c)
		} else {
			// Performs auth based on an API key
			hash := utils.GenerateHash(auth)
			accessKey, err := database.Client.GetAccessKeyByHash(hash)
			if err == utils.ErrAccessKeyNotFound {
				c.AbortWithStatusJSON(http.StatusForbidden, "invalid api key")
				return
			} else if err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}
			if hash != accessKey.KeyHash {
				c.AbortWithStatusJSON(http.StatusForbidden, "invalid api key")
				return
			}
			if time.Now().After(accessKey.Expiration) {
				c.AbortWithStatusJSON(http.StatusForbidden, "api key is expired")
				return
			}
			c.Set("scopes", accessKey.Scopes)
			c.Set("tenantId", accessKey.TenantId)
		}
		c.Next()
	}
}

func ValidateQueryStringAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, exists := c.GetQuery("api_key")
		if auth == "" && !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, "no authorization key provided")
			return
		}
		hash := utils.GenerateHash(auth)
		accessKey, err := database.Client.GetAccessKeyByHash(hash)
		if err == utils.ErrAccessKeyNotFound {
			c.AbortWithStatusJSON(http.StatusForbidden, "invalid api key")
			return
		} else if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		if hash != accessKey.KeyHash {
			c.AbortWithStatusJSON(http.StatusForbidden, "invalid api key")
			return
		}
		if time.Now().After(accessKey.Expiration) {
			c.AbortWithStatusJSON(http.StatusForbidden, "api key is expired")
			return
		}
		c.Set("scopes", accessKey.Scopes)
		c.Set("tenantId", accessKey.TenantId)
		c.Next()
	}
}

func ValidateHmacAuthWithTimestampValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		hash := c.GetHeader("HMAC-Hash")
		timestamp := c.GetHeader("HMAC-Timestamp")
		if hash == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, "no HMAC-Hash header value provided")
			return
		}
		if timestamp == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, "no HMAC-Timestamp header value provided")
			return
		}
		i, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, "no timestamp provided")
			return
		}
		requestTime := time.Unix(i, 0)
		maxTime := time.Now().Add(time.Minute * 5)
		if !requestTime.Before(maxTime) {
			c.AbortWithStatusJSON(http.StatusForbidden, "request timestamp exceeds 5 minute limit")
			return
		}
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "internal server error")
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		hashDecoded, err := hex.DecodeString(hash)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "internal server error")
			return
		}
		message := timestamp + ":" + string(bodyBytes)
		mac := hmac.New(sha256.New, []byte(secretKey))
		mac.Write([]byte(message))
		computedMac := mac.Sum(nil)
		if !hmac.Equal(hashDecoded, computedMac) {
			c.AbortWithStatusJSON(http.StatusForbidden, "message hashes do not match")
			return
		}
		c.Next()
	}
}

func ValidateHmacAuthWithoutTimestampValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		hash := c.GetHeader("HMAC-Hash")
		if hash == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, "no HMAC-Hash header value provided")
			return
		}
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "internal server error")
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		mac := hmac.New(sha256.New, []byte(secretKey))
		mac.Write(bodyBytes)
		computedMac := mac.Sum(nil)
		if !hmac.Equal([]byte(hash), computedMac) {
			c.AbortWithStatusJSON(http.StatusForbidden, "message hashes do not match")
			return
		}
		c.Next()
	}
}

// ValidateScopes
// Takes the scopes in an OAuth token or retrieves them from the database for API authentication and verifies if the
// actor requesting access to a resource has the required permissions.
func ValidateScopes(expectedScope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		scopes, _ := c.Get("scopes")
		if scopes == nil {
			authToken, _ := c.Get("token")
			token := authToken.(*auth0.ValidatedClaims)
			claims := token.CustomClaims.(*CustomClaims)
			c.Set("tenantId", claims.TenantId)
			if !hasScope(expectedScope, claims.Permissions) {
				c.AbortWithStatusJSON(http.StatusForbidden, "insufficient scope")
				return
			}
		} else {
			var constructedScopes []string
			switch scopesArray := scopes.(type) {
			case []string:
				for _, value := range scopesArray {
					constructedScopes = append(constructedScopes, value)
				}
			}
			if !hasScope(expectedScope, constructedScopes) {
				c.AbortWithStatusJSON(http.StatusForbidden, "insufficient scope")
				return
			}
		}
		c.Next()
	}
}

// hasScope
// Takes the expected scope and an interface containing scopes and checks if the expectedScope is in the provided scopes.
func hasScope(expectedScope string, scopes interface{}) bool {
	switch scope := scopes.(type) {
	case []string:
		for i := range scope {
			if scope[i] == expectedScope {
				return true
			}
		}
	case string:
		result := strings.Split(scope, " ")
		for i := range result {
			if result[i] == expectedScope {
				return true
			}
		}
	}
	return false
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}
