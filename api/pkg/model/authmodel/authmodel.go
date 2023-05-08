package authmodel

const (
	// Audit Log Scopes
	ReadAuditLog = "read:auditlog"

	// Access Key Scopes
	DeleteAccessKeys = "delete:accesskeys"
	ModifyAccessKeys = "modify:accesskeys"
	ReadAccessKeys   = "read:accesskeys"
	WriteAccessKeys  = "write:accesskeys"

	// Event Scopes
	ReadEvents = "read:events"

	// Incident Scopes
	DeleteIncident = "delete:incident"
	ModifyIncident = "modify:incident"
	ReadIncident   = "read:incident"
	WriteIncident  = "write:incident"

	// Incident Comment Scopes
	DeleteIncidentComment = "delete:incidentcomment"
	ModifyIncidentComment = "modify:incidentcomment"
	ReadIncidentComment   = "read:incidentcomment"
	WriteIncidentComment  = "write:incidentcomment"

	// Investigation Scopes
	DeleteInvestigations = "delete:investigations"
	ModifyInvestigations = "modify:investigations"
	ReadInvestigations   = "read:investigations"
	WriteInvestigations  = "write:investigations"

	// Investigation Template Scopes
	DeleteInvestigationTemplates = "delete:investigationtemplates"
	ModifyInvestigationTemplates = "modify:investigationtemplates"
	ReadInvestigationTemplates   = "read:investigationtemplates"
	WriteInvestigationTemplates  = "write:investigationtemplates"

	// Internal Scopes
	ReadScopes = "read:scopes"

	// Tenant Scopes
	DeleteTenants = "delete:tenants"
	ModifyTenants = "modify:tenants"
	ReadTenants   = "read:tenants"
	WriteTenants  = "write:tenants"

	// Users Scopes
	DeleteUsers = "delete:users"
	ModifyUsers = "modify:users"
	ReadUsers   = "read:users"
	WriteUsers  = "write:users"

	// Webhook Scopes
	DeleteWebhooks = "delete:webhooks"
	ModifyWebhooks = "modify:webhooks"
	ReadWebhooks   = "read:webhooks"
	WriteWebhooks  = "write:webhooks"
)

type ScopeResponse struct {
	Scopes []string `json:"scopes"`
}

var Scopes = []string{
	// Access Key Scopes
	DeleteAccessKeys, ModifyAccessKeys, ReadAccessKeys, WriteAccessKeys,
	// Evengt Scopes
	ReadEvents,
	// Investigation Scopes
	DeleteInvestigations, ModifyInvestigations, ReadInvestigations, WriteInvestigations,
	// Investigation Template Scopes
	DeleteInvestigationTemplates, ModifyInvestigationTemplates, ReadInvestigationTemplates, WriteInvestigationTemplates,
	// Scopes
	ReadScopes,
	// Tenant Scopes
	DeleteTenants, ModifyTenants, ReadTenants, WriteTenants,
	// Users Scopes
	DeleteUsers, ModifyUsers, ReadUsers, WriteUsers,
	// Webhook Scopes
	DeleteWebhooks, ModifyWebhooks, ReadWebhooks, WriteWebhooks,
}
