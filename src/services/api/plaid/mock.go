package main

import (
	"context"

	plaid "github.com/plaid/plaid-go/v12/plaid"
)

type Plaid interface {
	ToPlaidError(err error) (plaid.PlaidError, error)
	ItemPublicTokenExchange(ctx context.Context, publicToken string ) (plaid.ItemPublicTokenExchangeResponse, error)
}

type PlaidClient struct{}
type MockPlaidClient struct {
	PlaidError plaid.PlaidError
	Err error
	ExchangeResp plaid.ItemPublicTokenExchangeResponse
}

func createClient() *plaid.APIClient {
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", "63962697de7ba8001361d7fe")
	configuration.AddDefaultHeader("PLAID-SECRET", "e72fdfe4452e82f53b5d0f57aebf1d")
	configuration.UseEnvironment(plaid.Sandbox)
	client := plaid.NewAPIClient(configuration)
	return client
}

func (t MockPlaidClient) GetAccessToken(o *plaid.ItemPublicTokenExchangeResponse) string {
	return t.ExchangeResp.AccessToken
}
func (t MockPlaidClient) GetItemId(o *plaid.ItemPublicTokenExchangeResponse) string {
	return t.ExchangeResp.ItemId
}
func (t MockPlaidClient) ToPlaidError(err error) (plaid.PlaidError, error) {
	return t.PlaidError, t.Err
}
func (t MockPlaidClient) ItemPublicTokenExchange(ctx context.Context, publicToken string ) (plaid.ItemPublicTokenExchangeResponse, error) {
	return t.ExchangeResp, t.Err
}
func (t PlaidClient) GetAccessToken(o *plaid.ItemPublicTokenExchangeResponse) string {
	return o.GetAccessToken()
}
func (t PlaidClient) GetItemId(o *plaid.ItemPublicTokenExchangeResponse) string{
	return o.GetItemId()
}
func (t PlaidClient) ToPlaidError(err error) (plaid.PlaidError, error) {
	plaidError, err := plaid.ToPlaidError(err)
	return plaidError, err
}
func (t PlaidClient) ItemPublicTokenExchange(ctx context.Context, publicToken string ) (plaid.ItemPublicTokenExchangeResponse, error) {
	client := createClient()
	exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken),
	).Execute()
	return exchangePublicTokenResp, err
}