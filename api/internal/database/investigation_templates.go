package database

import (
	"context"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/investigationtemplatemodel"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateInvestigationTemplate
// Creates a new investigation template in the database.
func (d *database) CreateInvestigationTemplate(tenantId string, template investigationtemplatemodel.InvestigationTemplate) (*investigationtemplatemodel.InvestigationTemplateResponse, error) {
	var result investigationtemplatemodel.InvestigationTemplateResponse
	template.Id = uuid.NewString()
	_, err := d.InvestigationTemplatesCollection.InsertOne(context.TODO(), template)
	if err != nil {
		return nil, err
	}
	err = d.InvestigationTemplatesCollection.FindOne(context.TODO(), bson.D{{"id", template.Id}, {"tenantId", tenantId}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteInvestigationTemplate
// Deletes an investigation template from the database.
func (d *database) DeleteInvestigationTemplate(id string, tenantId string) error {
	result, err := d.InvestigationTemplatesCollection.DeleteOne(context.TODO(), bson.D{{"id", id}, {"tenantId", tenantId}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return utils.ErrInvestigationTemplateNotFound
	}
	return nil
}

// ListInvestigationTemplates
// Returns all investigation templates from the database.
func (d *database) ListInvestigationTemplates(limit int64, tenantId string) ([]investigationtemplatemodel.InvestigationTemplateResponse, error) {
	var investigationTemplates []investigationtemplatemodel.InvestigationTemplateResponse
	filter := bson.D{{"tenantId", tenantId}}
	opts := options.Find().SetLimit(limit)
	cursor, err := d.InvestigationTemplatesCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var investigationTemplate investigationtemplatemodel.InvestigationTemplateResponse
		err := cursor.Decode(&investigationTemplate)
		if err != nil {
			return nil, err
		}
		investigationTemplates = append(investigationTemplates, investigationTemplate)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(context.TODO())
	if len(investigationTemplates) == 0 {
		return nil, utils.ErrInvestigationTemplatesNotFound
	}
	return investigationTemplates, nil
}

// GetInvestigationTemplate
// Returns a single investigation template from the database.
func (d *database) GetInvestigationTemplate(id string, tenantId string) (*investigationtemplatemodel.InvestigationTemplateResponse, error) {
	var result investigationtemplatemodel.InvestigationTemplateResponse
	err := d.InvestigationTemplatesCollection.FindOne(context.TODO(), bson.D{{"id", id}, {"tenantId", tenantId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrInvestigationTemplateNotFound
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateInvestigationTemplate
// Updates an investigation in the database.
func (d *database) UpdateInvestigationTemplate(id string, tenantId string, template investigationtemplatemodel.UpdateInvestigationRequest) (*investigationtemplatemodel.InvestigationTemplateResponse, error) {
	var (
		oldResult investigationtemplatemodel.InvestigationTemplateResponse
		newResult investigationtemplatemodel.InvestigationTemplateResponse
	)
	filter := bson.D{{"id", id}, {"tenantId", tenantId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "description", Value: template.Status},
			primitive.E{Key: "severity", Value: template.Severity},
			primitive.E{Key: "tags", Value: template.Tags},
			primitive.E{Key: "titlePrefix", Value: template.TitlePrefix},
			primitive.E{Key: "tlp", Value: template.Tlp},
		}},
	}
	err := d.InvestigationTemplatesCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrInvestigationTemplateNotFound
	} else if err != nil {
		return nil, err
	}
	err = d.InvestigationTemplatesCollection.FindOne(context.TODO(), bson.D{{"id", oldResult.Id}, {"tenantId", tenantId}}).Decode(&newResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrInvestigationTemplateNotFound
	} else if err != nil {
		return nil, err
	}
	return &newResult, nil
}
