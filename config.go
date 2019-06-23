package dbpattern

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq" // import db driver
)

// PostgresConfig contains connection information for Postgres
type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
	MaxConn  int
}

// Config is global config var
var Config PostgresConfig

func (c PostgresConfig) dsn() string {
	hostString := fmt.Sprintf("user=%s password=%s host=%s port=%d database=%s", c.Username, c.Password, c.Host, c.Port, c.DBName)
	return hostString
}

// open sets the global db instance to a Postgres db
func (c PostgresConfig) open() error {
	db, err := sql.Open("postgres", c.dsn())
	if err != nil {
		return err
	}
	if db == nil {
		return errors.New("db is nil")
	}

	db.SetMaxOpenConns(c.MaxConn)
	setInstance(&DB{db})

	return nil
}

// Connect opens connection to the database and does a Ping to test the connection
func Connect() error {
	err := Config.open()
	if err != nil {
		return fmt.Errorf("db connect error: %v", err)
	}

	err = instance.Ping()
	if err != nil {
		return fmt.Errorf("db connect error: %v", err)
	}

	return nil
}
