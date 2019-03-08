# Currency converter

The application demonstrates the following:

1. Reads the content of the XML file from the web.
2. Loads the xml content into a database.
3. Provide an endpoint to display the loaded data from the database.

# The application flow
1. Application starts with the file 'main.go'
2. First it calls 'LoadData' (util/data_loader.go) function from the 'util package' to read the xml file from the given url and save it to the database.
3. 'LoadData' uses SaveData (see models/cube.go for more details) function from the models package to save the read data to the database.
4. All the database related operations are under 'models/cube.go'
5. Once the data is loaded. 'main.go' calls 'LoadRoutes' from 'routes' package to prepare the application for handling request and response.
6. 'routes/routes.go' defines all the paths that the application can serve and accept.
7. 'controllers/web.go' handles the processing of the web request and response. 

# Running the app

Assuming the you have already installed Go on your local machine and cloned this project.

1. Create a database using the sql file under 'sql/cube.sql'.  This project uses PostgreSQL, the default user and password used were 'postgres/postgres'.
2. Database configuration is under 'config/db.go'.  This provides db connection the moment that the application runs.
3. Go to 'currency-converter' directory and execute 'go run main.go'
4. Open 'http://localhost:8080', a message saying that the application is running will be displayed.
5. 'http://localhost:8080/rates/latest' returns the latest currency conversion rates.
6. 'http://localhost:8080/rates/2019-03-04' returns the currency conversion rates the given date.
7. 'http://localhost:8080/rates/analyze' returns the max, min and average conversion rates for every available currency.

# Unit tests

1. All the unit tests are in the root directory of this project, every file with '_test' suffix.
2. To execute the unit tests, go to the root directory via terminal or command line and just type 'go test -v' and hit enter.
3. See sample for passing tests below:

