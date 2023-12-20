package api

import (
	"github.com/k23dev/tango/api/routes"
	"github.com/k23dev/tango/pkg/webcore"
)

func ApiSetup(tapp *webcore.TangoApp) {

	tapp.App.DB.Primary.AutoMigrate()
	// features routes
	routes.SetupApiRoutes(tapp)

}
