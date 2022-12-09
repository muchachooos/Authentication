package storage

import "github.com/jmoiron/sqlx"

type UserStorage struct {
	DataBase *sqlx.DB
}
