package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, "127.0.0.1", 3306, "atp_db")

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
