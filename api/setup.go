package api

import (
	"github.com/k23dev/tango/api/routes"
	"github.com/k23dev/tango/pkg/webcore"
)

func ApiSetup(tapp *webcore.TangoApp) {

	// features routes
	routes.SetupApiRoutes(tapp)

}
