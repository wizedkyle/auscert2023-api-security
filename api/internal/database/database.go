package database

import (
	"context"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"os"
)

var (
	Client = database{}
)

type database struct {
	AccessKeysCollection             *mongo.Collection
	IncidentCollection               *mongo.Collection
	IncidentCommentsCollection       *mongo.Collection
	InvestigationCollection          *mongo.Collection
	InvestigationTemplatesCollection *mongo.Collection
	Client                           *mongo.Client
	TenantCollection                 *mongo.Collection
	UserCollection                   *mongo.Collection
	WebhooksCollection               *mongo.Collection
}

const (
	AccessKeysCollection             = "accessKeys"
	IncidentCollection               = "incidents"
	IncidentCommentCollection        = "incidentComments"
	InvestigationCollection          = "investigations"
	InvestigationTemplatesCollection = "investigationsTemplates"
	DatabaseName                     = "forgeResponse"
	TenantCollection                 = "tenants"
	UserCollection                   = "users"
	WebhooksCollection               = "webhooks"
)

func Init() {
	var (
		opts *options.ClientOptions
	)
	if os.Getenv("GIN_MODE") == "release" {
		opts = options.Client().ApplyURI("mongodb://localhost:27017")
	} else {
		opts = options.Client().ApplyURI("mongodb://mongodb:27017")
	}
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		utils.Logger.Fatal("failed to connect to database", zap.Error(err))
	}
	Client.AccessKeysCollection = client.Database(DatabaseName).Collection(AccessKeysCollection)
	Client.IncidentCollection = client.Database(DatabaseName).Collection(IncidentCollection)
	Client.IncidentCommentsCollection = client.Database(DatabaseName).Collection(IncidentCommentCollection)
	Client.InvestigationCollection = client.Database(DatabaseName).Collection(InvestigationCollection)
	Client.InvestigationTemplatesCollection = client.Database(DatabaseName).Collection(InvestigationTemplatesCollection)
	Client.TenantCollection = client.Database(DatabaseName).Collection(TenantCollection)
	Client.UserCollection = client.Database(DatabaseName).Collection(UserCollection)
	Client.WebhooksCollection = client.Database(DatabaseName).Collection(WebhooksCollection)
	Client.Client = client
}
