package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
)

var db *sql.DB

func Connect() (*sql.DB, error) {
	user := os.Getenv("DBUSER")
	if user == "" {
		return nil, fmt.Errorf("DBUSER environment variable not set")
	}

	pass := os.Getenv("DBPASS")
	if pass == "" {
		return nil, fmt.Errorf("DBPASS environment variable not set")
	}

	dbHost := os.Getenv("DBHOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}

	dbPort := 3306
	dbPortStr := os.Getenv("DBPORT")
	if dbPortStr != "" {
		if port, err := strconv.ParseInt(dbPortStr, 10, 64); err == nil {
			dbPort = int(port)
		}
	}

	dbName := os.Getenv("DBNAME")
	if dbName == "" {
		dbName = "atp_db"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, dbHost, dbPort, dbName)
	fmt.Printf("Connecting to database: %s\n", dsn) //TODO: remove this line

	// initialize the connection pool
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// set the maximum number of connections in the pool
	db.SetMaxOpenConns(100)

	return db, nil

}
