package tests

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func BrowserTestSetup(url string, debug bool, test func(*rod.Page)) {
	var browser *rod.Browser
	if debug == true {
		l := launcher.New().
			Headless(true)
		defer l.Cleanup()
	
		remoteUrl := l.MustLaunch()
	
		browser = rod.New().
			ControlURL(remoteUrl).
			Trace(true).
			SlowMotion(2 * time.Second).
			MustConnect()
	
		// ServeMonitor plays screenshots of each tab
		launcher.Open(browser.ServeMonitor(""))
		defer browser.MustClose()

	} else {
		browser = rod.New().MustConnect()
		defer browser.MustClose()
	}
	// Connect to main page
	page := browser.MustPage(url)
	defer page.Close()

	// Run enviroment test
	test(page)
}