```
go test -v
Database connection successful.
=== RUN   TestLoadData
Is data loaded to the database:  true
--- PASS: TestLoadData (2.55s)
=== RUN   TestAppHealth
Status Code:  200
Content Type:  text/plain; charset=utf-8
Response Body:  I am running...

--- PASS: TestAppHealth (0.00s)
=== RUN   TestRatesWithDate
Status Code:  200
Content Type:  application/json
Response Body:  {
        "Base": "EUR",
        "Rates": {
                "AUD  ": 1.6,
                "BGN  ": 1.9558,
                "BRL  ": 4.3037,
                "CAD  ": 1.4971,
                "CHF  ": 1.1363,
                "CNY  ": 7.6332,
                "CZK  ": 25.636,
                "DKK  ": 7.4613,
                "GBP  ": 0.85968,
                "HKD  ": 8.9344,
                "HRK  ": 7.432,
                "HUF  ": 316.06,
                "IDR  ": 16067,
                "ILS  ": 4.1331,
                "INR  ": 80.695,
                "ISK  ": 135.9,
                "JPY  ": 127.35,
                "KRW  ": 1282.12,
                "MXN  ": 21.994,
                "MYR  ": 4.6374,
                "NOK  ": 9.7268,
                "NZD  ": 1.6656,
                "PHP  ": 58.986,
                "PLN  ": 4.3096,
                "RON  ": 4.7431,
                "RUB  ": 74.9928,
                "SEK  ": 10.5003,
                "SGD  ": 1.5396,
                "THB  ": 36.113,
                "TRY  ": 6.123,
                "USD  ": 1.1383,
                "ZAR  ": 16.1426
        }
}

--- PASS: TestRatesWithDate (0.01s)
=== RUN   TestLatestRates
Status Code:  200
Content Type:  application/json
Response Body:  {
        "Base": "EUR",
        "Rates": {
                "AUD  ": 1.6014,
                "BGN  ": 1.9558,
                "BRL  ": 4.3234,
                "CAD  ": 1.5131,
                "CHF  ": 1.1355,
                "CNY  ": 7.5622,
                "CZK  ": 25.61,
                "DKK  ": 7.461,
                "GBP  ": 0.8588,
                "HKD  ": 8.8476,
                "HRK  ": 7.4158,
                "HUF  ": 315.36,
                "IDR  ": 15990.17,
                "ILS  ": 4.0782,
                "INR  ": 78.884,
                "ISK  ": 136.8,
                "JPY  ": 125.97,
                "KRW  ": 1273.15,
                "MXN  ": 21.8244,
                "MYR  ": 4.6065,
                "NOK  ": 9.786,
                "NZD  ": 1.6631,
                "PHP  ": 58.923,
                "PLN  ": 4.2991,
                "RON  ": 4.7415,
                "RUB  ": 74.3115,
                "SEK  ": 10.5625,
                "SGD  ": 1.5308,
                "THB  ": 35.819,
                "TRY  ": 6.1171,
                "USD  ": 1.1271,
                "ZAR  ": 16.1514
        }
}

--- PASS: TestLatestRates (0.01s)
=== RUN   TestAnalyze
Status Code:  200
Content Type:  application/json
Response Body:  {
        "Base": "EUR",
        "Rates": {
                "AUD  ": {
                        "Minimum": 1.5564,
                        "Maximum": 1.6314,
                        "Average": 1.5985427
                },
                "BGN  ": {
                        "Minimum": 1.9558,
                        "Maximum": 1.9558,
                        "Average": 1.9558
                },
                "BRL  ": {
                        "Minimum": 4.1231,
                        "Maximum": 4.8942,
                        "Average": 4.398165
                },
                "CAD  ": {
                        "Minimum": 1.4795,
                        "Maximum": 1.5605,
                        "Average": 1.5101217
                },
                "CHF  ": {
                        "Minimum": 1.1217,
                        "Maximum": 1.147,
                        "Average": 1.1343775
                },
                "CNY  ": {
                        "Minimum": 7.5622,
                        "Maximum": 8.0958,
                        "Average": 7.840493
                },
                "CZK  ": {
                        "Minimum": 25.434,
                        "Maximum": 26.032,
                        "Average": 25.749062
                },
                "DKK  ": {
                        "Minimum": 7.4536,
                        "Maximum": 7.4679,
                        "Average": 7.461886
                },
                "GBP  ": {
                        "Minimum": 0.85503,
                        "Maximum": 0.9068,
                        "Average": 0.884796
                },
                "HKD  ": {
                        "Minimum": 8.8161,
                        "Maximum": 9.2313,
                        "Average": 8.976299
                },
                "HRK  ": {
                        "Minimum": 7.387,
                        "Maximum": 7.4405,
                        "Average": 7.423148
                },
                "HUF  ": {
                        "Minimum": 315.36,
                        "Maximum": 327.81,
                        "Average": 321.71094
                },
                "IDR  ": {
                        "Minimum": 15844.69,
                        "Maximum": 17634.51,
                        "Average": 16693.254
                },
                "ILS  ": {
                        "Minimum": 4.0782,
                        "Maximum": 4.315,
                        "Average": 4.1947303
                },
                "INR  ": {
                        "Minimum": 78.884,
                        "Maximum": 85.7615,
                        "Average": 82.17948
                },
                "ISK  ": {
                        "Minimum": 124.6,
                        "Maximum": 141.2,
                        "Average": 134.9969
                },
                "JPY  ": {
                        "Minimum": 122.21,
                        "Maximum": 132.78,
                        "Average": 127.7252
                },
                "KRW  ": {
                        "Minimum": 1264.39,
                        "Maximum": 1320.7,
                        "Average": 1287.8138
                },
                "MXN  ": {
                        "Minimum": 21.4546,
                        "Maximum": 23.3643,
                        "Average": 22.221813
                },
                "MYR  ": {
                        "Minimum": 4.5928,
                        "Maximum": 4.8694,
                        "Average": 4.736449
                },
                "NOK  ": {
                        "Minimum": 9.4198,
                        "Maximum": 10.0025,
                        "Average": 9.672264
                },
                "NZD  ": {
                        "Minimum": 1.6398,
                        "Maximum": 1.7852,
                        "Average": 1.7065984
                },
                "PHP  ": {
                        "Minimum": 58.726,
                        "Maximum": 63.999,
                        "Average": 60.750675
                },
                "PLN  ": {
                        "Minimum": 4.271,
                        "Maximum": 4.3445,
                        "Average": 4.300973
                },
                "RON  ": {
                        "Minimum": 4.6308,
                        "Maximum": 4.7722,
                        "Average": 4.682694
                },
                "RUB  ": {
                        "Minimum": 74.0368,
                        "Maximum": 81.2688,
                        "Average": 76.32421
                },
                "SEK  ": {
                        "Minimum": 10.1753,
                        "Maximum": 10.6923,
                        "Average": 10.381704
                },
                "SGD  ": {
                        "Minimum": 1.5293,
                        "Maximum": 1.6083,
                        "Average": 1.5657116
                },
                "THB  ": {
                        "Minimum": 35.21,
                        "Maximum": 38.239,
                        "Average": 37.030945
                },
                "TRY  ": {
                        "Minimum": 5.9322,
                        "Maximum": 7.856,
                        "Average": 6.450107
                },
                "USD  ": {
                        "Minimum": 1.126,
                        "Maximum": 1.1777,
                        "Average": 1.1452489
                },
                "ZAR  ": {
                        "Minimum": 15.242,
                        "Maximum": 17.9906,
                        "Average": 16.315477
                }
        }
}

--- PASS: TestAnalyze (0.02s)
PASS
ok      github.com/rburawes/currency-converter  2.605s


```