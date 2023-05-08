package incidentcommentmodel

import "time"

type IncidentComment struct {
	Id         string    `bson:"id" json:"id"`
	IncidentId string    `bson:"incidentId" json:"incidentId"`
	Comment    string    `bson:"comment" json:"comment"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
	CreatedBy  string    `bson:"createdBy" json:"createdBy"`
	Order      int       `bson:"order" json:"order"`
	TenantId   string    `bson:"tenantId" json:"tenantId"`
}

type IncidentCommentResponse struct {
	Id        string    `json:"id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Order     int       `json:"order"`
}

type CreateIncidentComment struct {
	Comment string `json:"comment" binding:"required"`
}

type UpdateIncidentComment struct {
	Comment string `json:"comment" binding:"required"`
}
