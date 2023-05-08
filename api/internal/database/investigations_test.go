package database

import (
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/investigationmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestCreateInvestigation(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.InvestigationCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse(),
			mtest.CreateCursorResponse(1, "create.investigation", mtest.FirstBatch, bson.D{
				{"id", mocks.Investigation().Id},
				{"assignedTo", mocks.Investigation().AssignedTo},
				{"attachments", mocks.Investigation().Attachments},
				{"createdBy", mocks.Investigation().CreatedBy},
				{"createdAt", mocks.Investigation().CreatedAt},
				{"comments", mocks.Investigation().Comments},
				{"description", mocks.Investigation().Description},
				{"investigationId", mocks.Investigation().InvestigationId},
				{"severity", mocks.Investigation().Severity},
				{"status", mocks.Investigation().Status},
				{"tags", mocks.Investigation().Tags},
				{"title", mocks.Investigation().Title},
				{"tlp", mocks.Investigation().Tlp},
			}))
		createdInvestigation, err := Client.CreateInvestigation("", *mocks.Investigation())
		assert.Nil(t, err)
		assert.Equal(t, *mocks.InvestigationResponse(), *createdInvestigation)
	})
}

func TestDeleteInvestigation(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.InvestigationCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		err := Client.DeleteInvestigation(mocks.Investigation().Id, "")
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {
		Client.InvestigationCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		err := Client.DeleteInvestigation(mocks.Investigation().Id, "")
		assert.NotNil(t, err)
	})
}

func TestGetInvestigation(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.InvestigationCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "get.investigation", mtest.FirstBatch, bson.D{
			{"id", mocks.Investigation().Id},
			{"assignedTo", mocks.Investigation().AssignedTo},
			{"attachments", mocks.Investigation().Attachments},
			{"createdBy", mocks.Investigation().CreatedBy},
			{"createdAt", mocks.Investigation().CreatedAt},
			{"comments", mocks.Investigation().Comments},
			{"description", mocks.Investigation().Description},
			{"investigationId", mocks.Investigation().InvestigationId},
			{"severity", mocks.Investigation().Severity},
			{"status", mocks.Investigation().Status},
			{"tags", mocks.Investigation().Tags},
			{"title", mocks.Investigation().Title},
			{"tlp", mocks.Investigation().Tlp},
		}))
		investigation, err := Client.GetInvestigation(mocks.Investigation().Id, "")
		assert.Nil(t, err)
		assert.Equal(t, mocks.InvestigationResponse(), investigation)
	})
}

func TestListInvestigations(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.InvestigationCollection = mt.Coll
		first := mtest.CreateCursorResponse(1, "list.investigation", mtest.FirstBatch, bson.D{
			{"id", mocks.Investigation().Id},
			{"assignedTo", mocks.Investigation().AssignedTo},
			{"attachments", mocks.Investigation().Attachments},
			{"createdBy", mocks.Investigation().CreatedBy},
			{"createdAt", mocks.Investigation().CreatedAt},
			{"comments", mocks.Investigation().Comments},
			{"description", mocks.Investigation().Description},
			{"investigationId", mocks.Investigation().InvestigationId},
			{"severity", mocks.Investigation().Severity},
			{"status", mocks.Investigation().Status},
			{"tags", mocks.Investigation().Tags},
			{"title", mocks.Investigation().Title},
			{"tlp", mocks.Investigation().Tlp},
		})
		second := mtest.CreateCursorResponse(1, "list.investigation", mtest.NextBatch, bson.D{
			{"id", mocks.Investigation2().Id},
			{"assignedTo", mocks.Investigation2().AssignedTo},
			{"attachments", mocks.Investigation2().Attachments},
			{"createdBy", mocks.Investigation2().CreatedBy},
			{"createdAt", mocks.Investigation2().CreatedAt},
			{"comments", mocks.Investigation2().Comments},
			{"description", mocks.Investigation2().Description},
			{"investigationId", mocks.Investigation2().InvestigationId},
			{"severity", mocks.Investigation2().Severity},
			{"status", mocks.Investigation2().Status},
			{"tags", mocks.Investigation2().Tags},
			{"title", mocks.Investigation2().Title},
			{"tlp", mocks.Investigation2().Tlp},
		})
		killCursors := mtest.CreateCursorResponse(0, "list.investigation", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		investigations, err := Client.ListInvestigations(2, "")
		assert.Nil(t, err)
		assert.Equal(t, *mocks.Investigations(), investigations)
	})
}

func TestUpdateInvestigation(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.InvestigationCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", bson.D{
				{"description", mocks.Investigation().Description},
				{"severity", mocks.Investigation().Severity},
				{"status", mocks.Investigation().Status},
				{"tags", mocks.Investigation().Tags},
				{"title", mocks.Investigation().Title},
				{"tlp", mocks.Investigation().Tlp},
			}},
		}, mtest.CreateCursorResponse(1, "update.investigation", mtest.FirstBatch, bson.D{
			{"id", mocks.Investigation().Id},
			{"assignedTo", mocks.Investigation().AssignedTo},
			{"attachments", mocks.Investigation().Attachments},
			{"createdBy", mocks.Investigation().CreatedBy},
			{"createdAt", mocks.Investigation().CreatedAt},
			{"comments", mocks.Investigation().Comments},
			{"description", mocks.Investigation().Description},
			{"investigationId", mocks.Investigation().InvestigationId},
			{"severity", mocks.Investigation().Severity},
			{"status", mocks.Investigation().Status},
			{"tags", mocks.Investigation().Tags},
			{"title", mocks.Investigation().Title},
			{"tlp", mocks.Investigation().Tlp},
		}))
		updatedInvestigation, err := Client.UpdateInvestigation(mocks.Investigation().Id, "", investigationmodel.UpdateInvestigationRequest{
			Description: mocks.Investigation().Description,
			Severity:    mocks.Investigation().Severity,
			Status:      mocks.Investigation().Status,
			Tags:        mocks.Investigation().Tags,
			Title:       mocks.Investigation().Title,
			Tlp:         mocks.Investigation().Tlp,
		})
		assert.Nil(t, err)
		assert.Equal(t, mocks.InvestigationResponse(), updatedInvestigation)
	})
}
