package database

import (
	"context"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/tenantmodel"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateTenant
// Creates a new tenant in the database.
func (d *database) CreateTenant(tenant tenantmodel.Tenant) (*tenantmodel.Tenant, error) {
	var result tenantmodel.Tenant
	tenant.TenantId = uuid.NewString()
	_, err := d.TenantCollection.InsertOne(context.TODO(), tenant)
	if err != nil {
		return nil, err
	}
	err = d.TenantCollection.FindOne(context.TODO(), bson.D{{"tenantId", tenant.TenantId}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTenant
// Deletes a tenant from the database.
func (d *database) DeleteTenant(tenantId string) error {
	result, err := d.TenantCollection.DeleteOne(context.TODO(), bson.D{{"tenantId", tenantId}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return utils.ErrTenantNotFound
	}
	return err
}

// GetTenant
// Returns a single tenant from the database.
func (d *database) GetTenant(tenantId string) (*tenantmodel.Tenant, error) {
	var tenant tenantmodel.Tenant
	err := d.TenantCollection.FindOne(context.TODO(), bson.D{{"tenantId", tenantId}}).Decode(&tenant)
	if err == mongo.ErrNoDocuments {
		return nil, utils.ErrTenantNotFound
	} else if err != nil {
		return nil, err
	}
	return &tenant, err
}

// UpdateTenant
// Updates a tenant in the database.
func (d *database) UpdateTenant(tenantId string, tenant tenantmodel.UpdateTenantRequest) (*tenantmodel.Tenant, error) {
	var (
		oldResult tenantmodel.Tenant
		newResult tenantmodel.Tenant
	)
	filter := bson.D{{"tenantId", tenantId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "name", Value: tenant.Name},
		}},
	}
	err := d.TenantCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldResult)
	if err != nil {
		return nil, err
	}
	err = d.TenantCollection.FindOne(context.TODO(), bson.D{{"tenantId", oldResult.TenantId}}).Decode(&newResult)
	if err != nil {
		return nil, err
	}
	return &newResult, err
}
