package mocks

import (
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/usermodel"
	"time"
)

func StandardUser() *usermodel.User {
	return &usermodel.User{
		Id:          "2d0a5280-bf7c-4923-8157-e3629a050bdc",
		TenantId:    "",
		Email:       "test@test.com",
		FirstName:   "Test",
		LastName:    "Test",
		LastSignIn:  time.Time{},
		CreatedTime: time.Time{},
		Roles:       nil,
		IsActive:    true,
	}
}

func StandardUser2() *usermodel.User {
	return &usermodel.User{
		Id:          "687166ee-dfd7-49a0-8810-0d331ce83fe5",
		TenantId:    "",
		Email:       "test2@test.com",
		FirstName:   "Test2",
		LastName:    "Test2",
		LastSignIn:  time.Time{},
		CreatedTime: time.Time{},
		Roles:       nil,
		IsActive:    true,
	}
}

func UserResponse() *usermodel.UserResponse {
	return &usermodel.UserResponse{
		Id:          "2d0a5280-bf7c-4923-8157-e3629a050bdc",
		Email:       "test@test.com",
		FirstName:   "Test",
		LastName:    "Test",
		LastSignIn:  time.Time{},
		CreatedTime: time.Time{},
		Roles:       nil,
		IsActive:    true,
	}
}

func StandardUsers() *[]usermodel.UserResponse {
	return &[]usermodel.UserResponse{
		{
			Id:          "2d0a5280-bf7c-4923-8157-e3629a050bdc",
			Email:       "test@test.com",
			FirstName:   "Test",
			LastName:    "Test",
			LastSignIn:  time.Time{},
			CreatedTime: time.Time{},
			Roles:       nil,
			IsActive:    true,
		},
		{
			Id:          "687166ee-dfd7-49a0-8810-0d331ce83fe5",
			Email:       "test2@test.com",
			FirstName:   "Test2",
			LastName:    "Test2",
			LastSignIn:  time.Time{},
			CreatedTime: time.Time{},
			Roles:       nil,
			IsActive:    true,
		},
	}
}
