package database

import (
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/accesskeymodel"
	"github.com/ForgeResponse/ForgeResponse/v2/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestCreateAccessKey(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.AccessKeysCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse(),
			mtest.CreateCursorResponse(1, "create.accesskey", mtest.FirstBatch, bson.D{
				{"id", mocks.AccessKey().Id},
				{"description", mocks.AccessKey().Description},
				{"expiration", mocks.AccessKey().Expiration},
				{"keyHash", mocks.AccessKey().KeyHash},
				{"keyPrefix", mocks.AccessKey().KeyPrefix},
				{"tenantId", mocks.AccessKey().TenantId},
				{"scopes", mocks.AccessKey().Scopes},
				{"version", mocks.AccessKey().Version},
			}))
		createdAccessKey, err := Client.CreateAccessKey("", *mocks.AccessKey())
		assert.Nil(t, err)
		assert.Equal(t, *mocks.AccessKey(), *createdAccessKey)
	})
}

func TestDeleteAccessKey(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.AccessKeysCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		err := Client.DeleteAccessKey(mocks.AccessKey().Id, "")
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {
		Client.AccessKeysCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		err := Client.DeleteAccessKey(mocks.AccessKey().Id, "")
		assert.NotNil(t, err)
	})
}

func TestGetAccessKey(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.AccessKeysCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "get.accesskey", mtest.FirstBatch, bson.D{
			{"id", mocks.AccessKey().Id},
			{"description", mocks.AccessKey().Description},
			{"expiration", mocks.AccessKey().Expiration},
			{"keyPrefix", mocks.AccessKey().KeyPrefix},
			{"scopes", mocks.AccessKey().Scopes},
		}))
		accessKey, err := Client.GetAccessKey(mocks.AccessKey().Id, "")
		assert.Nil(t, err)
		assert.Equal(t, mocks.AccessKeyResponse(), accessKey)
	})
}

func TestGetAccessKeyByHash(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.AccessKeysCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "get.accesskeybyhash", mtest.FirstBatch, bson.D{
			{"id", mocks.AccessKey().Id},
			{"description", mocks.AccessKey().Description},
			{"expiration", mocks.AccessKey().Expiration},
			{"keyHash", mocks.AccessKey().KeyHash},
			{"keyPrefix", mocks.AccessKey().KeyPrefix},
			{"tenantId", mocks.AccessKey().TenantId},
			{"scopes", mocks.AccessKey().Scopes},
			{"version", mocks.AccessKey().Version},
		}))
		accessKey, err := Client.GetAccessKeyByHash(mocks.AccessKey().KeyHash)
		assert.Nil(t, err)
		assert.Equal(t, mocks.AccessKey(), accessKey)
	})
}

func TestListAccessKeys(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.AccessKeysCollection = mt.Coll
		first := mtest.CreateCursorResponse(1, "list.accesskeys", mtest.FirstBatch, bson.D{
			{"id", mocks.AccessKey().Id},
			{"description", mocks.AccessKey().Description},
			{"expiration", mocks.AccessKey().Expiration},
			{"keyPrefix", mocks.AccessKey().KeyPrefix},
			{"scopes", mocks.AccessKey().Scopes},
		})
		second := mtest.CreateCursorResponse(1, "list.accesskeys", mtest.NextBatch, bson.D{
			{"id", mocks.AccessKey2().Id},
			{"description", mocks.AccessKey2().Description},
			{"expiration", mocks.AccessKey2().Expiration},
			{"keyPrefix", mocks.AccessKey2().KeyPrefix},
			{"scopes", mocks.AccessKey2().Scopes},
		})
		killCursors := mtest.CreateCursorResponse(0, "list.accesskeys", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		accessKeys, err := Client.ListAccessKeys(2, "")
		assert.Nil(t, err)
		assert.Equal(t, *mocks.AccessKeys(), accessKeys)
	})
}

func TestRotateAccessKey(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.AccessKeysCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", bson.D{
				{"expiration", mocks.AccessKey().Expiration},
				{"keyHash", mocks.AccessKey().KeyHash},
				{"keyPrefix", mocks.AccessKey().KeyPrefix},
			}},
		}, mtest.CreateCursorResponse(1, "rotate.accesskey", mtest.FirstBatch, bson.D{
			{"id", mocks.AccessKey().Id},
			{"description", mocks.AccessKey().Description},
			{"expiration", mocks.AccessKey().Expiration},
			{"scopes", mocks.AccessKey().Scopes},
		}))
		updatedAccessKey, err := Client.RotateAccessKey(mocks.AccessKey().Id, "", accesskeymodel.AccessKey{
			Expiration: mocks.AccessKey().Expiration,
			KeyPrefix:  mocks.AccessKey().KeyPrefix,
			KeyHash:    mocks.AccessKey().KeyHash,
		})
		assert.Nil(t, err)
		assert.Equal(t, mocks.UpdateAccessKeyResponse(), updatedAccessKey)
	})
}

func TestUpdateAccessKey(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.AccessKeysCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", bson.D{
				{"description", mocks.AccessKey().Description},
				{"scopes", mocks.AccessKey().Scopes},
			}},
		}, mtest.CreateCursorResponse(1, "update.accesskey", mtest.FirstBatch, bson.D{
			{"id", mocks.AccessKey().Id},
			{"description", mocks.AccessKey().Description},
			{"expiration", mocks.AccessKey().Expiration},
			{"scopes", mocks.AccessKey().Scopes},
		}))
		updatedAccessKey, err := Client.UpdateAccessKey(mocks.AccessKey().Id, "", accesskeymodel.UpdateAccessKeyRequest{
			Description: mocks.AccessKey().Description,
			Scopes:      mocks.AccessKey().Scopes,
		})
		assert.Nil(t, err)
		assert.Equal(t, mocks.UpdateAccessKeyResponse(), updatedAccessKey)
	})
}
