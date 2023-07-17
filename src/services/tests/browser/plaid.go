package tests

import (
	"github.com/go-rod/rod"
)

func TestPlaidWorkFlow(page *rod.Page){
	// Select iframe context
	iframe := page.MustSearch("iframe#plaid-link-iframe-1").MustFrame()
	defer iframe.Close()

	// Complete plaid account link workflow
	user := "user_good"
	pass := "pass_good"
	iframe.MustElement("button#aut-button").MustClick()
	iframe.MustElement("button#aut-ins_33").MustClick()
	iframe.MustElement("input#aut-input-0").MustInput(user)
	iframe.MustElement("input#aut-input-1").MustInput(pass)
	iframe.MustElement("button#aut-button").MustClick()
	iframe.MustElement("button#aut-button").MustClick()
	iframe.MustElement("button#aut-button").MustClick()
	iframe.MustWaitStable()
}