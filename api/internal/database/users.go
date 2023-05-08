package database

import (
	"context"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/usermodel"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser
// Creates a new user in the database.
func (d *database) CreateUser(tenantId string, user usermodel.User) (*usermodel.User, error) {
	var result usermodel.User
	user.Id = uuid.NewString()
	_, err := d.UserCollection.InsertOne(context.TODO(), user)
	if mongo.IsDuplicateKeyError(err) {
		return nil, utils.ErrEmailAlreadyExists
	}
	if err != nil {
		return nil, err
	}
	err = d.UserCollection.FindOne(context.TODO(), bson.D{{"id", user.Id}, {"tenantId", tenantId}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteUser
// Deletes a user from the database.
func (d *database) DeleteUser(id string, tenantId string) error {
	result, err := d.UserCollection.DeleteOne(context.TODO(), bson.D{{"id", id}, {"tenantId", tenantId}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return utils.ErrUserNotFound
	}
	return err
}

// ListUsers
// Returns all users from the database.
func (d *database) ListUsers(limit int64, tenantId string) ([]usermodel.UserResponse, error) {
	var users []usermodel.UserResponse
	filter := bson.D{{"tenantId", tenantId}}
	opts := options.Find().SetLimit(limit)
	cursor, err := d.UserCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user usermodel.UserResponse
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(context.TODO())
	if len(users) == 0 {
		return nil, utils.ErrUsersNotFound
	}
	return users, nil
}

// GetUser
// Returns a single user from the database.
func (d *database) GetUser(id string, tenantId string) (*usermodel.UserResponse, error) {
	var result usermodel.UserResponse
	err := d.UserCollection.FindOne(context.TODO(), bson.D{{"id", id}, {"tenantId", tenantId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateUser
// Updates a user in the database.
func (d *database) UpdateUser(id string, tenantId string, user usermodel.UpdateUserRequest) (*usermodel.UserResponse, error) {
	var (
		oldResult usermodel.UserResponse
		newResult usermodel.UserResponse
	)
	filter := bson.D{{"id", id}, {"tenantId", tenantId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "firstName", Value: user.FirstName},
			primitive.E{Key: "lastName", Value: user.LastName},
			primitive.E{Key: "isActive", Value: user.IsActive},
		}},
	}
	err := d.UserCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	err = d.UserCollection.FindOne(context.TODO(), bson.D{{"id", oldResult.Id}, {"tenantId", tenantId}}).Decode(&newResult)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return &newResult, nil
}
