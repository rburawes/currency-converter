package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/rburawes/currency-converter/models"
	"net/http"
	"strings"
	"time"
)

const (
	pathRate             = "/rates/"
	convertPath          = "/convert/"
	latestPath           = "latest"
	headerContentTypeKey = "Content-Type"
	jsonType             = "application/json"
	jsonIndexPrefix      = ""
	jsonIndentValue      = "\t"
	errorMsg             = "Unable to retrieve data: "
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

	path := r.URL.Path[len(pathRate):]

	// if the path param value is 'latest'
	var timeInput time.Time

	// date value has been provided in the path as param.
	if path != latestPath && path != "" {
		v, err := time.Parse(models.TimeFormat, path)

		if err != nil {
			http.Error(w, "Time format error encountered: "+err.Error(), http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		timeInput = v
	}

	if path == "" {
		http.Error(w, errorMsg+"invalid parameter", http.StatusInternalServerError)
		fmt.Println(errorMsg + "invalid parameter")
		return
	}

	result, err := models.Get(timeInput)

	if err != nil {
		http.Error(w, errorMsg+err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	ConvertToJSON(w, result)

}

// Convert returns the conversion between two currencies.
func Convert(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path[len(convertPath):]
	params := strings.Split(path, "/")

	if len(params) < 2 {
		http.Error(w, errorMsg+"invalid parameter", http.StatusInternalServerError)
		fmt.Println(errorMsg + "invalid parameter")
		return
	}

	result, err := models.ConvertByCurrency(strings.ToUpper(params[0]), strings.ToUpper(params[1]))

	if err != nil {
		http.Error(w, errorMsg+err.Error(), http.StatusInternalServerError)
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
		http.Error(w, errorMsg+err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	ConvertToJSON(w, result)
}

// ConvertToJSON converts the target struct to json object
func ConvertToJSON(w http.ResponseWriter, s interface{}) {

	uj, err := json.MarshalIndent(s, jsonIndexPrefix, jsonIndentValue)
	if err != nil {
		http.Error(w, errorMsg+err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set(headerContentTypeKey, jsonType)
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}
