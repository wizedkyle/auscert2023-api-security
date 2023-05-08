package utils

import "github.com/pkg/errors"

type Error struct {
	InternalError error         `json:"internalError"`
	ExternalError ExternalError `json:"externalError"`
}

type ExternalError struct {
	Id            string `json:"id"`
	Message       string `json:"message"`
	Code          int    `json:"code"`
	TransactionId string `json:"transactionId"`
}

const (
	AccessKeyNotFound                 string = "access key not found"
	AccessKeysNotFound                string = "access keys not found"
	EmailAlreadyExists                string = "email address already exists"
	GenericInternalServerErrorMessage string = "internal server error"
	InvalidRequestBody                string = "invalid request body"
	InvalidNumberOfEvents             string = "number of events provided exceeds maximum allowable events"
	InvalidNumberOfScopes             string = "number of scopes provided exceeds maximum allowable scopes"
	InvalidEvent                      string = " is not a valid event"
	InvalidScope                      string = " is not a valid scope"
	IncidentNotFound                  string = "incident not found"
	IncidentsNotFound                 string = "incidents not found"
	IncidentCommentNotFound           string = "incident comment not found"
	IncidentCommentsNotFound          string = "incident comments not found"
	InvestigationNotFound             string = "investigation not found"
	InvestigationsNotFound            string = "investigations not found"
	InvestigationTemplateNotFound     string = "investigation template not found"
	InvestigationTemplatesNotFound    string = "investigation templates not found"
	TenantNotFound                    string = "tenant not found"
	UserNotFound                      string = "user not found"
	UsersNotFound                     string = "users not found"
	WebhookNotFound                   string = "webhook not found"
	WebhooksNotFound                  string = "webhooks not found"
)

var (
	ErrAccessKeyNotFound              error = errors.New(AccessKeyNotFound)
	ErrAccessKeysNotFound             error = errors.New(AccessKeysNotFound)
	ErrEmailAlreadyExists             error = errors.New(EmailAlreadyExists)
	ErrIncidentNotFound               error = errors.New(IncidentNotFound)
	ErrIncidentsNotFound              error = errors.New(IncidentsNotFound)
	ErrIncidentCommentNotFound        error = errors.New(IncidentCommentNotFound)
	ErrIncidentCommentsNotFound       error = errors.New(IncidentCommentsNotFound)
	ErrInvestigationNotFound          error = errors.New(InvestigationNotFound)
	ErrInvestigationsNotFound         error = errors.New(InvestigationsNotFound)
	ErrInvestigationTemplateNotFound  error = errors.New(InvestigationTemplateNotFound)
	ErrInvestigationTemplatesNotFound error = errors.New(InvestigationTemplatesNotFound)
	ErrTenantNotFound                 error = errors.New(TenantNotFound)
	ErrUserNotFound                   error = errors.New(UserNotFound)
	ErrUsersNotFound                  error = errors.New(UsersNotFound)
	ErrWebhookNotFound                error = errors.New(WebhookNotFound)
	ErrWebhooksNotFound               error = errors.New(WebhooksNotFound)
)

// GenerateError
// Generates an error message that is used when processing API requests.
func GenerateError(id string, message string, code int, transactionId string, err error) *Error {
	return &Error{
		InternalError: err,
		ExternalError: ExternalError{
			Id:            id,
			Message:       message,
			Code:          code,
			TransactionId: transactionId,
		},
	}
}
