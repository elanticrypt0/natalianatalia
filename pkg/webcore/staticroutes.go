package webcore

import (
	"github.com/labstack/echo/v4"
)

const publicDir = "./public"

func SetupStaticRoutes(server *echo.Echo) {

	server.Static("/", publicDir)
	server.Static("/", publicDir+"/assets")
	server.Static("/", publicDir+"/js")
	server.Static("/", publicDir+"/css")
	server.Static("/", publicDir+"/images")

}
