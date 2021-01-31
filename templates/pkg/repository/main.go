package repository

import (
	"database/sql"
)

// Holds all the dependencies that are being used by the database layer
type storage struct {
	RelationalDB *sql.DB
	DocumentDB   *sql.DB
}

// SetupDatabase sets up the database connections that will be used by the service
func SetupDatabase() *storage {
	return &storage{
		RelationalDB: nil,
		DocumentDB:   nil,
	}
}

// GetEmptyStorage returns a storage object with non-initialized database connections.
// It is meant to be used in test where we would setup the database connection to something like In Memory database
func GetEmptyStorage() *storage {
	return &storage{
		RelationalDB: nil,
		DocumentDB:   nil,
	}
}
