package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/rburawes/currency-converter/models"
	"net/http"
	"time"
)

const (
	pathRate             = "/rates/"
	latestPath           = "latest"
	headerContentTypeKey = "Content-Type"
	jsonType             = "application/json"
	jsonIndexPrefix      = ""
	jsonIndentValue      = "\t"
)

// Index is the default page of the application.
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "I am running...\n")
}

// Rates show the latest conversion rates.
// If there is date parameter in the path, use the value to search for the records.
func Rates(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	param := r.URL.Path[len(pathRate):]

	// if the path param value is 'latest'
	var timeInput time.Time

	// date value has been provided in the path as param.
	if param != latestPath && param != "" {
		v, err := time.Parse(models.TimeFormat, param)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		timeInput = v
	}

	result, err := models.Get(timeInput)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	ConvertToJSON(w, result)

}

// Analyze shows analyzed currency rates from the loaded data into the database.
func Analyze(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	result, err := models.AnalyseData()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	ConvertToJSON(w, result)
}

// ConvertToJSON converts the target struct to json object
func ConvertToJSON(w http.ResponseWriter, s interface{}) {

	uj, err := json.MarshalIndent(s, jsonIndexPrefix, jsonIndentValue)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set(headerContentTypeKey, jsonType)
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}
