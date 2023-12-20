package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/k23dev/go4it"
	"github.com/k23dev/tango/api"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/k23dev/tango/pkg/webcore_features"
)

func main() {

	app_config := go4it.NewApp("./config/appconfig")

	tapp := webcore.TangoApp{
		App: &app_config,
		Fiber: fiber.New(fiber.Config{
			Prefork:               false,
			CaseSensitive:         true,
			StrictRouting:         true,
			ServerHeader:          "Fiber",
			AppName:               app_config.Config.App_name,
			DisableStartupMessage: false,
			PassLocalsToViews:     true,
		}),
	}
	tapp.PrintAppInfo()

	// make the connection
	app_config.Connect2Db("local")
	app_config.DB.SetPrimaryDB(0)

	// middleware
	webcore.MiddlewareSetup(&tapp)

	// Routes setup

	// webcore setup routes
	if tapp.App.Config.App_setup_enabled {
		webcore_features.SetupRoutes(&tapp)
		webcore_features.SetupOnStartup(&tapp)
	}

	api.ApiSetup(&tapp)

	// static routes
	webcore.SetupStaticRoutes(tapp.Fiber)

	// go4it.OpenInBrowser("http://" + gas.GetAppUrl())

	log.Fatal(tapp.Fiber.Listen(":" + tapp.GetPortAsStr()))

}
