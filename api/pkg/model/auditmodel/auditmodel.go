package auditmodel

import "time"

const (
	// Audit Log Actions
	AuditLogRetrieved = "audit_log_retrieved"

	// Access Key Actions
	AccessKeyDeleted   = "access_key_deleted"
	AccessKeyUpdated   = "access_key_updated"
	AccessKeyRetrieved = "access_keys_retrieved"
	AccessKeyCreated   = "access_key_created"
	AccessKeyRotated   = "access_key_rotated"

	// Incident Actions
	IncidentDeleted   = "incident_deleted"
	IncidentUpdated   = "incident_updated"
	IncidentRetrieved = "incident_retrieved"
	IncidentCreated   = "incident_created"

	// Incident Comment Actions
	IncidentCommentDeleted   = "incident_comment_deleted"
	IncidentCommentUpdated   = "incident_comment_updated"
	IncidentCommentRetrieved = "incident_comment_retrieved"
	IncidentCommentCreated   = "incident_comment_created"

	// Investigation Actions
	InvestigationDeleted   = "investigation_deleted"
	InvestigationUpdated   = "investigation_updated"
	InvestigationRetrieved = "investigation_retrieved"
	InvestigationCreated   = "investigation_created"

	// Investigation Template Actions
	InvestigationTemplateDeleted   = "investigation_template_deleted"
	InvestigationTemplateUpdated   = "investigation_template_updated"
	InvestigationTemplateRetrieved = "investigation_template_retrieved"
	InvestigationTemplateCreated   = "investigation_template_created"

	// Tenant Actions
	TenantDeleted   = "tenant_deleted"
	TenantUpdated   = "tenant_updated"
	TenantRetrieved = "tenant_retrieved"
	TenantCreated   = "tenant_created"

	// User Actions
	UserDeleted   = "user_deleted"
	UserUpdated   = "user_updated"
	UserRetrieved = "user_retrieved"
	UserCreated   = "user_created"

	// Webhook Actions
	WebhookDeleted   = "webhook_deleted"
	WebhookUpdated   = "webhook_updated"
	WebhookRetrieved = "webhook_retrieved"
	WebhookCreated   = "webhook_created"
	WebhookRotated   = "webhook_rotated"
)

type AuditResponse struct {
	Events []string `json:"events"`
}

var AuditScopes = []string{
	// Access Key Actions
	AccessKeyCreated, AccessKeyDeleted, AccessKeyUpdated, AccessKeyRotated,
	// Investigation Actions
	InvestigationCreated, InvestigationDeleted, InvestigationUpdated,
	// Investigation Template Actions
	InvestigationTemplateCreated, InvestigationTemplateDeleted, InvestigationTemplateUpdated,
	// Users Actions
	UserCreated, UserDeleted, UserUpdated,
}

type Audit struct {
	Id        string     `json:"id"`
	TenantId  string     `json:"tenantId"`
	Time      time.Time  `json:"time"`
	Action    string     `json:"action"`
	Subject   string     `json:"subject"`
	Actor     AuditActor `json:"actor"`
	UserAgent string     `json:"userAgent"`
	IpAddress string     `json:"ipAddress"`
}

type AuditActor struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
