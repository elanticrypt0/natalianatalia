package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/k23dev/go4it"
	"github.com/k23dev/natalianatalia/api/nnconfig"
	"github.com/k23dev/natalianatalia/api/routes"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/k23dev/natalianatalia/pkg/webcore/helpers"
	"github.com/k23dev/natalianatalia/pkg/webcore_features"
)

func main() {

	app_config := go4it.NewApp("./config/appconfig")
	// make the connection
	app_config.Connect2Db("local")
	app_config.DB.SetPrimaryDB(0)

	gas := webcore.GasonlineApp{
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
		NNConfig: nnconfig.NewNNConfig(),
	}

	// CORS
	// necesario para poder utilizar la WUI
	gas.Fiber.Use(cors.New())
	gas.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: gas.App.Config.App_CORS_origins,
		AllowHeaders: gas.App.Config.App_CORS_headers,
	}))

	gas.Fiber.Use(recover.New())

	gas.PrintAppInfo()

	// features routes
	routes.SetupFeaturesRoutes(&gas)
	// webcore features
	webcore_features.SetupRoutes(&gas)
	// static routes
	webcore.SetupStaticRoutes(gas.Fiber)

	// siempre se migran las tablas cuando comienza la app
	if gas.App.Config.App_debug_mode && gas.App.Config.App_setup_enabled {
		fmt.Println("Migrando las bases de datos...")
		webcore_features.SetupOnStartup(&gas)
	}

	// fmt.Printf("%v", gas.NNConfig)

	portAsStr := fmt.Sprintf("%d", gas.App.Config.App_server_port)

	helpers.OpenInBrowser("http://" + gas.GetAppUrl())

	log.Fatal(gas.Fiber.Listen(":" + portAsStr))

}
