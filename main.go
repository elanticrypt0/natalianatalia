package main

import (
	"fmt"
	"log"

	"github.com/elanticrypt0/go4it"
	"github.com/elanticrypt0/natalianatalia/api/nnconfig"
	"github.com/elanticrypt0/natalianatalia/api/routes"
	"github.com/elanticrypt0/natalianatalia/pkg/webcore"
	"github.com/elanticrypt0/natalianatalia/pkg/webcore/helpers"
	"github.com/elanticrypt0/natalianatalia/pkg/webcore_features"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app_config := go4it.NewApp("./appconfig")
	// make the connection
	app_config.Connect2Db("local")
	app_config.DB.SetPrimaryDB(0)

	gas := webcore.GasonlineApp{
		App:      &app_config,
		Fiber:    fiber.New(fiber.Config{}),
		NNConfig: nnconfig.NewNNConfig(),
	}

	// CORS
	// necesario para poder utilizar la WUI
	gas.Fiber.Use(cors.New())
	gas.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: gas.App.Config.App_CORS_Origins,
		AllowHeaders: gas.App.Config.APP_CORS_Headers,
	}))

	gas.Fiber.Use(recover.New())

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
		fmt.Println("Iniciando app...")
	}

	fmt.Printf("%v", gas.NNConfig)

	portAsStr := fmt.Sprintf("%d", gas.App.Config.App_server_port)

	helpers.OpenInBrowser("http://" + gas.App.Config.App_server_host + ":" + portAsStr)

	log.Fatal(gas.Fiber.Listen(gas.App.Config.App_server_host+":"+portAsStr), "Server is running on port "+portAsStr)

}
