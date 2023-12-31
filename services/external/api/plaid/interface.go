package api

import (
	"context"
	"time"

	"github.com/plaid/plaid-go/v12/plaid"
)

type Plaid interface {
	ToPlaidError(err error) (plaid.PlaidError, error)
	ItemPublicTokenExchange(client *plaid.APIClient, ctx context.Context, publicToken string ) (plaid.ItemPublicTokenExchangeResponse, error)
	AccountsGet(client *plaid.APIClient, ctx context.Context, accessToken string ) ([]plaid.AccountBase, error)
	InvestmentsTransactionsGet(client *plaid.APIClient, ctx context.Context, accessToken string ) (plaid.InvestmentsTransactionsGetResponse, error)
	InvestmentsHoldingsGet(client *plaid.APIClient, ctx context.Context, accessToken string ) (plaid.InvestmentsHoldingsGetResponse, error)
	CreateLinkToken(client *plaid.APIClient, ctx context.Context, request *plaid.LinkTokenCreateRequest) (plaid.LinkTokenCreateResponse, error)
	NewLinkTokenCreateRequest(name string, user string, countryCodes []plaid.CountryCode, products []plaid.Products, redirectURI string) (*plaid.LinkTokenCreateRequest)
	NewTransactionsSyncRequest(client *plaid.APIClient, ctx context.Context, accessToken string, cursor *string) (plaid.TransactionsSyncResponse, error)
	NewTransactionsRecurringGetRequest(client *plaid.APIClient, ctx context.Context, accessToken string, accounts []string) (plaid.TransactionsRecurringGetResponse, error)
}

