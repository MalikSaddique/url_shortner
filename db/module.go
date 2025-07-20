package db

import "database/sql"

type URLDB interface {
}

type URLDBImpl struct {
	Db *sql.DB
}

func NewURLDB() URLDB {
	return &URLDBImpl{}
}
