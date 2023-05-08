package mocks

import (
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/accesskeymodel"
	"time"
)

func AccessKey() *accesskeymodel.AccessKey {
	return &accesskeymodel.AccessKey{
		Id:          "f9daf268-0101-4dd8-945a-6c1a84ec6d0d",
		Description: "Mock Key",
		Expiration:  time.Time{},
		KeyHash:     "c4835bad5904fa435da1838ddc2cff91cc5bb6d388ad0c2b19175038da3f3ef0",
		KeyPrefix:   "SQFo1b3",
		TenantId:    "",
		Scopes: []string{
			"read:users",
			"read:tenants",
		},
		Version: 1,
	}
}

func AccessKey2() *accesskeymodel.AccessKey {
	return &accesskeymodel.AccessKey{
		Id:          "3b541d38-ea64-4131-93cf-aedc1e534cec",
		Description: "Mock Key",
		Expiration:  time.Time{},
		KeyHash:     "4e1331f6a03b40e152ff71ba6aa36f5ead92a045cfa8048930f4294f639a8755",
		KeyPrefix:   "2dXWhK9",
		TenantId:    "",
		Scopes: []string{
			"read:users",
			"read:tenants",
		},
		Version: 1,
	}
}

func AccessKeyResponse() *accesskeymodel.AccessKeyResponse {
	return &accesskeymodel.AccessKeyResponse{
		Id:          "f9daf268-0101-4dd8-945a-6c1a84ec6d0d",
		Description: "Mock Key",
		Expiration:  time.Time{},
		KeyPrefix:   "SQFo1b3",
		Scopes: []string{
			"read:users",
			"read:tenants",
		},
	}
}

func AccessKeys() *[]accesskeymodel.AccessKeyResponse {
	return &[]accesskeymodel.AccessKeyResponse{
		{
			Id:          "f9daf268-0101-4dd8-945a-6c1a84ec6d0d",
			Description: "Mock Key",
			Expiration:  time.Time{},
			KeyPrefix:   "SQFo1b3",
			Scopes: []string{
				"read:users",
				"read:tenants",
			},
		},
		{
			Id:          "3b541d38-ea64-4131-93cf-aedc1e534cec",
			Description: "Mock Key",
			Expiration:  time.Time{},
			KeyPrefix:   "2dXWhK9",
			Scopes: []string{
				"read:users",
				"read:tenants",
			},
		},
	}
}

func UpdateAccessKeyResponse() *accesskeymodel.UpdateAccessKeyResponse {
	return &accesskeymodel.UpdateAccessKeyResponse{
		Id:          "f9daf268-0101-4dd8-945a-6c1a84ec6d0d",
		Description: "Mock Key",
		Expiration:  time.Time{},
		Scopes: []string{
			"read:users",
			"read:tenants",
		},
	}
}
