package mocks

import "github.com/ForgeResponse/ForgeResponse/v2/pkg/model/incidentmodel"

func Incident() *incidentmodel.Incident {
	return &incidentmodel.Incident{
		Id:          "87e31788-2625-4bb2-ae21-004e07f76d45",
		AssignedTo:  "36235e81-e789-49cd-b3d8-a7d164982a50",
		Attachments: nil,
		CreatedBy:   "36235e81-e789-49cd-b3d8-a7d164982a50",
		Description: "Test Description",
		Severity:    "critical",
		Status:      "in-progress",
		Tags:        []string{"test", "tag"},
		Tasks:       nil,
		TenantId:    "",
		Title:       "Test Incident",
		Tlp:         3,
		Version:     1,
	}
}

func Incident2() *incidentmodel.Incident {
	return &incidentmodel.Incident{
		Id:          "670dbf0d-fe67-4b25-8dc7-22c2557fbf06",
		AssignedTo:  "36235e81-e789-49cd-b3d8-a7d164982a50",
		Attachments: nil,
		CreatedBy:   "36235e81-e789-49cd-b3d8-a7d164982a50",
		Description: "Test Description",
		Severity:    "critical",
		Status:      "in-progress",
		Tags:        []string{"test", "tag"},
		Tasks:       nil,
		TenantId:    "",
		Title:       "Test Incident",
		Tlp:         3,
		Version:     1,
	}
}

func IncidentResponse() *incidentmodel.IncidentResponse {
	return &incidentmodel.IncidentResponse{
		Id:          "87e31788-2625-4bb2-ae21-004e07f76d45",
		AssignedTo:  "36235e81-e789-49cd-b3d8-a7d164982a50",
		Attachments: nil,
		CreatedBy:   "36235e81-e789-49cd-b3d8-a7d164982a50",
		Description: "Test Description",
		Severity:    "critical",
		Status:      "in-progress",
		Tags:        []string{"test", "tag"},
		Tasks:       nil,
		Title:       "Test Incident",
		Tlp:         3,
	}
}

func Incidents() *[]incidentmodel.IncidentResponse {
	return &[]incidentmodel.IncidentResponse{
		{
			Id:          "87e31788-2625-4bb2-ae21-004e07f76d45",
			AssignedTo:  "36235e81-e789-49cd-b3d8-a7d164982a50",
			Attachments: nil,
			CreatedBy:   "36235e81-e789-49cd-b3d8-a7d164982a50",
			Description: "Test Description",
			Severity:    "critical",
			Status:      "in-progress",
			Tags:        []string{"test", "tag"},
			Tasks:       nil,
			Title:       "Test Incident",
			Tlp:         3,
		},
		{
			Id:          "670dbf0d-fe67-4b25-8dc7-22c2557fbf06",
			AssignedTo:  "36235e81-e789-49cd-b3d8-a7d164982a50",
			Attachments: nil,
			CreatedBy:   "36235e81-e789-49cd-b3d8-a7d164982a50",
			Description: "Test Description",
			Severity:    "critical",
			Status:      "in-progress",
			Tags:        []string{"test", "tag"},
			Tasks:       nil,
			Title:       "Test Incident",
			Tlp:         3,
		},
	}
}
