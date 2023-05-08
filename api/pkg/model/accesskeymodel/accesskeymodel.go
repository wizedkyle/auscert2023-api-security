package accesskeymodel

import "time"

type AccessKey struct {
	Id          string    `bson:"id" json:"id"`
	Description string    `bson:"description" json:"description"`
	Expiration  time.Time `bson:"expiration" json:"expiration"`
	KeyHash     string    `bson:"keyHash" json:"keyHash"`
	KeyPrefix   string    `bson:"keyPrefix" json:"keyPrefix"`
	TenantId    string    `bson:"tenantId" json:"tenantId"`
	Scopes      []string  `bson:"scopes" json:"scopes"`
	Version     int       `bson:"version" json:"version"`
}

type CreateAccessKeyRequest struct {
	Description string   `json:"description" binding:"required"`
	Duration    int      `json:"duration"`
	Scopes      []string `json:"scopes" binding:"required"`
}

type CreateAccessKeyResponse struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Expiration  time.Time `json:"expiration"`
	Key         string    `json:"key"`
	Scopes      []string  `json:"scopes"`
}

type AccessKeyResponse struct {
	Id          string    `bson:"id" json:"id"`
	Description string    `bson:"description" json:"description"`
	Expiration  time.Time `bson:"expiration" json:"expiration"`
	KeyPrefix   string    `bson:"keyPrefix" json:"keyPrefix"`
	Scopes      []string  `bson:"scopes" json:"scopes"`
}

type RotateAccessKeyRequest struct {
	Duration int `json:"duration" binding:"required"`
}

type RotateAccessKeyResponse struct {
	Id         string    `json:"id"`
	Expiration time.Time `json:"expiration"`
	Key        string    `json:"key"`
}

type UpdateAccessKeyRequest struct {
	Description string   `json:"description" binding:"required"`
	Scopes      []string `json:"scopes" binding:"required"`
}

type UpdateAccessKeyResponse struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Expiration  time.Time `json:"expiration"`
	Scopes      []string  `json:"scopes"`
}
