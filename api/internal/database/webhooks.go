package database

import (
	"context"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/webhookmodel"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateWebhook
// Creates a new webhook.
func (d *database) CreateWebhook(tenantId string, webhook webhookmodel.Webhook) (*webhookmodel.WebhookResponse, error) {
	var result webhookmodel.WebhookResponse
	webhook.Id = uuid.NewString()
	_, err := d.WebhooksCollection.InsertOne(context.TODO(), webhook)
	if err != nil {
		return nil, err
	}
	err = d.WebhooksCollection.FindOne(context.TODO(), bson.D{{"id", webhook.Id}, {"tenantId", tenantId}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteWebhook
// Deletes a webhook from the database.
func (d *database) DeleteWebhook(id string, tenantId string) error {
	result, err := d.WebhooksCollection.DeleteOne(context.TODO(), bson.D{{"id", id}, {"tenantId", tenantId}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return utils.ErrWebhookNotFound
	}
	return nil
}

// ListWebhooks
// Returns all webhooks from the database.
func (d *database) ListWebhooks(limit int64, tenantId string) ([]webhookmodel.WebhookResponse, error) {
	var webhooks []webhookmodel.WebhookResponse
	filter := bson.D{{"tenantId", tenantId}}
	opts := options.Find().SetLimit(limit)
	cursor, err := d.WebhooksCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var webhook webhookmodel.WebhookResponse
		err := cursor.Decode(&webhook)
		if err != nil {
			return nil, err
		}
		webhooks = append(webhooks, webhook)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(context.TODO())
	if len(webhooks) == 0 {
		return nil, utils.ErrWebhooksNotFound
	}
	return webhooks, nil
}

// GetWebhook
// Returns a single webhook from the database.
func (d *database) GetWebhook(id string, tenantId string) (*webhookmodel.WebhookResponse, error) {
	var result webhookmodel.WebhookResponse
	err := d.WebhooksCollection.FindOne(context.TODO(), bson.D{{"id", id}, {"tenantId", tenantId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrWebhookNotFound
	} else if err != nil {
		return nil, err
	}
	return &result, err
}

// RotateWebhookSecret
// Updates the secret of an existing webhook in the database.
func (d *database) RotateWebhookSecret(id string, tenantId string, secret string) (*webhookmodel.WebhookResponse, error) {
	var (
		oldResult webhookmodel.WebhookResponse
		newResult webhookmodel.WebhookResponse
	)
	filter := bson.D{{"id", id}, {"tenantId", tenantId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "secret", Value: secret},
		}},
	}
	err := d.WebhooksCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrWebhookNotFound
	} else if err != nil {
		return nil, err
	}
	err = d.WebhooksCollection.FindOne(context.TODO(), bson.D{{"id", oldResult.Id}, {"tenantId", tenantId}}).Decode(&newResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrWebhookNotFound
	} else if err != nil {
		return nil, err
	}
	return &newResult, err
}

// UpdateWebhook
// Updates a webhook in the database.
func (d *database) UpdateWebhook(id string, tenantId string, webhook webhookmodel.UpdateWebhookRequest) (*webhookmodel.WebhookResponse, error) {
	var (
		oldResult webhookmodel.WebhookResponse
		newResult webhookmodel.WebhookResponse
	)
	filter := bson.D{{"id", id}, {"tenantId", tenantId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "description", Value: webhook.Description},
			primitive.E{Key: "events", Value: webhook.Events},
			primitive.E{Key: "url", Value: webhook.Url},
		}},
	}
	err := d.WebhooksCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrWebhookNotFound
	} else if err != nil {
		return nil, err
	}
	err = d.WebhooksCollection.FindOne(context.TODO(), bson.D{{"id", oldResult.Id}, {"tenantId", tenantId}}).Decode(&newResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrWebhookNotFound
	} else if err != nil {
		return nil, err
	}
	return &newResult, nil
}
