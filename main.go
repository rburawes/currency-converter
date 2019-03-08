package main

import (
	"fmt"
	"github.com/rburawes/currency-converter/routes"
	"github.com/rburawes/currency-converter/util"
)

func main() {

	_, err := util.LoadData("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	routes.LoadRoutes()
}
