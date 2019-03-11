package models

import (
	"database/sql"
	"github.com/rburawes/currency-converter/config"
	"log"
	"strconv"
	"time"
)

// The number of goroutines to use in inserting records.
// More goroutines the faster the insert execution.
const (
	TimeFormat  = "2006-01-02"
	defaultBase = "EUR"
)

var entries []CubeEntry

// Cube holds data for the currency and rate.
type cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}

// cubeCollection the parent of {cube} and contains date and time.
type cubeCollection struct {
	Time string `xml:"time,attr"`
	Cube []cube `xml:"Cube"`
}

// cubeData a list of {cubeCollection}
type cubeData struct {
	CubeCollection []cubeCollection `xml:"Cube"`
}

// dataSender holds the info about source of the data.
type dataSender struct {
	Name string `xml:"name"`
}

// Data wraps everything and holds all the info.
type Data struct {
	Subject  string     `xml:"subject"`
	Sender   dataSender `xml:"Sender"`
	CubeData cubeData   `xml:"Cube"`
}

// CubeEntry converts the read data from the url to a format easy
// for database processes.
type CubeEntry struct {
	Currency string
	Rate     float32
	RateTime time.Time
}

// CubeResult an object that will hold result data from db query
// and will be shown in the API as JSON.
type CubeResult struct {
	Base  string
	Rates map[string]float32
}

// RateAnalysis holds the data for minimum, maximum and average rates
// for currencies.
type RateAnalysis struct {
	Minimum float32
	Maximum float32
	Average float32
}

// AnalysisData is the wrapper for analysis data.
type AnalysisData struct {
	Base  string
	Rates map[string]RateAnalysis
}

// SaveCube saves the cube data to the database.
func SaveCube(c *Data) (bool, error) {

	for _, data := range c.CubeData.CubeCollection {

		t, err := time.Parse(TimeFormat, data.Time)

		if err != nil {
			return false, err
		}

		for _, cube := range data.Cube {
			entry := CubeEntry{}
			entry.RateTime = t
			entry.Currency = cube.Currency

			if v, err := strconv.ParseFloat(cube.Rate, 32); err == nil {
				entry.Rate = float32(v)
			}
			entries = append(entries, entry)
		}

	}

	// the statement to use in inserting a record
	// ignore data if found to be a duplicate.
	stmt := "INSERT INTO cube (currency, rate, rate_time) VALUES ($1, $2, $3) ON CONFLICT ON CONSTRAINT cube_crr_index DO NOTHING"

	// use goroutines to save record to the database.
	// multiple qoroutines can be used to process huge dataset
	go inserter(stmt)

	return true, nil
}

// A function to be executed as a new goroutine for
// executing database insert.
func inserter(statement string) {

	for _, entry := range entries {
		stmt, err := config.Database.Prepare(statement)
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(entry.Currency, entry.Rate, entry.RateTime)

		if err != nil {
			log.Fatal(err)
		}

		stmt.Close()

	}

}

// Get the conversion rates based on the given date.
// If given date is null then get the latest.
func Get(targetTime time.Time) (CubeResult, error) {

	var inputTime time.Time
	rows := &sql.Rows{}
	var err error

	if !targetTime.IsZero() {
		inputTime = targetTime
		rows, err = config.Database.Query("SELECT DISTINCT c.currency, c.rate FROM cube c WHERE c.rate_time = $1 ORDER BY c.rate ASC", inputTime)
	} else {
		rows, err = config.Database.Query("SELECT DISTINCT c.currency, c.rate FROM cube c WHERE c.rate_time IN (SELECT max(rate_time) rt FROM cube GROUP BY currency) ORDER BY c.currency ASC ")
	}

	if err != nil {
		return CubeResult{}, err
	}

	defer rows.Close()

	cr := CubeResult{}
	cr.Base = defaultBase
	dataMap := make(map[string]float32)
	for rows.Next() {
		var key string
		var value float32
		err := rows.Scan(&key, &value)
		if err != nil {
			return CubeResult{}, err
		}
		dataMap[key] = value
	}

	cr.Rates = dataMap

	return cr, nil

}

// Count finds records based on the current date.
func Count() int64 {

	var count int64
	row := config.Database.QueryRow("SELECT COUNT(*) FROM cube c WHERE c.rate_time = current_date")
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count

}

// AnalyseData provides max, min and average currency rates from the data loaded into the database.
func AnalyseData() (AnalysisData, error) {

	ad := AnalysisData{}
	rows, err := config.Database.Query("SELECT currency, max(rate), min(rate), avg(rate) FROM cube GROUP BY currency ORDER BY currency ASC")

	if err != nil {
		return ad, err
	}

	defer rows.Close()

	dataMap := make(map[string]RateAnalysis)
	for rows.Next() {
		var curr string
		ra := RateAnalysis{}
		err := rows.Scan(&curr, &ra.Maximum, &ra.Minimum, &ra.Average)
		if err != nil {
			return ad, err
		}

		dataMap[curr] = ra
	}

	ad.Base = defaultBase
	ad.Rates = dataMap

	return ad, nil
}
