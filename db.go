package dbpattern

import (
	"database/sql"
)

// Database interface is a restrictive wrapper interface around sql.DB methods
// allows for custom use, query logging, mocking, tracing
type Database interface {
	Ping() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Close() error
}

var instance Database

// Instance returns the global Database instance
func Instance() Database { return instance }

func setInstance(db Database) {
	if instance != nil {
		instance.Close()
	}
	instance = db
}

// DB contains global db object with all the connections
// implements Database interface
type DB struct{ internal *sql.DB }

// Ping does a ping to the DB
func (db *DB) Ping() error {
	return db.internal.Ping()
}

// Close closes the DB
func (db *DB) Close() error {
	return db.internal.Close()
}

// Query does a query on DB
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// TODO: add logging
	return db.internal.Query(query, args...)
}
