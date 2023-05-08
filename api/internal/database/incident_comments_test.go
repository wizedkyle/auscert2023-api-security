package database

import (
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/incidentcommentmodel"
	"testing"

	"github.com/ForgeResponse/ForgeResponse/v2/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestDeleteIncidentComment(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.IncidentCommentsCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}})
		err := Client.DeleteIncidentComment(mocks.IncidentComment().Id, mocks.IncidentComment().IncidentId, "")
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {
		Client.IncidentCommentsCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 0}})
		err := Client.DeleteIncidentComment(mocks.IncidentComment().Id, mocks.IncidentComment().IncidentId, "")
		assert.NotNil(t, err)
	})
}

func TestGetIncidentComment(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.IncidentCommentsCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "get.incidentcomment", mtest.FirstBatch, bson.D{
			{Key: "id", Value: mocks.IncidentComment().Id},
			{Key: "incidentId", Value: mocks.IncidentComment().IncidentId},
			{Key: "comment", Value: mocks.IncidentComment().Comment},
			{Key: "createdBy", Value: mocks.IncidentComment().CreatedBy},
			{Key: "createdAt", Value: mocks.IncidentComment().CreatedAt},
			{Key: "order", Value: mocks.IncidentComment().Order},
		}))
		incidentComment, err := Client.GetIncidentComment(mocks.IncidentComment().Id, mocks.IncidentComment().IncidentId, "")
		assert.Nil(t, err)
		assert.Equal(t, mocks.IncidentCommentResponse(), incidentComment)
	})
}

func TestListIncidentComments(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.IncidentCommentsCollection = mt.Coll
		first := mtest.CreateCursorResponse(1, "list.incidentcomment", mtest.FirstBatch, bson.D{
			{Key: "id", Value: mocks.IncidentComment().Id},
			{Key: "incidentId", Value: mocks.IncidentComment().IncidentId},
			{Key: "comment", Value: mocks.IncidentComment().Comment},
			{Key: "createdBy", Value: mocks.IncidentComment().CreatedBy},
			{Key: "createdAt", Value: mocks.IncidentComment().CreatedAt},
			{Key: "order", Value: mocks.IncidentComment().Order},
		})
		second := mtest.CreateCursorResponse(1, "list.incidentcomment", mtest.NextBatch, bson.D{
			{Key: "id", Value: mocks.IncidentComment2().Id},
			{Key: "incidentId", Value: mocks.IncidentComment2().IncidentId},
			{Key: "comment", Value: mocks.IncidentComment2().Comment},
			{Key: "createdBy", Value: mocks.IncidentComment2().CreatedBy},
			{Key: "createdAt", Value: mocks.IncidentComment2().CreatedAt},
			{Key: "order", Value: mocks.IncidentComment2().Order},
		})
		killCursors := mtest.CreateCursorResponse(0, "list.incidentcomment", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		incidentComments, err := Client.ListIncidentComments(2, mocks.IncidentComment().IncidentId, "")
		assert.Nil(t, err)
		assert.Equal(t, *mocks.IncidentComments(), incidentComments)
	})
}

func TestUpdateIncidentComment(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.IncidentCommentsCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "comment", Value: mocks.IncidentComment().Comment},
			}},
		}, mtest.CreateCursorResponse(1, "update.incidentcomment", mtest.FirstBatch, bson.D{
			{Key: "id", Value: mocks.IncidentComment().Id},
			{Key: "incidentId", Value: mocks.IncidentComment().IncidentId},
			{Key: "comment", Value: mocks.IncidentComment().Comment},
			{Key: "createdBy", Value: mocks.IncidentComment().CreatedBy},
			{Key: "createdAt", Value: mocks.IncidentComment().CreatedAt},
			{Key: "order", Value: mocks.IncidentComment().Order},
		}))
		updatedIncidentComment, err := Client.UpdateIncidentComment(mocks.IncidentComment().Id, mocks.IncidentComment().IncidentId, "", incidentcommentmodel.UpdateIncidentComment{
			Comment: mocks.IncidentComment().Comment,
		})
		assert.Nil(t, err)
		assert.Equal(t, mocks.IncidentCommentResponse(), updatedIncidentComment)
	})
}
