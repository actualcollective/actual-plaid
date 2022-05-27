package plaid

import (
	"github.com/plaid/plaid-go/v3/plaid"
	"math/rand"
)

func Mock() (plaid.Institution, []plaid.AccountBase) {
	var products []plaid.Products
	var countryCodes []plaid.CountryCode
	products = append(products, plaid.PRODUCTS_TRANSACTIONS)
	countryCodes = append(countryCodes, plaid.COUNTRYCODE_NL)

	institution := plaid.Institution{
		InstitutionId:             "1234",
		Name:                      "ActualBank",
		Products:                  products,
		CountryCodes:              countryCodes,
		Url:                       plaid.NullableString{},
		PrimaryColor:              plaid.NullableString{},
		Logo:                      plaid.NullableString{},
		RoutingNumbers:            nil,
		Oauth:                     false,
		Status:                    plaid.NullableInstitutionStatus{},
		PaymentInitiationMetadata: plaid.NullablePaymentInitiationMetadata{},
		AuthMetadata:              plaid.NullableAuthMetadata{},
		AdditionalProperties:      nil,
	}

	var accounts []plaid.AccountBase

	float := rand.Float64()
	_mask := "5932"
	mask := plaid.NewNullableString(&_mask)

	accounts = append(accounts, plaid.AccountBase{
		AccountId:            "1234ActualBank5678",
		Balances:             plaid.AccountBalance{Available: *plaid.NewNullableFloat64(&float), Current: *plaid.NewNullableFloat64(&float)},
		Mask:                 *mask,
		Name:                 "ActualAccount",
		OfficialName:         plaid.NullableString{},
		Type:                 plaid.ACCOUNTTYPE_DEPOSITORY,
		Subtype:              plaid.NullableAccountSubtype{},
		VerificationStatus:   nil,
		AdditionalProperties: nil,
	})

	return institution, accounts
}
