package database

import (
	"context"

	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/accesskeymodel"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateAccessKey
// Creates a new access key in the database.
func (d *database) CreateAccessKey(tenantId string, accessKey accesskeymodel.AccessKey) (*accesskeymodel.AccessKey, error) {
	var result accesskeymodel.AccessKey
	accessKey.Id = uuid.NewString()
	_, err := d.AccessKeysCollection.InsertOne(context.TODO(), accessKey)
	if err != nil {
		return nil, err
	}
	err = d.AccessKeysCollection.FindOne(context.TODO(), bson.D{{Key: "id", Value: accessKey.Id}, {Key: "tenantId", Value: tenantId}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAccessKey
// Deletes an access key from the database.
func (d *database) DeleteAccessKey(id string, tenantId string) error {
	result, err := d.AccessKeysCollection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}, {Key: "tenantId", Value: tenantId}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return utils.ErrAccessKeyNotFound
	}
	return nil
}

// ListAccessKeys
// Returns all access keys from the database.
func (d *database) ListAccessKeys(limit int64, tenantId string) ([]accesskeymodel.AccessKeyResponse, error) {
	var accessKeys []accesskeymodel.AccessKeyResponse
	filter := bson.D{{Key: "tenantId", Value: tenantId}}
	opts := options.Find().SetLimit(limit)
	cursor, err := d.AccessKeysCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var accessKey accesskeymodel.AccessKeyResponse
		err := cursor.Decode(&accessKey)
		if err != nil {
			return nil, err
		}
		accessKeys = append(accessKeys, accessKey)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(context.TODO())
	if len(accessKeys) == 0 {
		return nil, utils.ErrAccessKeysNotFound
	}
	return accessKeys, err
}

// GetAccessKey
// Returns a single access key from the database.
func (d *database) GetAccessKey(id string, tenantId string) (*accesskeymodel.AccessKeyResponse, error) {
	var result accesskeymodel.AccessKeyResponse
	err := d.AccessKeysCollection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}, {Key: "tenantId", Value: tenantId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrAccessKeyNotFound
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAccessKeyByHash
// Returns a single access key from the database based on the provided hash.
func (d *database) GetAccessKeyByHash(hash string) (*accesskeymodel.AccessKey, error) {
	var result accesskeymodel.AccessKey
	err := d.AccessKeysCollection.FindOne(context.TODO(), bson.D{{Key: "keyHash", Value: hash}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrAccessKeyNotFound
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// RotateAccessKey
// Updates the keyHash and expiration of an existing access key in the database.
func (d *database) RotateAccessKey(id string, tenantId string, accessKey accesskeymodel.AccessKey) (*accesskeymodel.UpdateAccessKeyResponse, error) {
	var (
		oldResult accesskeymodel.UpdateAccessKeyResponse
		newResult accesskeymodel.UpdateAccessKeyResponse
	)
	filter := bson.D{{Key: "id", Value: id}, {Key: "tenantId", Value: tenantId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "expiration", Value: accessKey.Expiration},
			primitive.E{Key: "keyHash", Value: accessKey.KeyHash},
			primitive.E{Key: "keyPrefix", Value: accessKey.KeyPrefix},
		}},
	}
	err := d.AccessKeysCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrAccessKeyNotFound
	} else if err != nil {
		return nil, err
	}
	err = d.AccessKeysCollection.FindOne(context.TODO(), bson.D{{Key: "id", Value: oldResult.Id}, {Key: "tenantId", Value: tenantId}}).Decode(&newResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrAccessKeyNotFound
	} else if err != nil {
		return nil, err
	}
	return &newResult, nil
}

// UpdateAccessKey
// Updates an access key in the database.
func (d *database) UpdateAccessKey(id string, tenantId string, accessKey accesskeymodel.UpdateAccessKeyRequest) (*accesskeymodel.UpdateAccessKeyResponse, error) {
	var (
		oldResult accesskeymodel.UpdateAccessKeyResponse
		newResult accesskeymodel.UpdateAccessKeyResponse
	)
	filter := bson.D{{Key: "id", Value: id}, {Key: "tenantId", Value: tenantId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "description", Value: accessKey.Description},
			primitive.E{Key: "scopes", Value: accessKey.Scopes},
		}},
	}
	err := d.AccessKeysCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrAccessKeyNotFound
	} else if err != nil {
		return nil, err
	}
	err = d.AccessKeysCollection.FindOne(context.TODO(), bson.D{{Key: "id", Value: oldResult.Id}, {Key: "tenantId", Value: tenantId}}).Decode(&newResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrAccessKeyNotFound
	} else if err != nil {
		return nil, err
	}
	return &newResult, nil
}
