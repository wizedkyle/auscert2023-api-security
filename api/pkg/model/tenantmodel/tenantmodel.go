package tenantmodel

type Tenant struct {
	TenantId string `bson:"tenantId" json:"tenantId"`
	Name     string `bson:"name" json:"name"`
}

type CreateTenantRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTenantRequest struct {
	Name string `json:"name" binding:"required"`
}
