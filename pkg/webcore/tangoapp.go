package webcore

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/k23dev/go4it"
)

type TangoApp struct {
	App   *go4it.App
	Fiber *fiber.App
}

func (tapp *TangoApp) PrintAppInfo() {
	fmt.Printf("Starting app: %s v%s", tapp.App.Config.App_name, tapp.App.Config.App_version)
}

func (tapp *TangoApp) GetAppUrl() string {
	return fmt.Sprintf("%s:%d", tapp.App.Config.App_server_host, tapp.App.Config.App_server_port)
}

func (tapp *TangoApp) GetPortAsStr() string {
	return fmt.Sprintf("%d", tapp.App.Config.App_server_port)
}
