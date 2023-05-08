package mocks

import "github.com/ForgeResponse/ForgeResponse/v2/pkg/model/tenantmodel"

func Tenant() *tenantmodel.Tenant {
	return &tenantmodel.Tenant{
		TenantId: "ef47cb43-6a8e-47af-bcab-074b17b1bec0",
		Name:     "Test",
	}
}

func Tenant2() *tenantmodel.Tenant {
	return &tenantmodel.Tenant{
		TenantId: "5ee389c2-c013-42a1-bf25-f02fbc8fcbb1",
		Name:     "Test2",
	}
}

func Tenants() *[]tenantmodel.Tenant {
	return &[]tenantmodel.Tenant{
		{
			TenantId: "ef47cb43-6a8e-47af-bcab-074b17b1bec0",
			Name:     "Test",
		},
		{
			TenantId: "5ee389c2-c013-42a1-bf25-f02fbc8fcbb1",
			Name:     "Test2",
		},
	}
}
