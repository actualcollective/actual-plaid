package routes

import (
	"actual-plaid/internal/actual"
	"actual-plaid/internal/config"
	"actual-plaid/internal/plaid"
	"github.com/gofiber/fiber/v2"
)

type SuccessFrom struct {
	Token       string                    `json:"token"`
	BankCtx     string                    `json:"bankCtx"`
	PublicToken string                    `json:"publicToken"`
	Metadata    plaid.SuccessResponseMeta `json:"metadata"`
}

type AccountForm struct {
	Token   string `json:"token"`
	BankCtx string `json:"bankCtx"`
}

type TransactionForm struct {
	Token     string `json:"token"`
	BankCtx   string `json:"bankCtx"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	AcctId    string `json:"acctId"`
	Count     int32  `json:"count"`
	Offset    int32  `json:"offset"`
}

func PlaidRouter(app fiber.Router) {
	app.Post("/plaid/success", Success())

	app.Post("/plaid/accounts", Accounts())
	app.Post("/plaid/transactions", Transactions())
}

func Success() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		form := new(SuccessFrom)

		err = c.BodyParser(form)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		accessToken, err := plaid.ExchangePublicToken(&form.PublicToken)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		err = actual.SetBankContext(config.Config.ActualSecret, form.BankCtx, *accessToken, form.PublicToken)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		err = actual.PutTokenContent(form.Token, form.PublicToken, form.Metadata.Institution, form.Metadata.Accounts)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		return c.JSON(fiber.Map{"status": "ok"})
	}
}

func Accounts() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		form := new(AccountForm)

		err = c.BodyParser(form)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		err = actual.ValidateToken(form.Token)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		accessToken, err := actual.GetBankContext(config.Config.ActualSecret, form.BankCtx)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		accounts, err := plaid.GetAccounts(accessToken)

		return c.JSON(fiber.Map{"status": "ok", "data": accounts})
	}
}

func Transactions() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		form := new(TransactionForm)

		err = c.BodyParser(form)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		err = actual.ValidateToken(form.Token)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		accessToken, err := actual.GetBankContext(config.Config.ActualSecret, form.BankCtx)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "reason": err})
		}

		res, err := plaid.GetTransactions(accessToken, form.StartDate, form.EndDate, form.AcctId, form.Count, form.Offset)

		return c.JSON(fiber.Map{"status": "ok", "data": res})
	}
}
