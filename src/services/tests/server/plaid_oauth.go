package main

import (
	"context"
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"os"

	plaid "github.com/plaid/plaid-go/v12/plaid"
)

type IndexPageData struct {
    LinkToken string
}
type PlaidLinkCallback struct {
	PublicToken string `json:"public_token"`
}

func main() {
	ctx := context.Background()

	// Set test enviornment
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", "63962697de7ba8001361d7fe")
	configuration.AddDefaultHeader("PLAID-SECRET", "e72fdfe4452e82f53b5d0f57aebf1d")
	configuration.UseEnvironment(plaid.Sandbox)
	client := plaid.NewAPIClient(configuration)

	// Create link token
	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: "Test_User",
	}
	request := plaid.NewLinkTokenCreateRequest(
	  "Plaid Test",
	  "en",
	  []plaid.CountryCode{plaid.COUNTRYCODE_US},
	  user,
	)
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})
	resp, _, _ := client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	http.HandleFunc("/plaid-callback", handlePlaidCallback)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Retrieve the link token
        linkToken := resp.GetLinkToken()

        // Create the data to pass to the HTML template
        data := IndexPageData{
            LinkToken: linkToken,
        }

        // Parse the HTML template
		wd, _ := os.Getwd()
        tmpl, err := template.ParseFiles(filepath.Join(wd, "index.html"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Execute the template with the data
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    http.ListenAndServe(":8080", nil)
}

func handlePlaidCallback(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var callbackData PlaidLinkCallback
		err := json.NewDecoder(r.Body).Decode(&callbackData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		wd, _ := os.Getwd()
		file, _ := os.Create(filepath.Join(wd, "public_token.txt"))
		defer file.Close()
		file.WriteString(callbackData.PublicToken)
	}
}