package main

import "database/sql"

type Accessor interface {
	get(db *sql.DB) error
	update(db *sql.DB) error
	delete(db *sql.DB) error
	create(db *sql.DB) error
}
