package database

import (
	"context"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/incidentcommentmodel"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateIncidentComment
// Creates a new comment against the specified incident in the database.
func (d *database) CreateIncidentComment(tenantId string, comment incidentcommentmodel.IncidentComment) (*incidentcommentmodel.IncidentCommentResponse, error) {
	var (
		latestRecord incidentcommentmodel.IncidentComment
		result       incidentcommentmodel.IncidentCommentResponse
	)
	comment.Id = uuid.NewString()
	opts := options.FindOne().SetSort(bson.M{"$natural": -1})
	err := d.IncidentCommentsCollection.FindOne(context.TODO(), bson.D{{Key: "tenantId", Value: tenantId}, {Key: "incidentId", Value: comment.IncidentId}}, opts).Decode(&latestRecord)
	if err == mongo.ErrNoDocuments {
		comment.Order = 0
	} else if err == nil {
		comment.Order = latestRecord.Order + 1
	} else {
		return nil, err
	}
	_, err = d.IncidentCommentsCollection.InsertOne(context.TODO(), comment)
	if err != nil {
		return nil, err
	}
	err = d.IncidentCommentsCollection.FindOne(context.TODO(), bson.D{{Key: "id", Value: comment.Id}, {Key: "tenantId", Value: tenantId}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteIncidentComment
// Deletes a comment from the specified incident.
func (d *database) DeleteIncidentComment(commentId string, incidentId string, tenantId string) error {
	result, err := d.IncidentCommentsCollection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: commentId}, {Key: "incidentId", Value: incidentId}, {Key: "tenantId", Value: tenantId}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return utils.ErrIncidentCommentNotFound
	}
	return err
}

// ListIncidentComments
// Returns all comments associated to the specific incident.
func (d *database) ListIncidentComments(limit int64, incidentId string, tenantId string) ([]incidentcommentmodel.IncidentCommentResponse, error) {
	var comments []incidentcommentmodel.IncidentCommentResponse
	filter := bson.D{{Key: "incidentId", Value: incidentId}, {Key: "tenantId", Value: tenantId}}
	opts := options.Find().SetLimit(limit)
	cursor, err := d.IncidentCommentsCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var comment incidentcommentmodel.IncidentCommentResponse
		err := cursor.Decode(&comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(context.TODO())
	if len(comments) == 0 {
		return nil, utils.ErrIncidentCommentsNotFound
	}
	return comments, nil
}

// GetIncidentComment
// Returns a single incident comment from the specific incident.
func (d *database) GetIncidentComment(commentId string, incidentId string, tenantId string) (*incidentcommentmodel.IncidentCommentResponse, error) {
	var result incidentcommentmodel.IncidentCommentResponse
	err := d.IncidentCommentsCollection.FindOne(context.TODO(), bson.D{
		{Key: "id", Value: commentId},
		{Key: "incidentId", Value: incidentId},
		{Key: "tenantId", Value: tenantId},
	}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrIncidentCommentNotFound
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateIncidentComment
// Updates an incident comment against the specified incident in the database.
func (d *database) UpdateIncidentComment(commentId string, incidentId string, tenantId string, comment incidentcommentmodel.UpdateIncidentComment) (*incidentcommentmodel.IncidentCommentResponse, error) {
	var (
		oldResult incidentcommentmodel.IncidentCommentResponse
		newResult incidentcommentmodel.IncidentCommentResponse
	)
	filter := bson.D{
		{Key: "id", Value: commentId},
		{Key: "incidentId", Value: incidentId},
		{Key: "tenantId", Value: tenantId},
	}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "comment", Value: comment.Comment},
		}},
	}
	err := d.IncidentCommentsCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrIncidentCommentNotFound
	} else if err != nil {
		return nil, err
	}
	err = d.IncidentCommentsCollection.FindOne(context.TODO(), filter).Decode(&newResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrIncidentCommentNotFound
	} else if err != nil {
		return nil, err
	}
	return &newResult, nil
}
