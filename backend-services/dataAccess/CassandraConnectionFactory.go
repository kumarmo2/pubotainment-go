package dataaccess

import (
	"errors"
	"log"
	"sync"

	"github.com/gocql/gocql"
)

var cassandraFactory cassandraConnectionFactory
var cassandraInitError error
var casOnce sync.Once

func GetCassandraConnectionFactory() ICassandraConnectionFactory {
	return &cassandraFactory
}

type ICassandraConnectionFactory interface {
	GetConnection() (*gocql.Session, error)
}

type cassandraConnectionFactory struct {
	session *gocql.Session
}

func initConnection() {
	// NOTE: always use the IP address to connect to the cassandra cluster
	// and not the domain.
	log.Println("========= initializing cassandra connection ========")
	config := gocql.NewCluster("127.0.0.1:9042")
	if config == nil {
		cassandraInitError = errors.New("Could not create cluster config")
		return
	}

	config.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandr1",
	}
	config.Consistency = gocql.One

	session, err := config.CreateSession() // this session is thread-safe.
	if err != nil {
		cassandraInitError = err
	}
	cassandraFactory.session = session
}

func (factory *cassandraConnectionFactory) GetConnection() (*gocql.Session, error) {
	casOnce.Do(initConnection)
	return cassandraFactory.session, cassandraInitError
}