type PlaidClient struct{}
type MockPlaidClient struct {
	User string
	Name string
	CountryCode []plaid.CountryCode
	Products []plaid.Products
	RedirectURI string
	PlaidError plaid.PlaidError
	Err error
	Accounts []plaid.AccountBase
	ExchangeResp plaid.ItemPublicTokenExchangeResponse
	AccountsResp plaid.AccountsGetResponse
	InvestTransResp plaid.InvestmentsTransactionsGetResponse
	InvestHoldResp plaid.InvestmentsHoldingsGetResponse
	TokenResp plaid.LinkTokenCreateResponse
	SyncResp plaid.TransactionsSyncResponse
	RecurringResp plaid.TransactionsRecurringGetResponse
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
func (t MockPlaidClient) ItemPublicTokenExchange(client *plaid.APIClient, ctx context.Context, publicToken string ) (plaid.ItemPublicTokenExchangeResponse, error) {
	return t.ExchangeResp, t.Err
}
func (t MockPlaidClient) AccountsGet(client *plaid.APIClient, ctx context.Context, accessToken string ) ([]plaid.AccountBase, error) {
	return t.Accounts, t.Err
}
func (t MockPlaidClient) InvestmentsTransactionsGet(client *plaid.APIClient, ctx context.Context, accessToken string ) (plaid.InvestmentsTransactionsGetResponse, error) {
	return t.InvestTransResp, t.Err
}
func (t MockPlaidClient) InvestmentsHoldingsGet(client *plaid.APIClient, ctx context.Context, accessToken string ) (plaid.InvestmentsHoldingsGetResponse, error) {
	return t.InvestHoldResp, t.Err
}
func (t MockPlaidClient) CreateLinkToken(client *plaid.APIClient, ctx context.Context, request *plaid.LinkTokenCreateRequest) (plaid.LinkTokenCreateResponse, error) {
	return t.TokenResp, t.Err
}
func (t MockPlaidClient) NewLinkTokenCreateRequest(name string, user string, countryCodes []plaid.CountryCode, products []plaid.Products, redirectURI string) (*plaid.LinkTokenCreateRequest) {
	request := plaid.NewLinkTokenCreateRequest(
		t.Name,
		"en",
		t.CountryCode,
		plaid.LinkTokenCreateRequestUser{ ClientUserId: t.User },
	)
	request.SetProducts(t.Products)
	if redirectURI != "" {
		request.SetRedirectUri(t.RedirectURI)
	}
	return request
}
func (t MockPlaidClient) NewTransactionsSyncRequest(client *plaid.APIClient, ctx context.Context, accessToken string, cursor *string) (plaid.TransactionsSyncResponse, error) {
	return t.SyncResp, t.Err
}
func (t MockPlaidClient) NewTransactionsRecurringGetRequest(client *plaid.APIClient, ctx context.Context, accessToken string, accounts []string) (plaid.TransactionsRecurringGetResponse, error) {

	return t.RecurringResp, t.Err
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
func (t PlaidClient) ItemPublicTokenExchange(client *plaid.APIClient, ctx context.Context, publicToken string ) (plaid.ItemPublicTokenExchangeResponse, error) {
	exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken),
	).Execute()
	return exchangePublicTokenResp, err
}
func (t PlaidClient) AccountsGet(client *plaid.APIClient, ctx context.Context, accessToken string ) ([]plaid.AccountBase, error) {
	accountsGetResp, _, err := client.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*plaid.NewAccountsGetRequest(accessToken),
	).Execute()
	return accountsGetResp.Accounts, err
}
func (t PlaidClient) InvestmentsTransactionsGet(client *plaid.APIClient, ctx context.Context, accessToken string ) (plaid.InvestmentsTransactionsGetResponse, error) {
	endDate := time.Now().Local().Format("2006-01-02")
	startDate := time.Now().Local().Add(-30 * 24 * time.Hour).Format("2006-01-02")
	request := plaid.NewInvestmentsTransactionsGetRequest(accessToken, startDate, endDate)
	invTxResp, _, err := client.PlaidApi.InvestmentsTransactionsGet(ctx).InvestmentsTransactionsGetRequest(*request).Execute()
	return invTxResp, err
}
func (t PlaidClient) InvestmentsHoldingsGet(client *plaid.APIClient, ctx context.Context, accessToken string ) (plaid.InvestmentsHoldingsGetResponse, error) {
	holdingsGetResp, _, err := client.PlaidApi.InvestmentsHoldingsGet(ctx).InvestmentsHoldingsGetRequest(
		*plaid.NewInvestmentsHoldingsGetRequest(accessToken),
	).Execute()
	return holdingsGetResp, err
}
func (t PlaidClient) CreateLinkToken(client *plaid.APIClient, ctx context.Context, request *plaid.LinkTokenCreateRequest) (plaid.LinkTokenCreateResponse, error) {
	linkTokenCreateResp, _, err := client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()
	return linkTokenCreateResp, err
}
func (t PlaidClient) NewLinkTokenCreateRequest(name string, user string, countryCodes []plaid.CountryCode, products []plaid.Products, redirectURI string) (*plaid.LinkTokenCreateRequest) {
	request := plaid.NewLinkTokenCreateRequest(
		name,
		"en",
		countryCodes,
		plaid.LinkTokenCreateRequestUser{ ClientUserId: user },
	)
	request.SetProducts(products)
	if redirectURI != "" {
		request.SetRedirectUri(redirectURI)
	}
	return request
}
func (t PlaidClient) NewTransactionsSyncRequest(client *plaid.APIClient, ctx context.Context, accessToken string, cursor *string) (plaid.TransactionsSyncResponse, error) {
	include := true
	options := plaid.TransactionsSyncRequestOptions{
    	IncludePersonalFinanceCategory: &include,
	}
	request := plaid.NewTransactionsSyncRequest(accessToken)
	request.SetOptions(options)
	if cursor != nil {
		request.SetCursor(*cursor)
	}
	resp, _, err := client.PlaidApi.TransactionsSync(ctx).TransactionsSyncRequest(*request).Execute()
	return resp, err
}
func (t PlaidClient) NewTransactionsRecurringGetRequest(client *plaid.APIClient, ctx context.Context, accessToken string, accounts []string) (plaid.TransactionsRecurringGetResponse, error) {
	include := true
	options := plaid.TransactionsRecurringGetRequestOptions{
    	IncludePersonalFinanceCategory: &include,
	}
	request := plaid.NewTransactionsRecurringGetRequest(accessToken, accounts)
	request.SetOptions(options)
	resp, _, err := client.PlaidApi.TransactionsRecurringGet(ctx).TransactionsRecurringGetRequest(*request).Execute()
	return resp, err
}