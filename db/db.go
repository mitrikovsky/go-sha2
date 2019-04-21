package db

import (
	"../config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/op/go-logging"
)

var (
	db  *sql.DB
	log = logging.MustGetLogger("db")
)

func init() {
	var err error
	// connect to DB
	params := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	db, err = sql.Open("postgres", params)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}

	// check connection
	if err = db.Ping(); err != nil {
		log.Fatalf("db.Ping: %v", err)
	}

	log.Info("DB ok")
}

// close db connection
func Close() {
	db.Close()
}
