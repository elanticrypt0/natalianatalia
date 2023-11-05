package webcore

import (
	"github.com/elanticrypt0/go4it"
	"github.com/elanticrypt0/natalianatalia/api/nnconfig"
	"github.com/gofiber/fiber/v2"
)

type GasonlineApp struct {
	App      *go4it.App
	Fiber    *fiber.App
	NNConfig *nnconfig.NNConfig
}
