package incidentmodel

type Incident struct {
	Id          string                `bson:"id" json:"id"`
	AssignedTo  string                `bson:"assignedTo" json:"assignedTo"`
	Attachments []IncidentAttachments `bson:"attachments" json:"attachments"`
	CreatedBy   string                `bson:"createdBy" json:"createdBy"`
	Description string                `bson:"description" json:"description"`
	Severity    string                `bson:"severity" json:"severity"`
	Status      string                `bson:"status" json:"status"`
	Tags        []string              `bson:"tags" json:"tags"`
	Tasks       []IncidentTasks       `bson:"tasks" json:"tasks"`
	TenantId    string                `bson:"tenantId" json:"tenantId"`
	Title       string                `bson:"title" json:"title"`
	Tlp         int                   `bson:"tlp" json:"tlp"`
	Version     int                   `bson:"version" json:"version"`
}

type IncidentAttachments struct {
	Id string `bson:"id" json:"id"`
}

type IncidentTasks struct {
	Id string `bson:"id" json:"id"`
}

type IncidentResponse struct {
	Id          string                `json:"id"`
	AssignedTo  string                `json:"assignedTo"`
	Attachments []IncidentAttachments `json:"attachments"`
	CreatedBy   string                `json:"createdBy"`
	Description string                `json:"description"`
	Severity    string                `json:"severity"`
	Status      string                `json:"status"`
	Tags        []string              `json:"tags"`
	Tasks       []IncidentTasks       `json:"tasks"`
	Title       string                `json:"title"`
	Tlp         int                   `json:"tlp"`
}

type CreateIncidentRequest struct {
	AssignedTo  string   `json:"assignedTo"`
	Description string   `json:"description"`
	Severity    string   `json:"severity"`
	Status      string   `json:"status"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
	Tlp         int      `json:"tlp"`
}

type UpdateIncidentRequest struct {
	AssignedTo  string   `json:"assignedTo"`
	Description string   `json:"description"`
	Severity    string   `json:"severity"`
	Status      string   `json:"status"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
	Tlp         int      `json:"tlp"`
}
