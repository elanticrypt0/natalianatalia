package api

import (
	"github.com/k23dev/natalianatalia/app/routes"
	"github.com/k23dev/natalianatalia/pkg/webcore"
)

func ApiSetup(tapp *webcore.TangoApp) {

	// features routes
	routes.SetupApiRoutes(tapp)

}
