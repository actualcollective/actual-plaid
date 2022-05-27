package actual

import (
	"actual-plaid/internal/config"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func MakeRequest(method string, path string, data *fiber.Map, dst interface{}, destruct bool) (a *fiber.Agent, err error) {
	a = fiber.AcquireAgent()
	a = a.Debug()
	a.Request().Header.SetMethod(method)
	a.Request().SetRequestURI(config.Config.ActualUrl + "/" + path)

	a = a.JSON(&data)

	if err := a.Parse(); err != nil {
		return nil, err
	}

	code, raw, errs := a.Bytes()
	if errs != nil {
		err = errs[0]
		return nil, err
	}

	if code == 200 && destruct {
		err = json.Unmarshal(raw, &dst)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		return a, nil
	}

	return nil, err
}

func ValidateToken(token string) (err error) {
	if config.Config.ActualSecret == token {
		return nil
	}

	return fmt.Errorf("token does not match with actual secret")
}
