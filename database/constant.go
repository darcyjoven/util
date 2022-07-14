package database

import "github.com/jmoiron/sqlx"

type DB struct {
	db *sqlx.DB
}

type Rows struct {
	rows *sqlx.Rows
}
