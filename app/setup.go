package app

import (
	"github.com/k23dev/natalianatalia/app/models"
	"github.com/k23dev/natalianatalia/app/routes"
	"github.com/k23dev/natalianatalia/pkg/webcore"
)

func AppSetup(tapp *webcore.TangoApp) {

	// features routes
	routes.SetupAppRoutes(tapp)

	if tapp.App.Config.App_debug_mode {
		tapp.App.DB.Primary.AutoMigrate(&models.Tanga{}, &models.TangaField{}, &models.Scripts{})
	}

}
