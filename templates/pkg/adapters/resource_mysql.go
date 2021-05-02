package adapters

import (
	"database/sql"
	"log"
	"module_placeholder/pkg/app/query"
	domain "module_placeholder/pkg/domain/resource"
	"time"

	"github.com/pkg/errors"
)

// The struct should be renamed to present the identified domain model
type MySqlResourceRepository struct {
	db *sql.DB
}

func NewMySqlResourceRepository(db *sql.DB) *MySqlResourceRepository {
	if db == nil {
		panic("missing sql.DB object")
	}

	return &MySqlResourceRepository{
		db: db,
	}
}

func (repo MySqlResourceRepository) GetResource(ID uint64) (*query.Resource, error) {
	return nil, nil
}

func (repo MySqlResourceRepository) AllResources() ([]*query.Resource, error) {
	return nil, nil
}

func (repo MySqlResourceRepository) CreateResource(obj *domain.Resource) (uint64, error) {
	return 0, nil
}

func (repo MySqlResourceRepository) UpdateResource(obj *domain.Resource) error {
	return nil
}

func (repo MySqlResourceRepository) DeleteResource(ID uint64) error {
	return nil
}

func InitMySQLConn(driverName, connection string) (*sql.DB, error) {
	db, err := openAndPingDatabase(driverName, connection)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err, "Issues when closing database")
	}
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
