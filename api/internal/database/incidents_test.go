package database

import (
	"testing"

	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/incidentmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateIncident(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.IncidentCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse(),
			mtest.CreateCursorResponse(1, "create.incident", mtest.FirstBatch, bson.D{
				{Key: "id", Value: mocks.Incident().Id},
				{Key: "assignedTo", Value: mocks.Incident().AssignedTo},
				{Key: "attachments", Value: mocks.Incident().Attachments},
				{Key: "createdBy", Value: mocks.Incident().CreatedBy},
				{Key: "description", Value: mocks.Incident().Description},
				{Key: "severity", Value: mocks.Incident().Severity},
				{Key: "status", Value: mocks.Incident().Status},
				{Key: "tags", Value: mocks.Incident().Tags},
				{Key: "tasks", Value: mocks.Incident().Tasks},
				{Key: "title", Value: mocks.Incident().Title},
				{Key: "tlp", Value: mocks.Incident().Tlp},
			}))
		createdIncident, err := Client.CreateIncident("", *mocks.Incident())
		assert.Nil(t, err)
		assert.Equal(t, *mocks.IncidentResponse(), *createdIncident)
	})
}

func TestDeleteIncident(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.IncidentCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {"n", 1}})
		err := Client.DeleteIncident(mocks.Incident().Id, "")
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {
		Client.IncidentCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {"n", 0}})
		err := Client.DeleteIncident(mocks.Incident().Id, "")
		assert.NotNil(t, err)
	})
}

func TestGetIncident(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.IncidentCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "get.incident", mtest.FirstBatch, bson.D{
			{Key: "id", Value: mocks.Incident().Id},
			{Key: "assignedTo", Value: mocks.Incident().AssignedTo},
			{Key: "attachments", Value: mocks.Incident().Attachments},
			{Key: "createdBy", Value: mocks.Incident().CreatedBy},
			{Key: "description", Value: mocks.Incident().Description},
			{Key: "severity", Value: mocks.Incident().Severity},
			{Key: "status", Value: mocks.Incident().Status},
			{Key: "tags", Value: mocks.Incident().Tags},
			{Key: "tasks", Value: mocks.Incident().Tasks},
			{Key: "title", Value: mocks.Incident().Title},
			{Key: "tlp", Value: mocks.Incident().Tlp},
		}))
		incident, err := Client.GetIncident(mocks.Incident().Id, "")
		assert.Nil(t, err)
		assert.Equal(t, mocks.IncidentResponse(), incident)
	})
}

func TestListIncidents(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.IncidentCollection = mt.Coll
		first := mtest.CreateCursorResponse(1, "list.incident", mtest.FirstBatch, bson.D{
			{Key: "id", Value: mocks.Incident().Id},
			{Key: "assignedTo", Value: mocks.Incident().AssignedTo},
			{Key: "attachments", Value: mocks.Incident().Attachments},
			{Key: "createdBy", Value: mocks.Incident().CreatedBy},
			{Key: "description", Value: mocks.Incident().Description},
			{Key: "severity", Value: mocks.Incident().Severity},
			{Key: "status", Value: mocks.Incident().Status},
			{Key: "tags", Value: mocks.Incident().Tags},
			{Key: "tasks", Value: mocks.Incident().Tasks},
			{Key: "title", Value: mocks.Incident().Title},
			{Key: "tlp", Value: mocks.Incident().Tlp},
		})
		second := mtest.CreateCursorResponse(1, "list.incident", mtest.NextBatch, bson.D{
			{Key: "id", Value: mocks.Incident2().Id},
			{Key: "assignedTo", Value: mocks.Incident2().AssignedTo},
			{Key: "attachments", Value: mocks.Incident2().Attachments},
			{Key: "createdBy", Value: mocks.Incident2().CreatedBy},
			{Key: "description", Value: mocks.Incident2().Description},
			{Key: "severity", Value: mocks.Incident2().Severity},
			{Key: "status", Value: mocks.Incident2().Status},
			{Key: "tags", Value: mocks.Incident2().Tags},
			{Key: "tasks", Value: mocks.Incident2().Tasks},
			{Key: "title", Value: mocks.Incident2().Title},
			{Key: "tlp", Value: mocks.Incident2().Tlp},
		})
		killCursors := mtest.CreateCursorResponse(0, "list.incident", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		incidents, err := Client.ListIncidents(2, "")
		assert.Nil(t, err)
		assert.Equal(t, *mocks.Incidents(), incidents)
	})
}

func TestUpdateIncident(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.IncidentCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "assignedTo", Value: mocks.Incident().AssignedTo},
				{Key: "description", Value: mocks.Incident().Description},
				{Key: "severity", Value: mocks.Incident().Severity},
				{Key: "status", Value: mocks.Incident().Status},
				{Key: "tags", Value: mocks.Incident().Tags},
				{Key: "title", Value: mocks.Incident().Title},
				{Key: "tlp", Value: mocks.Incident().Tlp},
			}},
		}, mtest.CreateCursorResponse(1, "update.incident", mtest.FirstBatch, bson.D{
			{Key: "id", Value: mocks.Incident().Id},
			{Key: "assignedTo", Value: mocks.Incident().AssignedTo},
			{Key: "attachments", Value: mocks.Incident().Attachments},
			{Key: "createdBy", Value: mocks.Incident().CreatedBy},
			{Key: "description", Value: mocks.Incident().Description},
			{Key: "severity", Value: mocks.Incident().Severity},
			{Key: "status", Value: mocks.Incident().Status},
			{Key: "tags", Value: mocks.Incident().Tags},
			{Key: "tasks", Value: mocks.Incident().Tasks},
			{Key: "title", Value: mocks.Incident().Title},
			{Key: "tlp", Value: mocks.Incident().Tlp},
		}))
		updatedIncident, err := Client.UpdateIncident(mocks.Incident().Id, "", incidentmodel.UpdateIncidentRequest{
			AssignedTo:  mocks.Incident().AssignedTo,
			Description: mocks.Incident().Description,
			Severity:    mocks.Incident().Severity,
			Status:      mocks.Incident().Status,
			Tags:        mocks.Incident().Tags,
			Title:       mocks.Incident().Title,
			Tlp:         mocks.Incident().Tlp,
		})
		assert.Nil(t, err)
		assert.Equal(t, mocks.IncidentResponse(), updatedIncident)
	})
}
