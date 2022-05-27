package actual

import (
	plaidCtx "actual-plaid/internal/plaid"
	"github.com/gofiber/fiber/v2"
	"github.com/plaid/plaid-go/v3/plaid"
)

func PutTokenContent(token string, plaidToken string, institution plaid.Institution, accounts []plaid.AccountBase) (err error) {
	data := fiber.Map{
		"token": token,
		"data": plaidCtx.SuccessResponse{
			PublicToken: plaidToken,
			Metadata: plaidCtx.SuccessResponseMeta{
				Institution: institution,
				Accounts:    accounts,
			},
		},
	}

	_, err = MakeRequest(fiber.MethodPost, "/plaid/put-web-token-contents", &data, nil, false)
	if err != nil {
		return err
	}

	return nil
}
