package main

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/database"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/routes"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/routes/accesskeys_management"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/routes/event_management"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/routes/healthcheck_management"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/routes/incident_management"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/routes/investigation_management"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/routes/scope_management"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/routes/tenant_management"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/routes/user_management"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/routes/webhook_management"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
)

func main() {
	utils.GenerateLogger()
	database.Init()
	router := routes.GenerateRouter()
	healthcheck_management.GenerateHealthCheckManagementRoutes(router)
	accesskeys_management.GenerateAccessKeyManagementRoutes(router)
	event_management.GenerateEventManagementRoutes(router)
	incident_management.GenerateIncidentManagementRoutes(router)
	investigation_management.GenerateInvestigationManagementRoutes(router)
	scope_management.GenerateScopeManagementRoutes(router)
	tenant_management.GenerateTenantManagementRoutes(router)
	user_management.GenerateUserManagementRoutes(router)
	webhook_management.GenerateWebhookManagementRoutes(router)
	router.Run(":9000")
}
