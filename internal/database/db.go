// Pacakge databas provides a connection to the database
package database

import (
	"database/sql"
	"fmt"

	"synf/internal/config"
)

func Connect() (*sql.DB, error) {

	var lo config.Login
	lo.LoadCredentials()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", lo.User, lo.Pass, lo.IP, lo.Port, lo.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		print("error")
	}

	fmt.Println("Connected to database")
	return db, nil
}

func Close(db *sql.DB) {
	db.Close()
}
