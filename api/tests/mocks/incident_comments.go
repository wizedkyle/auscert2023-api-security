package mocks

import (
	"time"

	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/incidentcommentmodel"
)

func IncidentComment() *incidentcommentmodel.IncidentComment {
	return &incidentcommentmodel.IncidentComment{
		Id:         "86e9e3d0-cb3f-49a5-8074-819164a8debd",
		IncidentId: "87e31788-2625-4bb2-ae21-004e07f76d45",
		Comment:    "Test Comment",
		CreatedAt:  time.Time{},
		CreatedBy:  "",
		Order:      0,
		TenantId:   "",
	}
}

func IncidentComment2() *incidentcommentmodel.IncidentComment {
	return &incidentcommentmodel.IncidentComment{
		Id:         "5c836f45-120d-42f6-b000-4c8005bc42e5",
		IncidentId: "87e31788-2625-4bb2-ae21-004e07f76d45",
		Comment:    "Test Comment",
		CreatedAt:  time.Time{},
		CreatedBy:  "",
		Order:      1,
		TenantId:   "",
	}
}

func IncidentCommentResponse() *incidentcommentmodel.IncidentCommentResponse {
	return &incidentcommentmodel.IncidentCommentResponse{
		Id:        "86e9e3d0-cb3f-49a5-8074-819164a8debd",
		Comment:   "Test Comment",
		CreatedAt: time.Time{},
		CreatedBy: "",
		Order:     0,
	}
}

func IncidentComments() *[]incidentcommentmodel.IncidentCommentResponse {
	return &[]incidentcommentmodel.IncidentCommentResponse{
		{
			Id:        "86e9e3d0-cb3f-49a5-8074-819164a8debd",
			Comment:   "Test Comment",
			CreatedAt: time.Time{},
			CreatedBy: "",
			Order:     0,
		},
		{
			Id:        "5c836f45-120d-42f6-b000-4c8005bc42e5",
			Comment:   "Test Comment",
			CreatedAt: time.Time{},
			CreatedBy: "",
			Order:     1,
		},
	}
}
