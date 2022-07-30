package dataaccess

import "github.com/jmoiron/sqlx"

type IDbConnectionFactory interface {
	GetConnection() (*sqlx.DB, error)
}
