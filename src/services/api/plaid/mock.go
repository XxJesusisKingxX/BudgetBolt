package main

import (
	"context"
	"time"

	plaid "github.com/plaid/plaid-go/v12/plaid"
)

type Plaid interface {
	ToPlaidError(err error) (plaid.PlaidError, error)
	ItemPublicTokenExchange(ctx context.Context, publicToken string ) (plaid.ItemPublicTokenExchangeResponse, error)
	AccountsGet(ctx context.Context, accessToken string ) (plaid.AccountsGetResponse, error)
	InvestmentsTransactionsGet(ctx context.Context, accessToken string ) (plaid.InvestmentsTransactionsGetResponse, error)
	InvestmentsHoldingsGet(ctx context.Context, accessToken string ) (plaid.InvestmentsHoldingsGetResponse, error)	
}

type PlaidClient struct{}
type MockPlaidClient struct {
	PlaidError plaid.PlaidError
	Err error
	ExchangeResp plaid.ItemPublicTokenExchangeResponse
	AccountsResp plaid.AccountsGetResponse
	InvestTransResp plaid.InvestmentsTransactionsGetResponse
	InvestHoldResp plaid.InvestmentsHoldingsGetResponse
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
func (t MockPlaidClient) AccountsGet(ctx context.Context, accessToken string ) (plaid.AccountsGetResponse, error) {
	return t.AccountsResp, t.Err
}
func (t MockPlaidClient) InvestmentsTransactionsGet(ctx context.Context, accessToken string ) (plaid.InvestmentsTransactionsGetResponse, error) {
	return t.InvestTransResp, t.Err
}
func (t MockPlaidClient) InvestmentsHoldingsGet(ctx context.Context, accessToken string ) (plaid.InvestmentsHoldingsGetResponse, error) {
	return t.InvestHoldResp, t.Err
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
func (t PlaidClient) AccountsGet(ctx context.Context, accessToken string ) (plaid.AccountsGetResponse, error) {
	client := createClient()
	accountsGetResp, _, err := client.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*plaid.NewAccountsGetRequest(accessToken),
	).Execute()
	return accountsGetResp, err
}
func (t PlaidClient) InvestmentsTransactionsGet(ctx context.Context, accessToken string ) (plaid.InvestmentsTransactionsGetResponse, error) {
	endDate := time.Now().Local().Format("2006-01-02")
	startDate := time.Now().Local().Add(-30 * 24 * time.Hour).Format("2006-01-02")
	request := plaid.NewInvestmentsTransactionsGetRequest(accessToken, startDate, endDate)
	invTxResp, _, err := client.PlaidApi.InvestmentsTransactionsGet(ctx).InvestmentsTransactionsGetRequest(*request).Execute()
	return invTxResp, err
}
func (t PlaidClient) InvestmentsHoldingsGet(ctx context.Context, accessToken string ) (plaid.InvestmentsHoldingsGetResponse, error) {
	holdingsGetResp, _, err := client.PlaidApi.InvestmentsHoldingsGet(ctx).InvestmentsHoldingsGetRequest(
		*plaid.NewInvestmentsHoldingsGetRequest(accessToken),
	).Execute()
	return holdingsGetResp, err
}