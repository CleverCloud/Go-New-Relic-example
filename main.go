package main

import (
	"fmt"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	app := initNRAgent()
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", indexPage))
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w, `<h1>Hello, World</h1>
    <p>Go is running on Clever Cloud 💡☁️,</p>
    <p>you should give it a try too!</p>`)

}

func initNRAgent() *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Your Application Name"),
		newrelic.ConfigLicense("YOUR_NEW_RELIC_LICENSE_KEY"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil {
		fmt.Errorf("New Relic Agent is having a nervous breakdown.")
	}
	return app
}

//do anything
