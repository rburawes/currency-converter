package util

import (
	"encoding/xml"
	"fmt"
	"github.com/rburawes/currency-converter/models"
	"io"
	"net/http"
)

// LoadData reads xml data from the url.
func LoadData(url string) (bool, error) {

	resp, err := http.Get(url)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unable to retrieve data from the given url %d", resp.StatusCode)
	}

	data, err := readCubes(resp.Body)

	if err != nil {
		return false, fmt.Errorf("an error has occurred while reading data: %v", err)
	}

	isSaved, err := models.SaveCube(&data)

	if err != nil {
		return isSaved, fmt.Errorf("an error has occurred while saving the data: %v", err)
	}

	return isSaved, nil

}

// Reads the content of the http response body to get the xml data.
func readCubes(reader io.Reader) (models.Data, error) {
	var data models.Data
	if err := xml.NewDecoder(reader).Decode(&data); err != nil {
		return data, fmt.Errorf("something went wrong while processing the xml data: %v", err)
	}
	return data, nil
}
