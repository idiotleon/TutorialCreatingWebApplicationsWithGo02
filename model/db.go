package model

import (
	"database/sql"
)

var db *sql.DB

/**
 *	To inject the database instance into the model layer
 */
func SetDatabase(database *sql.DB) {
	db = database
}
