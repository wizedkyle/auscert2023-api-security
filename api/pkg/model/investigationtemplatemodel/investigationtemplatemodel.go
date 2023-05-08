package investigationtemplatemodel

import "time"

type InvestigationTemplate struct {
	Id          string    `bson:"id" json:"id"`
	TenantId    string    `bson:"tenantId" json:"tenantId"`
	CreatedBy   string    `bson:"createdBy" json:"createdBy"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	Description string    `bson:"description" json:"description"`
	TitlePrefix string    `bson:"titlePrefix" json:"titlePrefix"`
	Severity    string    `bson:"severity" json:"severity"`
	Status      string    `bson:"status" json:"status"`
	Tags        []string  `bson:"tags" json:"tags"`
	Tlp         int       `bson:"tlp" json:"tlp"`
}

type InvestigationTemplateResponse struct {
	Id          string    `json:"id"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	TitlePrefix string    `json:"titlePrefix"`
	Severity    string    `json:"severity"`
	Status      string    `json:"status"`
	Tags        []string  `json:"tags"`
	Tlp         int       `json:"tlp"`
}

type CreateInvestigationRequest struct {
	Description string   `json:"description"`
	TitlePrefix string   `json:"titlePrefix" binding:"required"`
	Severity    string   `json:"severity" binding:"required"`
	Status      string   `json:"status" binding:"required"`
	Tags        []string `json:"tags" binding:"required"`
	Tlp         int      `json:"tlp" binding:"required"`
}

type UpdateInvestigationRequest struct {
	Description string   `json:"description"`
	TitlePrefix string   `json:"titlePrefix" binding:"required"`
	Severity    string   `json:"severity" binding:"required"`
	Status      string   `json:"status" binding:"required"`
	Tags        []string `json:"tags" binding:"required"`
	Tlp         int      `json:"tlp" binding:"required"`
}
