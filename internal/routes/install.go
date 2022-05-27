package routes

import (
	"actual-plaid/internal/plaid"
	"encoding/base64"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type InstallOptions struct {
	TokenId   string `json:"tokenId"`
	BankCtxId string `json:"bankCtxId"`
	Timestamp int64  `json:"timestamp"`
}

type InstallForm struct {
	Options string `validate:"required"`
}

func InstallRouter(app fiber.Router) {
	app.Post("/install", Install())
}

func Install() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		form := new(InstallForm)

		err = c.BodyParser(form)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}
		c.Context()
		unescaped, _ := url.PathUnescape(form.Options)
		options := InstallOptions{}

		raw, err := base64.StdEncoding.DecodeString(unescaped)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		err = json.Unmarshal(raw, &options)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		plaidToken, err := plaid.LinkToken(options.BankCtxId)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		return c.JSON(fiber.Map{"status": "ok", "token": options.TokenId, "bankCtx": options.BankCtxId, "plaidToken": plaidToken})
	}
}
