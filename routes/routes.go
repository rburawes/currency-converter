package routes

import (
	"github.com/rburawes/currency-converter/controllers"
	"net/http"
)

// LoadRoutes handles routes to pages of the application.
func LoadRoutes() {

	// Index or main page.
	http.HandleFunc("/", controllers.Index)
	// Returns the latest rates or the rates based on the passed date value.
	http.HandleFunc("/rates/", controllers.Rates)
	// Returns the currency analysis e.g. max, min and average for every available currency.
	http.HandleFunc("/rates/analyze", controllers.Analyze)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// Listens and serve requests.
	http.ListenAndServe(":8080", nil)

}
