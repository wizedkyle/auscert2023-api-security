package webhookmodel

type Webhook struct {
	Id          string   `bson:"id" json:"id"`
	Algorithm   string   `bson:"algorithm" json:"algorithm"`
	Description string   `bson:"description" json:"description"`
	Events      []string `bson:"events" json:"events"`
	Secret      string   `bson:"secret" json:"secret"`
	TenantId    string   `bson:"tenantId" json:"tenantId"`
	Url         string   `bson:"url" json:"url"`
	Version     int      `bson:"version" json:"version"`
}

type WebhookResponse struct {
	Id          string   `json:"id"`
	Algorithm   string   `json:"algorithm"`
	Description string   `json:"description"`
	Events      []string `json:"events"`
	Secret      string   `json:"secret"`
	Url         string   `json:"url"`
}

type CreateWebhookRequest struct {
	Description string   `json:"description" binding:"required"`
	Events      []string `json:"events"`
	Url         string   `json:"url" binding:"required"`
}

type UpdateWebhookRequest struct {
	Description string   `json:"description"`
	Events      []string `json:"events"`
	Url         string   `json:"url"`
}
