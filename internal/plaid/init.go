package plaid

import (
	"actual-plaid/internal/config"
	"context"
	"fmt"
	"github.com/plaid/plaid-go/v3/plaid"
	"github.com/samber/lo"
	"strings"
)

type SuccessResponse struct {
	PublicToken string              `json:"publicToken"`
	Metadata    SuccessResponseMeta `json:"metadata"`
}

type SuccessResponseMeta struct {
	Institution plaid.Institution   `json:"institution"`
	Accounts    []plaid.AccountBase `json:"accounts"`
}

func craftPlaidUser(bankCtx string) *plaid.LinkTokenCreateRequestUser {
	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: bankCtx,
	}

	return &user
}

func switchEnv(plaidMode string) plaid.Environment {
	switch plaidMode {
	case "sandbox":
		return plaid.Sandbox
	case "development":
		return plaid.Development
	case "production":
		return plaid.Production
	}

	return plaid.Sandbox
}

func InitClient() *plaid.APIClient {
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", config.Config.PlaidClientId)
	configuration.AddDefaultHeader("PLAID-SECRET", config.Config.PlaidSecret)
	configuration.UseEnvironment(switchEnv(config.Config.PlaidMode))
	client := plaid.NewAPIClient(configuration)

	return client
}

func LinkToken(bankCtx string) (token *string, err error) {
	countries := strings.Split(config.Config.PlaidCountryCodes, ",")

	countryCodes := lo.Map[string, plaid.CountryCode](countries, func(country string, _ int) plaid.CountryCode {
		countryCode, _ := plaid.NewCountryCodeFromValue(strings.ToUpper(country))
		return *countryCode
	})

	user := craftPlaidUser(bankCtx)
	request := plaid.NewLinkTokenCreateRequest(
		config.Config.PlaidClientName,
		config.Config.PlaidLanguage,
		countryCodes,
		*user,
	)

	request.SetProducts([]plaid.Products{plaid.PRODUCTS_TRANSACTIONS})
	request.SetLinkCustomizationName("default")

	client := InitClient()
	resp, raw, err := client.PlaidApi.LinkTokenCreate(context.Background()).LinkTokenCreateRequest(*request).Execute()
	if err != nil {
		fmt.Println(err, raw)
		return nil, err
	}
	linkToken := resp.GetLinkToken()

	return &linkToken, nil
}

func ExchangePublicToken(publicToken *string) (token *string, err error) {
	request := plaid.NewItemPublicTokenExchangeRequest(*publicToken)
	client := InitClient()
	resp, raw, err := client.PlaidApi.ItemPublicTokenExchange(context.Background()).ItemPublicTokenExchangeRequest(*request).Execute()
	if err != nil {
		fmt.Println(err, raw)
		return nil, err
	}

	accessToken := resp.GetAccessToken()

	return &accessToken, nil
}

func GetAccounts(accessToken *string) (accounts []plaid.AccountBase, err error) {
	request := plaid.NewAccountsGetRequest(*accessToken)
	client := InitClient()
	resp, raw, err := client.PlaidApi.AccountsGet(context.Background()).AccountsGetRequest(*request).Execute()
	if err != nil {
		fmt.Println(err, raw)
		return nil, err
	}

	return resp.Accounts, nil
}

func GetTransactions(accessToken *string, startDate string, endDate string, acctId string, count int32, offset int32) (res *plaid.TransactionsGetResponse, err error) {
	request := plaid.NewTransactionsGetRequest(*accessToken, startDate, endDate)

	var accountsIds []string
	accountsIds = append(accountsIds, acctId)

	options := plaid.TransactionsGetRequestOptions{
		AccountIds: &accountsIds,
		Count:      &count,
		Offset:     &offset,
	}

	request.SetOptions(options)
	client := InitClient()
	response, raw, err := client.PlaidApi.TransactionsGet(context.Background()).TransactionsGetRequest(*request).Execute()
	if err != nil {
		fmt.Println(err, raw)
		return nil, err
	}

	res = &response

	return res, nil
}
