package usermodel

import (
	"time"
)

type User struct {
	Id           string      `bson:"id" json:"id"`
	TenantId     string      `bson:"tenantId" json:"tenantId"`
	Email        string      `bson:"email" json:"email"`
	FirstName    string      `bson:"firstName"json:"firstName"`
	LastName     string      `bson:"lastName" json:"lastName"`
	LastSignIn   time.Time   `bson:"lastSignIn" json:"lastSignIn"`
	CreatedTime  time.Time   `bson:"createdTime" json:"createdTime"`
	Roles        []UserRoles `bson:"roles" json:"roles"`
	IsActive     bool        `bson:"isActive" json:"isActive"`
	AccountOwner bool        `bson:"accountOwner" json:"accountOwner"`
}

type UserRoles struct {
	Id string `json:"id,omitempty"`
}

type UserResponse struct {
	Id          string      `bson:"id" json:"id"`
	Email       string      `bson:"email" json:"email"`
	FirstName   string      `bson:"firstName"json:"firstName"`
	LastName    string      `bson:"lastName" json:"lastName"`
	LastSignIn  time.Time   `bson:"lastSignIn" json:"lastSignIn"`
	CreatedTime time.Time   `bson:"createdTime" json:"createdTime"`
	Roles       []UserRoles `bson:"roles" json:"roles"`
	IsActive    bool        `bson:"isActive" json:"isActive"`
}

type CreateUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"firstName" binding:"required,alpha"`
	LastName  string `json:"lastName" binding:"required,alpha"`
}

type UpdateUserRequest struct {
	FirstName string `json:"firstName" binding:"required,alpha"`
	LastName  string `json:"lastName" binding:"required,alpha"`
	IsActive  *bool  `json:"isActive" binding:"required"`
}
