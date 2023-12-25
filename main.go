package main

import (
	"github.com/k23dev/natalianatalia/app/routes"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/k23dev/natalianatalia/pkg/webcore_features"

	"github.com/k23dev/go4it"
	"github.com/labstack/echo/v4"
)

func main() {

	app_config := go4it.NewApp("./config/appconfig")

	tapp := webcore.TangoApp{
		App:    &app_config,
		Server: echo.New(),
	}

	// Database connections
	app_config.Connect2Db("local")
	app_config.DB.SetPrimaryDB(0)

	tapp.PrintAppInfo()

	// Middleware
	webcore.MiddlewareSetup(&tapp)

	//  Routes

	webcore_features.SetupRoutes(&tapp)

	if tapp.App.Config.App_setup_enabled && tapp.App.Config.App_debug_mode {
		routes.SetupApiRoutes(&tapp)
	}

	webcore.SetupStaticRoutes(tapp.Server)

	// open app in default browser
	go4it.OpenInBrowser("http://" + tapp.GetAppUrl())

	// Start server
	tapp.Server.Logger.Fatal(tapp.Server.Start(":" + tapp.GetPortAsStr()))

}
