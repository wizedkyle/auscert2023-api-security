package database

import (
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/webhookmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestCreateWebhook(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.WebhooksCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse(),
			mtest.CreateCursorResponse(1, "create.webhook", mtest.FirstBatch, bson.D{
				{"id", mocks.Webhook().Id},
				{"algorithm", mocks.Webhook().Algorithm},
				{"description", mocks.Webhook().Description},
				{"events", mocks.Webhook().Events},
				{"secret", mocks.Webhook().Secret},
				{"tenantId", mocks.Webhook().TenantId},
				{"url", mocks.Webhook().Url},
				{"version", mocks.Webhook().Version},
			}))
	})
	createdWebhook, err := Client.CreateWebhook("", *mocks.Webhook())
	assert.Nil(t, err)
	assert.Equal(t, *mocks.WebhookResponse(), *createdWebhook)
}

func TestDeleteWebhook(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.WebhooksCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		err := Client.DeleteWebhook(mocks.Webhook().Id, "")
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {
		Client.WebhooksCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		err := Client.DeleteWebhook(mocks.Webhook().Id, "")
		assert.NotNil(t, err)
	})
}

func TestGetWebhook(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.WebhooksCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "get.webhook", mtest.FirstBatch, bson.D{
			{"id", mocks.Webhook().Id},
			{"algorithm", mocks.Webhook().Algorithm},
			{"description", mocks.Webhook().Description},
			{"events", mocks.Webhook().Events},
			{"secret", mocks.Webhook().Secret},
			{"url", mocks.Webhook().Url},
		}))
		webhook, err := Client.GetWebhook(mocks.Webhook().Id, "")
		assert.Nil(t, err)
		assert.Equal(t, mocks.WebhookResponse(), webhook)
	})
}

func TestListWebhooks(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.WebhooksCollection = mt.Coll
		first := mtest.CreateCursorResponse(1, "list.webhooks", mtest.FirstBatch, bson.D{
			{"id", mocks.Webhook().Id},
			{"algorithm", mocks.Webhook().Algorithm},
			{"description", mocks.Webhook().Description},
			{"events", mocks.Webhook().Events},
			{"secret", mocks.Webhook().Secret},
			{"url", mocks.Webhook().Url},
		})
		second := mtest.CreateCursorResponse(1, "list.webhooks", mtest.NextBatch, bson.D{
			{"id", mocks.Webhook2().Id},
			{"algorithm", mocks.Webhook2().Algorithm},
			{"description", mocks.Webhook2().Description},
			{"events", mocks.Webhook2().Events},
			{"secret", mocks.Webhook2().Secret},
			{"url", mocks.Webhook2().Url},
		})
		killCursors := mtest.CreateCursorResponse(0, "list.webhooks", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		webhooks, err := Client.ListWebhooks(2, "")
		assert.Nil(t, err)
		assert.Equal(t, *mocks.Webhooks(), webhooks)
	})
}

func TestRotateWebhookSecret(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.WebhooksCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", bson.D{
				{"secret", mocks.Webhook().Secret},
			}},
		}, mtest.CreateCursorResponse(1, "rotate.webhook", mtest.FirstBatch, bson.D{
			{"id", mocks.Webhook().Id},
			{"algorithm", mocks.Webhook().Algorithm},
			{"description", mocks.Webhook().Description},
			{"events", mocks.Webhook().Events},
			{"secret", mocks.Webhook().Secret},
			{"url", mocks.Webhook().Url},
		}))
		rotatedWebhook, err := Client.RotateWebhookSecret(mocks.Webhook().Id, "", mocks.Webhook().Secret)
		assert.Nil(t, err)
		assert.Equal(t, mocks.WebhookResponse(), rotatedWebhook)
	})
}

func TestUpdateWebhook(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.WebhooksCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", bson.D{
				{"description", mocks.Webhook().Description},
				{"events", mocks.Webhook().Events},
				{"url", mocks.Webhook().Url},
			}},
		}, mtest.CreateCursorResponse(1, "update.webhook", mtest.FirstBatch, bson.D{
			{"id", mocks.Webhook().Id},
			{"algorithm", mocks.Webhook().Algorithm},
			{"description", mocks.Webhook().Description},
			{"events", mocks.Webhook().Events},
			{"secret", mocks.Webhook().Secret},
			{"url", mocks.Webhook().Url},
		}))
		updatedWebhook, err := Client.UpdateWebhook(mocks.Webhook().Id, "", webhookmodel.UpdateWebhookRequest{
			Description: mocks.Webhook().Description,
			Events:      mocks.Webhook().Events,
			Url:         mocks.Webhook().Url,
		})
		assert.Nil(t, err)
		assert.Equal(t, mocks.WebhookResponse(), updatedWebhook)
	})
}
