package main

import (
	"fmt"
	"github.com/rburawes/currency-converter/controllers"
	"github.com/rburawes/currency-converter/util"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Tests the loading of data from the url to the database.
// This also test the database connection, read and write to the tables.
func TestLoadData(t *testing.T) {

	isOK, err := util.LoadData("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")

	if err != nil {
		t.Errorf("Failed to load data.")
	}

	fmt.Println("Is data loaded to the database: ", isOK)

}

// Tests the application if it will be accessible.
// Tells the application is running.
func TestAppHealth(t *testing.T) {
	executeTest("/", controllers.Index, t)
}

// Tests the Rates API with input date.
// Reutrns the rates based on the given date.
func TestRatesWithDate(t *testing.T) {
	executeTest("/rates/2019-03-01", controllers.Rates, t)
}

// Tests the Rate API that returns latest conversion rates data.
// Returns the latest rates.
func TestLatestRates(t *testing.T) {
	executeTest("/rates/latest", controllers.Rates, t)
}

// Tests the API that analyses the loaded data from the given url.
// Returns the minimum, maximum and average rates for every currency.
func TestAnalyze(t *testing.T) {
	executeTest("/rates/analyze", controllers.Analyze, t)
}

// Tests the API that returns conversion between currencies.
// Returns the value of GBP in USD
func TestConversion(t *testing.T) {
	executeTest("/convert/gbp/usd", controllers.Convert, t)
}

// Executes the test based on the given conditions.
func executeTest(url string, f func(w http.ResponseWriter, req *http.Request), t *testing.T) {
	// Create a request to pass to our a handler that manages latest rates API.
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Records the response.
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(f)

	handler.ServeHTTP(rec, req)

	// The status code must be '200'
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("request returns unexpected result with status code: RECEIVED: %v EXPECTED %v",
			status, http.StatusOK)
	}

	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Status Code: ", resp.StatusCode)
	fmt.Println("Content Type: ", resp.Header.Get("Content-Type"))
	fmt.Println("Response Body: ", string(body))

}
