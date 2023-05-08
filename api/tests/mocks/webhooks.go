package mocks

import (
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/auditmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/webhookmodel"
)

func Webhook() *webhookmodel.Webhook {
	return &webhookmodel.Webhook{
		Id:          "6a9fd917-7829-492e-b540-f0b835a8f74b",
		Algorithm:   "SHA256",
		Description: "Test Webhook",
		Events: []string{
			auditmodel.AccessKeyCreated,
			auditmodel.AccessKeyUpdated,
		},
		Secret:   "houehfjnjhouiiepwfsdkcmisdfksdmfksd",
		TenantId: "",
		Url:      "https://test.com/webhook",
		Version:  1,
	}
}

func Webhook2() *webhookmodel.Webhook {
	return &webhookmodel.Webhook{
		Id:          "485e0335-7fee-4909-9b58-4247deb54ffd",
		Algorithm:   "SHA256",
		Description: "Test Webhook",
		Events: []string{
			auditmodel.AccessKeyCreated,
			auditmodel.AccessKeyUpdated,
		},
		Secret:   "fsuhfiuafdanfljdjfasdnfkljadfakdfna",
		TenantId: "",
		Url:      "https://test.com/webhook",
		Version:  1,
	}
}

func WebhookResponse() *webhookmodel.WebhookResponse {
	return &webhookmodel.WebhookResponse{
		Id:          "6a9fd917-7829-492e-b540-f0b835a8f74b",
		Algorithm:   "SHA256",
		Description: "Test Webhook",
		Events: []string{
			auditmodel.AccessKeyCreated,
			auditmodel.AccessKeyUpdated,
		},
		Secret: "houehfjnjhouiiepwfsdkcmisdfksdmfksd",
		Url:    "https://test.com/webhook",
	}
}

func Webhooks() *[]webhookmodel.WebhookResponse {
	return &[]webhookmodel.WebhookResponse{
		{
			Id:          "6a9fd917-7829-492e-b540-f0b835a8f74b",
			Algorithm:   "SHA256",
			Description: "Test Webhook",
			Events: []string{
				auditmodel.AccessKeyCreated,
				auditmodel.AccessKeyUpdated,
			},
			Secret: "houehfjnjhouiiepwfsdkcmisdfksdmfksd",
			Url:    "https://test.com/webhook",
		},
		{
			Id:          "485e0335-7fee-4909-9b58-4247deb54ffd",
			Algorithm:   "SHA256",
			Description: "Test Webhook",
			Events: []string{
				auditmodel.AccessKeyCreated,
				auditmodel.AccessKeyUpdated,
			},
			Secret: "fsuhfiuafdanfljdjfasdnfkljadfakdfna",
			Url:    "https://test.com/webhook",
		},
	}
}
