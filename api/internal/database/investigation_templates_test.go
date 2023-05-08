package database

import (
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/investigationtemplatemodel"
	"github.com/ForgeResponse/ForgeResponse/v2/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestCreateInvestigationTemplate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.InvestigationTemplatesCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse(),
			mtest.CreateCursorResponse(1, "create.investigationtemplate", mtest.FirstBatch, bson.D{
				{"id", mocks.InvestigationTemplate().Id},
				{"createdBy", mocks.InvestigationTemplate().CreatedBy},
				{"createdAt", mocks.InvestigationTemplate().CreatedAt},
				{"titlePrefix", mocks.InvestigationTemplate().TitlePrefix},
				{"severity", mocks.InvestigationTemplate().Severity},
				{"status", mocks.InvestigationTemplate().Status},
				{"tags", mocks.InvestigationTemplate().Tags},
				{"tlp", mocks.InvestigationTemplate().Tlp},
			}))
		createdInvestigationTemplate, err := Client.CreateInvestigationTemplate("", *mocks.InvestigationTemplate())
		assert.Nil(t, err)
		assert.Equal(t, *mocks.InvestigationTemplateResponse(), *createdInvestigationTemplate)
	})
}

func TestDeleteInvestigationTemplate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.InvestigationTemplatesCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		err := Client.DeleteInvestigationTemplate(mocks.InvestigationTemplate().Id, "")
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {
		Client.InvestigationTemplatesCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		err := Client.DeleteInvestigationTemplate(mocks.InvestigationTemplate().Id, "")
		assert.NotNil(t, err)
	})
}

func TestGetInvestigationTemplate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.InvestigationTemplatesCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "get.investigationtemplate", mtest.FirstBatch, bson.D{
			{"id", mocks.InvestigationTemplate().Id},
			{"createdBy", mocks.InvestigationTemplate().CreatedBy},
			{"createdAt", mocks.InvestigationTemplate().CreatedAt},
			{"titlePrefix", mocks.InvestigationTemplate().TitlePrefix},
			{"severity", mocks.InvestigationTemplate().Severity},
			{"status", mocks.InvestigationTemplate().Status},
			{"tags", mocks.InvestigationTemplate().Tags},
			{"tlp", mocks.InvestigationTemplate().Tlp},
		}))
		investigationTemplate, err := Client.GetInvestigationTemplate(mocks.InvestigationTemplate().Id, "")
		assert.Nil(t, err)
		assert.Equal(t, mocks.InvestigationTemplateResponse(), investigationTemplate)
	})
}

func TestListInvestigationTemplates(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.InvestigationTemplatesCollection = mt.Coll
		first := mtest.CreateCursorResponse(1, "list.investigationtemplate", mtest.FirstBatch, bson.D{
			{"id", mocks.InvestigationTemplate().Id},
			{"createdBy", mocks.InvestigationTemplate().CreatedBy},
			{"createdAt", mocks.InvestigationTemplate().CreatedAt},
			{"titlePrefix", mocks.InvestigationTemplate().TitlePrefix},
			{"severity", mocks.InvestigationTemplate().Severity},
			{"status", mocks.InvestigationTemplate().Status},
			{"tags", mocks.InvestigationTemplate().Tags},
			{"tlp", mocks.InvestigationTemplate().Tlp},
		})
		second := mtest.CreateCursorResponse(1, "list.investigationtemplate", mtest.NextBatch, bson.D{
			{"id", mocks.InvestigationTemplate2().Id},
			{"createdBy", mocks.InvestigationTemplate2().CreatedBy},
			{"createdAt", mocks.InvestigationTemplate2().CreatedAt},
			{"titlePrefix", mocks.InvestigationTemplate2().TitlePrefix},
			{"severity", mocks.InvestigationTemplate2().Severity},
			{"status", mocks.InvestigationTemplate2().Status},
			{"tags", mocks.InvestigationTemplate2().Tags},
			{"tlp", mocks.InvestigationTemplate2().Tlp},
		})
		killCursors := mtest.CreateCursorResponse(0, "list.investigationtemplate", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		investigationTemplates, err := Client.ListInvestigationTemplates(2, "")
		assert.Nil(t, err)
		assert.Equal(t, *mocks.InvestigationTemplates(), investigationTemplates)
	})
}

func TestUpdateInvestigationTemplate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	Client.Client = mt.Client
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		Client.InvestigationTemplatesCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", bson.D{
				{"titlePrefix", mocks.InvestigationTemplate().TitlePrefix},
				{"severity", mocks.InvestigationTemplate().Severity},
				{"status", mocks.InvestigationTemplate().Status},
				{"tags", mocks.InvestigationTemplate().Tags},
				{"tlp", mocks.InvestigationTemplate().Tlp},
			}},
		}, mtest.CreateCursorResponse(1, "update.investigationtemplate", mtest.FirstBatch, bson.D{
			{"id", mocks.InvestigationTemplate().Id},
			{"createdBy", mocks.InvestigationTemplate().CreatedBy},
			{"createdAt", mocks.InvestigationTemplate().CreatedAt},
			{"titlePrefix", mocks.InvestigationTemplate().TitlePrefix},
			{"severity", mocks.InvestigationTemplate().Severity},
			{"status", mocks.InvestigationTemplate().Status},
			{"tags", mocks.InvestigationTemplate().Tags},
			{"tlp", mocks.InvestigationTemplate().Tlp},
		}))
		updatedInvestigationTemplate, err := Client.UpdateInvestigationTemplate(mocks.InvestigationTemplate().Id, "", investigationtemplatemodel.UpdateInvestigationRequest{
			TitlePrefix: mocks.InvestigationTemplate().TitlePrefix,
			Severity:    mocks.InvestigationTemplate().Severity,
			Status:      mocks.InvestigationTemplate().Status,
			Tags:        mocks.InvestigationTemplate().Tags,
			Tlp:         mocks.InvestigationTemplate().Tlp,
		})
		assert.Nil(t, err)
		assert.Equal(t, mocks.InvestigationTemplateResponse(), updatedInvestigationTemplate)
	})
}
