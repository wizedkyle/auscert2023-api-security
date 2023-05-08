package mocks

import (
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/investigationmodel"
	"time"
)

func Investigation() *investigationmodel.Investigation {
	return &investigationmodel.Investigation{
		Id:              "f1406436-c890-45d4-85a8-da69194fe33b",
		AssignedTo:      "36235e81-e789-49cd-b3d8-a7d164982a50",
		Attachments:     nil,
		CreatedBy:       "36235e81-e789-49cd-b3d8-a7d164982a50",
		CreatedAt:       time.Time{},
		Comments:        nil,
		Description:     "Test Description",
		InvestigationId: "1",
		Severity:        "critical",
		Status:          "in-progress",
		Tags:            []string{"test", "tag"},
		TenantId:        "",
		Title:           "Test Investigation",
		Tlp:             3,
		Version:         1,
	}
}

func Investigation2() *investigationmodel.Investigation {
	return &investigationmodel.Investigation{
		Id:              "fa70bc9e-d74c-41e6-9e24-65be21d9148b",
		AssignedTo:      "0c3dc9af-a9d8-4368-80e3-390f0a48799c",
		Attachments:     nil,
		CreatedBy:       "6963b733-4fdd-4c3f-ab44-13c085d1890e",
		CreatedAt:       time.Time{},
		Comments:        nil,
		Description:     "Test Description 2",
		InvestigationId: "1",
		Severity:        "critical",
		Status:          "in-progress",
		Tags:            []string{"test", "tag"},
		TenantId:        "",
		Title:           "Test Investigation 2",
		Tlp:             3,
		Version:         1,
	}
}

func InvestigationResponse() *investigationmodel.InvestigationResponse {
	return &investigationmodel.InvestigationResponse{
		Id:              "f1406436-c890-45d4-85a8-da69194fe33b",
		AssignedTo:      "36235e81-e789-49cd-b3d8-a7d164982a50",
		Attachments:     nil,
		CreatedBy:       "36235e81-e789-49cd-b3d8-a7d164982a50",
		CreatedAt:       time.Time{},
		Comments:        nil,
		Description:     "Test Description",
		InvestigationId: "1",
		Severity:        "critical",
		Status:          "in-progress",
		Tags:            []string{"test", "tag"},
		Title:           "Test Investigation",
		Tlp:             3,
	}
}

func Investigations() *[]investigationmodel.InvestigationResponse {
	return &[]investigationmodel.InvestigationResponse{
		{
			Id:              "f1406436-c890-45d4-85a8-da69194fe33b",
			AssignedTo:      "36235e81-e789-49cd-b3d8-a7d164982a50",
			Attachments:     nil,
			CreatedBy:       "36235e81-e789-49cd-b3d8-a7d164982a50",
			CreatedAt:       time.Time{},
			Comments:        nil,
			Description:     "Test Description",
			InvestigationId: "1",
			Severity:        "critical",
			Status:          "in-progress",
			Tags:            []string{"test", "tag"},
			Title:           "Test Investigation",
			Tlp:             3,
		},
		{
			Id:              "fa70bc9e-d74c-41e6-9e24-65be21d9148b",
			AssignedTo:      "0c3dc9af-a9d8-4368-80e3-390f0a48799c",
			Attachments:     nil,
			CreatedBy:       "6963b733-4fdd-4c3f-ab44-13c085d1890e",
			CreatedAt:       time.Time{},
			Comments:        nil,
			Description:     "Test Description 2",
			InvestigationId: "1",
			Severity:        "critical",
			Status:          "in-progress",
			Tags:            []string{"test", "tag"},
			Title:           "Test Investigation 2",
			Tlp:             3,
		},
	}
}
