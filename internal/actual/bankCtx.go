package actual

import (
	"github.com/gofiber/fiber/v2"
)

func SetBankContext(secret string, bankCtxId string, payload string, publicToken string) (err error) {
	data := fiber.Map{
		"token":      secret,
		"contextId":  bankCtxId,
		"payload":    payload,
		"externalId": publicToken,
	}

	_, err = MakeRequest(fiber.MethodPost, "/integrations/set-context", &data, nil, false)
	if err != nil {
		return err
	}

	return nil
}

func GetBankContext(secret string, bankCtxId string) (payload *string, err error) {
	data := fiber.Map{
		"token":     secret,
		"contextId": bankCtxId,
	}

	dst := struct {
		Status string `json:"status"`
		Data   string `json:"data"`
	}{}

	_, err = MakeRequest(fiber.MethodPost, "/integrations/get-context", &data, &dst, true)
	if err != nil {
		return nil, err
	}

	return &dst.Data, nil
}
