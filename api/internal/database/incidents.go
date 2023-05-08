package database

import (
	"context"

	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/incidentmodel"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateIncident
// Creates a new incident in the database.
func (d *database) CreateIncident(tenantId string, incident incidentmodel.Incident) (*incidentmodel.IncidentResponse, error) {
	var result incidentmodel.IncidentResponse
	incident.Id = uuid.NewString()
	_, err := d.IncidentCollection.InsertOne(context.TODO(), incident)
	if err != nil {
		return nil, err
	}
	err = d.IncidentCollection.FindOne(context.TODO(), bson.D{{"id", incident.Id}, {"tenantId", tenantId}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteIncident
// Deletes an incident from the database.
func (d *database) DeleteIncident(id string, tenantId string) error {
	result, err := d.IncidentCollection.DeleteOne(context.TODO(), bson.D{{"id", id}, {"tenantId", tenantId}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return utils.ErrIncidentNotFound
	}
	return err
}

// ListIncidents
// Returns all incidents from the database.
func (d *database) ListIncidents(limit int64, tenantId string) ([]incidentmodel.IncidentResponse, error) {
	var incidents []incidentmodel.IncidentResponse
	filter := bson.D{{"tenantId", tenantId}}
	opts := options.Find().SetLimit(limit)
	cursor, err := d.IncidentCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var incident incidentmodel.IncidentResponse
		err := cursor.Decode(&incident)
		if err != nil {
			return nil, err
		}
		incidents = append(incidents, incident)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(context.TODO())
	if len(incidents) == 0 {
		return nil, utils.ErrIncidentsNotFound
	}
	return incidents, nil
}

// GetIncident
// Returns a single incident from the database.
func (d *database) GetIncident(id string, tenantId string) (*incidentmodel.IncidentResponse, error) {
	var result incidentmodel.IncidentResponse
	err := d.IncidentCollection.FindOne(context.TODO(), bson.D{{"id", id}, {"tenantId", tenantId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrIncidentNotFound
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateIncident
// Updates an incident in the database.
func (d *database) UpdateIncident(id string, tenantId string, incident incidentmodel.UpdateIncidentRequest) (*incidentmodel.IncidentResponse, error) {
	var (
		oldResult incidentmodel.IncidentResponse
		newResult incidentmodel.IncidentResponse
	)
	filter := bson.D{{"id", id}, {"tenantId", tenantId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "assignedTo", Value: incident.AssignedTo},
			primitive.E{Key: "description", Value: incident.Description},
			primitive.E{Key: "severity", Value: incident.Severity},
			primitive.E{Key: "status", Value: incident.Status},
			primitive.E{Key: "tags", Value: incident.Tags},
			primitive.E{Key: "title", Value: incident.Title},
			primitive.E{Key: "tlp", Value: incident.Tlp},
		}},
	}
	err := d.IncidentCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrIncidentNotFound
	} else if err != nil {
		return nil, err
	}
	err = d.IncidentCollection.FindOne(context.TODO(), bson.D{{"id", id}, {"tenantId", tenantId}}).Decode(&newResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrIncidentNotFound
	} else if err != nil {
		return nil, err
	}
	return &newResult, nil
}
