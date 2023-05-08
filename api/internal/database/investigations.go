package database

import (
	"context"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/investigationmodel"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateInvestigation
// Creates a new investigation in the database.
func (d *database) CreateInvestigation(tenantId string, investigation investigationmodel.Investigation) (*investigationmodel.InvestigationResponse, error) {
	var result investigationmodel.InvestigationResponse
	investigation.Id = uuid.NewString()
	_, err := d.InvestigationCollection.InsertOne(context.TODO(), investigation)
	if err != nil {
		return nil, err
	}
	err = d.InvestigationCollection.FindOne(context.TODO(), bson.D{{"id", investigation.Id}, {"tenantId", tenantId}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteInvestigation
// Deletes an investigation from the database.
func (d *database) DeleteInvestigation(id string, tenantId string) error {
	result, err := d.InvestigationCollection.DeleteOne(context.TODO(), bson.D{{"id", id}, {"tenantId", tenantId}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return utils.ErrInvestigationNotFound
	}
	return err
}

// ListInvestigations
// Returns all investigations from the database.
func (d *database) ListInvestigations(limit int64, tenantId string) ([]investigationmodel.InvestigationResponse, error) {
	var investigations []investigationmodel.InvestigationResponse
	filter := bson.D{{"tenantId", tenantId}}
	opts := options.Find().SetLimit(limit)
	cursor, err := d.InvestigationCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var investigation investigationmodel.InvestigationResponse
		err := cursor.Decode(&investigation)
		if err != nil {
			return nil, err
		}
		investigations = append(investigations, investigation)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(context.TODO())
	if len(investigations) == 0 {
		return nil, utils.ErrInvestigationsNotFound
	}
	return investigations, nil
}

// GetInvestigation
// Returns a single investigation from the database.
func (d *database) GetInvestigation(id string, tenantId string) (*investigationmodel.InvestigationResponse, error) {
	var result investigationmodel.InvestigationResponse
	err := d.InvestigationCollection.FindOne(context.TODO(), bson.D{{"id", id}, {"tenantId", tenantId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrInvestigationNotFound
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateInvestigation
// Updates an investigation in the database.
func (d *database) UpdateInvestigation(id string, tenantId string, investigation investigationmodel.UpdateInvestigationRequest) (*investigationmodel.InvestigationResponse, error) {
	var (
		oldResult investigationmodel.InvestigationResponse
		newResult investigationmodel.InvestigationResponse
	)
	filter := bson.D{{"id", id}, {"tenantId", tenantId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "description", Value: investigation.Description},
			primitive.E{Key: "severity", Value: investigation.Severity},
			primitive.E{Key: "status", Value: investigation.Status},
			primitive.E{Key: "tags", Value: investigation.Tags},
			primitive.E{Key: "title", Value: investigation.Title},
			primitive.E{Key: "tlp", Value: investigation.Tlp},
		}},
	}
	err := d.InvestigationCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrInvestigationNotFound
	} else if err != nil {
		return nil, err
	}
	err = d.InvestigationCollection.FindOne(context.TODO(), bson.D{{"id", oldResult.Id}, {"tenantId", tenantId}}).Decode(&newResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrInvestigationNotFound
	} else if err != nil {
		return nil, err
	}
	return &newResult, nil
}
