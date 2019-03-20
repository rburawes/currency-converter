package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	// Needed only for initialization.
	_ "github.com/lib/pq"
)

// Database is the object uses by the models for accessing
// database tables and executing queries.
var Database *sql.DB

// Config holds the connection properties for the database.
type Config struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Database   string `json:"database"`
	SslMode    string `json:"sslMode"`
	URL        string `json:"url"`
	DriverName string `json:"driverName"`
}

// Initializes database connection.
func init() {
	var err error

	config, err := loadConnectionProperties()

	if err != nil {
		panic(err)
		fmt.Println("Invalid connection details: " + err.Error())
	}

	Database, err = sql.Open(config.DriverName, config.GetConnectionString())
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the database: " + err.Error())
	}

	if err = Database.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database connection successful.")
}

// Parses the config file and load to Config struct.
func loadConnectionProperties() (Config, error) {

	var config Config
	data, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil

}

// GetConnectionString returns the data source value or connection string.
func (config Config) GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", config.Username, config.Password, config.URL, config.Database, config.SslMode)
}
