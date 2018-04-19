package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// GetUser returns the user for a given device.
func GetUser(stmt *sql.Stmt, device string) (string, error) {
	var result string
	err := stmt.QueryRow(device).Scan(&result)
	if err != nil {
		return "", err
	}
	return result, nil
}

// SetUser sets the user for a given device.
func SetUser(stmt *sql.Stmt, device, user string) error {
	err := exec(stmt, user, device)
	if err != nil {
		return err
	}
	return nil
}

// SetSW sets the software for a given device.
func SetSW(stmt *sql.Stmt, device, sw string) error {
	err := exec(stmt, sw, device)
	if err != nil {
		return err
	}
	return nil
}

// execAffectingOneRow executes a given statement, expecting one row to be affected.
func exec(stmt *sql.Stmt, args ...interface{}) error {
	_, err := stmt.Exec(args...)
	if err != nil {
		return fmt.Errorf("mysql: could not execute statement: %v", err)
	}
	return nil
}
