package main

import (
	"github.com/k23dev/natalianatalia/app"
	"github.com/k23dev/natalianatalia/pkg/nn"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/k23dev/natalianatalia/pkg/webcore_features"

	"github.com/k23dev/go4it"
	"github.com/labstack/echo/v4"
)

func main() {

	app_config := go4it.NewApp("./config/appconfig")

	tapp := webcore.TangoApp{
		App:      &app_config,
		Server:   echo.New(),
		NNConfig: nn.NewNNConfig(),
	}

	// Database connections
	app_config.Connect2Db("local")
	app_config.DB.SetPrimaryDB(0)

	tapp.PrintAppInfo()

	// Middleware
	webcore.MiddlewareSetup(&tapp)

	//  Routes
	if tapp.App.Config.App_setup_enabled && tapp.App.Config.App_debug_mode {
		webcore_features.SetupRoutes(&tapp)
	}

	webcore.SetupStaticRoutes(tapp.Server)

	app.AppSetup(&tapp)

	// open app in default browser
	go4it.OpenInBrowser("http://" + tapp.GetAppUrl())

	// Start server
	// tapp.Server.Logger.Fatal(tapp.Server.Start(":" + tapp.GetPortAsStr()))
	tapp.Server.Start(":" + tapp.GetPortAsStr())

}
