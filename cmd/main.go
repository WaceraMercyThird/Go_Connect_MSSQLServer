// Main applications for this project.
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

// Access configuration settings
type Config struct {
	Database struct {
		ServerName               string `json:"serverName"`
		Port                     int    `json:"port"`
		UserID                   string `json:"userId"`
		Password                 string `json:"password"`
		DatabaseName             string `json:"databaseName"`
		MultipleActiveResultSets bool   `json:"multipleActiveResultSets"`
		TrustServerCertificate   bool   `json:"trustServerCertificate"`
		MaxPoolSize              int    `json:"maxPoolSize"`
	} `json:"database"`
}

func loadConfig(filename string) (Config, error) {

	var config Config

	file, err := os.Open(filename)

	if err != nil {

		return config, err

	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)

	return config, err
}

func connectToSQLServer() (*sql.DB, error) {

	// Load configuration from app_settings.json
	config, err := loadConfig("../app_settings.json")

	if err != nil {

		return nil, err

	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		config.Database.ServerName, config.Database.UserID, config.Database.Password,
		config.Database.Port, config.Database.DatabaseName)

	db, err := sql.Open("sqlserver", connString)

	if err != nil {

		return nil, err
	}

	// Ping the SQL Server to ensure connectivity
	err = db.PingContext(context.Background())

	if err != nil {

		return nil, err
	}

	log.Println("Connected to SQL Server!")

	return db, nil
}

func main() {

	// Connect to SQL Server
	db, err := connectToSQLServer()

	if err != nil {

		log.Printf("Error connecting to SQL Server: %v", err)

		return
	}
	
	defer db.Close()

	// Perform any additional logic or database operations here

	// For example, you might want to execute queries, fetch data, etc.

	// Fetch transactionToDisplayDtos

	log.Println("Application logic executed successfully!")
}

