package investigationmodel

import "time"

type Investigation struct {
	Id              string                     `bson:"id" json:"id"`
	AssignedTo      string                     `bson:"assignedTo" json:"assignedTo"`
	Attachments     []InvestigationAttachments `bson:"attachments" json:"attachments"`
	CreatedBy       string                     `bson:"createdBy" json:"createdBy"`
	CreatedAt       time.Time                  `bson:"createdAt" json:"createdAt"`
	Comments        []InvestigationComments    `bson:"comments" json:"comments"`
	Description     string                     `bson:"description" json:"description"`
	InvestigationId string                     `bson:"investigationId" json:"investigationId"`
	Severity        string                     `bson:"severity" json:"severity"`
	Status          string                     `bson:"status" json:"status"`
	Tags            []string                   `bson:"tags" json:"tags"`
	TenantId        string                     `bson:"tenantId" json:"tenantId"`
	Title           string                     `bson:"title" json:"title"`
	Tlp             int                        `bson:"tlp" json:"tlp"`
	Version         int                        `bson:"version" json:"version"`
}

type InvestigationAttachments struct {
	Id string `bson:"id" json:"id"`
}

type InvestigationComments struct {
	Id string `bson:"id" json:"id"`
}

type InvestigationResponse struct {
	Id              string                     `json:"id"`
	AssignedTo      string                     `json:"assignedTo"`
	Attachments     []InvestigationAttachments `json:"attachments"`
	CreatedBy       string                     `json:"createdBy"`
	CreatedAt       time.Time                  `json:"createdAt"`
	Comments        []InvestigationComments    `json:"comments"`
	Description     string                     `json:"description"`
	InvestigationId string                     `json:"investigationId"`
	Severity        string                     `json:"severity"`
	Status          string                     `json:"status"`
	Tags            []string                   `json:"tags"`
	Title           string                     `json:"title"`
	Tlp             int                        `json:"tlp"`
}

type CreateInvestigationRequest struct {
	AssignedTo  string   `json:"assignedTo"`
	Description string   `json:"description"`
	Severity    string   `json:"severity"`
	Status      string   `json:"status"`
	Tags        []string `json:"tags"`
	TemplateId  string   `json:"templateId"`
	Title       string   `json:"title"`
	Tlp         int      `json:"tlp"`
}

type UpdateInvestigationRequest struct {
	AssignedTo  string   `json:"assignedTo"`
	Description string   `json:"description"`
	Severity    string   `json:"severity"`
	Status      string   `json:"status"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
	Tlp         int      `json:"tlp"`
}
