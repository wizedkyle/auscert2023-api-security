package database

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/usermodel"
	"testing"

	"github.com/ForgeResponse/ForgeResponse/v2/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.UserCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse(),
			mtest.CreateCursorResponse(1, "create.user", mtest.FirstBatch, bson.D{
				{"id", mocks.StandardUser().Id},
				{"tenantId", mocks.StandardUser().TenantId},
				{"email", mocks.StandardUser().Email},
				{"firstName", mocks.StandardUser().FirstName},
				{"lastName", mocks.StandardUser().LastName},
				{"lastSignIn", mocks.StandardUser().LastSignIn},
				{"createdTime", mocks.StandardUser().CreatedTime},
				{"roles", mocks.StandardUser().Roles},
				{"isActive", mocks.StandardUser().IsActive},
			}))
		createdUser, err := Client.CreateUser("", *mocks.StandardUser())
		assert.Nil(t, err)
		assert.Equal(t, *mocks.StandardUser(), *createdUser)
	})

	mt.Run("duplicate email error", func(mt *mtest.T) {
		Client.UserCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))
		createdUser, err := Client.CreateUser("", *mocks.StandardUser())
		assert.Nil(t, createdUser)
		assert.NotNil(t, err)
		assert.Equal(t, utils.ErrEmailAlreadyExists, err)
	})
}

func TestDeleteUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.UserCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		err := Client.DeleteUser(mocks.StandardUser().Id, "")
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {
		Client.UserCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		err := Client.DeleteUser(mocks.StandardUser().Id, "")
		assert.NotNil(t, err)
	})
}

func TestGetUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.UserCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "get.user", mtest.FirstBatch, bson.D{
			{"id", mocks.StandardUser().Id},
			{"email", mocks.StandardUser().Email},
			{"firstName", mocks.StandardUser().FirstName},
			{"lastName", mocks.StandardUser().LastName},
			{"lastSignIn", mocks.StandardUser().LastSignIn},
			{"createdTime", mocks.StandardUser().CreatedTime},
			{"roles", mocks.StandardUser().Roles},
			{"isActive", mocks.StandardUser().IsActive},
		}))
		user, err := Client.GetUser(mocks.StandardUser().Id, "")
		assert.Nil(t, err)
		assert.Equal(t, mocks.UserResponse(), user)
	})
}

func TestListUsers(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.UserCollection = mt.Coll
		first := mtest.CreateCursorResponse(1, "list.users", mtest.FirstBatch, bson.D{
			{"id", mocks.StandardUser().Id},
			{"email", mocks.StandardUser().Email},
			{"firstName", mocks.StandardUser().FirstName},
			{"lastName", mocks.StandardUser().LastName},
			{"lastSignIn", mocks.StandardUser().LastSignIn},
			{"createdTime", mocks.StandardUser().CreatedTime},
			{"roles", mocks.StandardUser().Roles},
			{"isActive", mocks.StandardUser().IsActive},
		})
		second := mtest.CreateCursorResponse(2, "list.users", mtest.NextBatch, bson.D{
			{"id", mocks.StandardUser2().Id},
			{"email", mocks.StandardUser2().Email},
			{"firstName", mocks.StandardUser2().FirstName},
			{"lastName", mocks.StandardUser2().LastName},
			{"lastSignIn", mocks.StandardUser2().LastSignIn},
			{"createdTime", mocks.StandardUser2().CreatedTime},
			{"roles", mocks.StandardUser2().Roles},
			{"isActive", mocks.StandardUser2().IsActive},
		})
		killCursors := mtest.CreateCursorResponse(0, "list.users", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		users, err := Client.ListUsers(2, "")
		assert.Nil(t, err)
		assert.Equal(t, *mocks.StandardUsers(), users)
	})
}

func TestUpdateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.UserCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", bson.D{
				{"id", mocks.StandardUser().Id},
				{"email", mocks.StandardUser().Email},
				{"firstName", mocks.StandardUser().FirstName},
				{"lastName", mocks.StandardUser().LastName},
				{"lastSignIn", mocks.StandardUser().LastSignIn},
				{"createdTime", mocks.StandardUser().CreatedTime},
				{"roles", mocks.StandardUser().Roles},
				{"isActive", mocks.StandardUser().IsActive},
			}},
		}, mtest.CreateCursorResponse(1, "update.user", mtest.FirstBatch, bson.D{
			{"id", mocks.StandardUser().Id},
			{"email", mocks.StandardUser().Email},
			{"firstName", mocks.StandardUser().FirstName},
			{"lastName", mocks.StandardUser().LastName},
			{"lastSignIn", mocks.StandardUser().LastSignIn},
			{"createdTime", mocks.StandardUser().CreatedTime},
			{"roles", mocks.StandardUser().Roles},
			{"isActive", mocks.StandardUser().IsActive},
		}))
		updatedUser, err := Client.UpdateUser(mocks.StandardUser().Id, "", usermodel.UpdateUserRequest{
			FirstName: mocks.StandardUser().FirstName,
			LastName:  mocks.StandardUser().LastName,
			IsActive:  &mocks.StandardUser().IsActive,
		})
		assert.Nil(t, err)
		assert.Equal(t, mocks.UserResponse(), updatedUser)
	})
}
