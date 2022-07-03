package dataaccess

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
)

type DbConnectionFactory struct {
	User     string
	Password string
	DbName   string
	SSLMode  bool
}

func GetConnectionFactory() IDbConnectionFactory {
	once.Do(func() {
		factory = &DbConnectionFactory{
			User:     "postgres",
			Password: "admin",
			DbName:   "pubotainment",
			SSLMode:  false,
		}
	})
	return factory
}

var factory *DbConnectionFactory
var once sync.Once

func (factory *DbConnectionFactory) GetConnection() (*sqlx.DB, error) {

	var sslMode string
	if factory.SSLMode {
		sslMode = "enable"
	} else {
		sslMode = "disable"
	}
	connectionString := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v", factory.User, factory.Password, factory.DbName, sslMode)

	// TODO: use connection pooling
	db, err := sqlx.Connect("postgres", connectionString)
	return db, err
}
