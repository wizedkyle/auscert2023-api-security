package mocks

import (
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/investigationtemplatemodel"
	"time"
)

func InvestigationTemplate() *investigationtemplatemodel.InvestigationTemplate {
	return &investigationtemplatemodel.InvestigationTemplate{
		Id:          "ffd05e2d-4efb-4882-b1c1-2ee2bea26422",
		TenantId:    "",
		CreatedBy:   "36235e81-e789-49cd-b3d8-a7d164982a50",
		CreatedAt:   time.Time{},
		TitlePrefix: "Test prefix",
		Severity:    "critical",
		Status:      "in-progress",
		Tags:        []string{"test", "tag"},
		Tlp:         3,
	}
}

func InvestigationTemplate2() *investigationtemplatemodel.InvestigationTemplate {
	return &investigationtemplatemodel.InvestigationTemplate{
		Id:          "44940f00-feea-4a53-98ad-2d1e8d24cbfd",
		TenantId:    "",
		CreatedBy:   "36235e81-e789-49cd-b3d8-a7d164982a50",
		CreatedAt:   time.Time{},
		TitlePrefix: "Test prefix 2",
		Severity:    "critical",
		Status:      "in-progress",
		Tags:        []string{"test", "tag"},
		Tlp:         3,
	}
}

func InvestigationTemplateResponse() *investigationtemplatemodel.InvestigationTemplateResponse {
	return &investigationtemplatemodel.InvestigationTemplateResponse{
		Id:          "ffd05e2d-4efb-4882-b1c1-2ee2bea26422",
		CreatedBy:   "36235e81-e789-49cd-b3d8-a7d164982a50",
		CreatedAt:   time.Time{},
		TitlePrefix: "Test prefix",
		Severity:    "critical",
		Status:      "in-progress",
		Tags:        []string{"test", "tag"},
		Tlp:         3,
	}
}

func InvestigationTemplates() *[]investigationtemplatemodel.InvestigationTemplateResponse {
	return &[]investigationtemplatemodel.InvestigationTemplateResponse{
		{
			Id:          "ffd05e2d-4efb-4882-b1c1-2ee2bea26422",
			CreatedBy:   "36235e81-e789-49cd-b3d8-a7d164982a50",
			CreatedAt:   time.Time{},
			TitlePrefix: "Test prefix",
			Severity:    "critical",
			Status:      "in-progress",
			Tags:        []string{"test", "tag"},
			Tlp:         3,
		},
		{
			Id:          "44940f00-feea-4a53-98ad-2d1e8d24cbfd",
			CreatedBy:   "36235e81-e789-49cd-b3d8-a7d164982a50",
			CreatedAt:   time.Time{},
			TitlePrefix: "Test prefix 2",
			Severity:    "critical",
			Status:      "in-progress",
			Tags:        []string{"test", "tag"},
			Tlp:         3,
		},
	}
}
