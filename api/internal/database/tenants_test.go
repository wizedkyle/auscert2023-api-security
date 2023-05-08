package database

import (
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/tenantmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestCreateTenant(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.TenantCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse(),
			mtest.CreateCursorResponse(1, "create.tenant", mtest.FirstBatch, bson.D{
				{"tenantId", mocks.Tenant().TenantId},
				{"name", mocks.Tenant().Name},
			}))
		createdTenant, err := Client.CreateTenant(*mocks.Tenant())
		assert.Nil(t, err)
		assert.Equal(t, *mocks.Tenant(), *createdTenant)
	})
}

func TestDeleteTenant(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.TenantCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		err := Client.DeleteTenant(mocks.Tenant().TenantId)
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {
		Client.TenantCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		err := Client.DeleteTenant(mocks.Tenant().TenantId)
		assert.NotNil(t, err)
	})
}

func TestGetTenant(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.TenantCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "get.tenant", mtest.FirstBatch, bson.D{
			{"tenantId", mocks.Tenant().TenantId},
			{"name", mocks.Tenant().Name},
		}))
		tenant, err := Client.GetTenant(mocks.Tenant().TenantId)
		assert.Nil(t, err)
		assert.Equal(t, mocks.Tenant(), tenant)
	})
}

func TestUpdateTenant(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.TenantCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", bson.D{
				{"tenantId", mocks.Tenant().TenantId},
				{"name", mocks.Tenant().Name},
			}},
		}, mtest.CreateCursorResponse(1, "update.tenant", mtest.FirstBatch, bson.D{
			{"tenantId", mocks.Tenant().TenantId},
			{"name", mocks.Tenant().Name},
		}))
		updatedTenant, err := Client.UpdateTenant(mocks.Tenant().TenantId, tenantmodel.UpdateTenantRequest{
			Name: mocks.Tenant().Name,
		})
		assert.Nil(t, err)
		assert.Equal(t, mocks.Tenant(), updatedTenant)
	})
}
