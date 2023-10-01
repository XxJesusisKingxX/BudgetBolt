package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v12/plaid"

	user "services/internal/user_management/db/model"
	request "services/internal/utils/http"
)

func CreateLinkToken(c *gin.Context, ps Plaid, httpClient request.HTTP, plaidapi *plaid.APIClient) {
	ctx := context.Background()
	uid, _ := c.Cookie("UID")

	var profile user.Profile
	body := fmt.Sprintf("uid=%v", uid)
	status, resp, err := httpClient.POST("profile/get", body)
	request.ParseResponse(resp, &profile)

	if status != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	countryCodes := convertCountryCodes(strings.Split("US", ","))
	products := convertProducts(strings.Split("transactions", ","))
	request := ps.NewLinkTokenCreateRequest(uid, uid, countryCodes, products, "")
	linkTokenCreateResp, err := ps.CreateLinkToken(plaidapi, ctx, request)
	if err != nil {
		RenderError(c, err, PlaidClient{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"link_token": linkTokenCreateResp.GetLinkToken()})
}

func CreateAccessToken(c *gin.Context, ps Plaid, plaidapi *plaid.APIClient, httpClient request.HTTP, debug bool) {
	ctx := context.Background()
	publicToken := c.PostForm("public_token")
	uid, _ := c.Cookie("UID")
	exchangePublicTokenResp, err := ps.ItemPublicTokenExchange(plaidapi, ctx, publicToken)

	if err != nil {
		RenderError(c, err, PlaidClient{})
		return
	}
	var id int64
	accessToken := exchangePublicTokenResp.GetAccessToken()
	itemID := exchangePublicTokenResp.GetItemId()

	var profile user.Profile
	body := fmt.Sprintf("uid=%v", uid)
	status, resp, err := httpClient.POST("profile/get", body)
	request.ParseResponse(resp, &profile)
	if status != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	id = profile.ID
	body = fmt.Sprintf("id=%v&token=%v&itemId=%v", id, accessToken, itemID)
	status, _, _ = httpClient.POST("token/create", body)
	if status != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func CreateAccounts(c *gin.Context, ps Plaid, plaidapi *plaid.APIClient, httpClient request.HTTP, debug bool) {
	ctx := context.Background()
	uid, _ := c.Cookie("UID")

	var profile user.Profile
	body := fmt.Sprintf("uid=%v", uid)
	status, resp, _ := httpClient.POST("profile/get", body)
	request.ParseResponse(resp, &profile)

	if status != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var token user.Tokens
	url := fmt.Sprintf("token/get?uid=%v", uid)
	status, resp, _ = httpClient.GET(url)
	request.ParseResponse(resp, &token)

	if status != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	
	for _, token := range token.Tokens {
		accessToken := token.Token
		accounts, err := ps.AccountsGet(plaidapi, ctx, accessToken)
		
		if err != nil {
			RenderError(c, err, PlaidClient{})
			return
		}
		if !debug {
			id := profile.ID
			accountsJson, _ := json.Marshal(accounts)
			body = fmt.Sprintf("id=%v&accounts=%v", id, string(accountsJson))
			httpClient.POST("accounts/store", body)
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}

func CreateTransactions(c *gin.Context, ps Plaid, plaidapi *plaid.APIClient, httpClient request.HTTP, debug bool) {
	ctx := context.Background()
	uid, _ := c.Cookie("UID")

	var profile user.Profile
	body := fmt.Sprintf("uid=%v", uid)
	status, resp, err := httpClient.POST("profile/get", body)
	request.ParseResponse(resp, &profile)

	if status != 200 {
		err = errors.New("")
		RenderError(c, err, PlaidClient{})
		return
	}

	var token user.Tokens
	url := fmt.Sprintf("token/get?uid=%v", uid)
	status, resp, _ = httpClient.GET(url)
	request.ParseResponse(resp, &token)

	if status != 200 {
		err := errors.New("")
		RenderError(c, err, PlaidClient{})
		return
	}

	for _, token := range token.Tokens{
		accessToken := token.Token
		var cursor *string
		var transactions []plaid.Transaction
		hasMore := true
		for hasMore {
			resp, err := ps.NewTransactionsSyncRequest(plaidapi, ctx, accessToken, cursor)
			if err != nil {
				RenderError(c, err, PlaidClient{})
				return
			}
			transactions = append(transactions, resp.GetAdded()...)
			hasMore = resp.GetHasMore()
			nextCursor := resp.GetNextCursor()
			cursor = &nextCursor
		}
		if !debug {
			id := profile.ID
			transactionJson, _ := json.Marshal(transactions)
			body = fmt.Sprintf("id=%v&transactions=%v", id, string(transactionJson))
			httpClient.POST("transactions/store", body)
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}