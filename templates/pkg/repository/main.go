package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/pkg/errors"
)

// Holds all the dependencies that are being used by the database layer
type storage struct {
	DB *sql.DB
}

// NewStorage creates a new instance of storage struct which holds the connection to the database
func NewStorage() *storage {
	return &storage{
		DB: nil,
	}
}

func (s *storage) CloseConnection() {
	err := s.DB.Close()
	if err != nil {
		log.Fatal(err, "Issues when closing database")
	}
}

func (s *storage) InitMySQLConn(driverName, connection string) error {
	db, err := openAndPingDatabase(driverName, connection)
	if err != nil {
		return err
	}

	s.DB = db
	return nil
}

func openAndPingDatabase(driverName, connection string) (*sql.DB, error) {
	// Calling Open doesn't open a connection. It only returns the database handler
	db, err := sql.Open(driverName, connection)
	if err != nil {
		return nil, errors.Wrap(err, "issue opening database")
	}

	// Ping the database to verify the connection
	if err = db.Ping(); err != nil {
		// On first run the database is not yet initilized.
		// Wait for the database to initialize and try again
		time.Sleep(15 * time.Second)
		if err := db.Ping(); err != nil {
			return nil, errors.Wrap(err, "can't ping database")
		}
	}

	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxOpenConns(10)

	return db, nil
}
