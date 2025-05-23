package database

import (
	"database/sql"
	"fmt"
	"go-marketplace/cmd/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConn() (*sql.DB, error) {

	rootDB, err := sql.Open("mysql", config.GetDsnRoot())
	if err != nil {
		return nil, fmt.Errorf("error opening root connection: %w", err)
	}
	defer rootDB.Close()

	if err := rootDB.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging root connection: %w", err)
	}

	checkErr := existsCheck(rootDB, config.GetDB_Name())
	if checkErr == sql.ErrNoRows {
		log.Print("DATABASE NOT FOUND. CREATING....")
		if err := createDB(rootDB, config.GetDB_Name()); err != nil {
			return nil, err
		}
		log.Print("DATABASE CREATED SUCCESSFULLY")
	} else if checkErr != nil {
		return nil, checkErr
	} else {
		log.Print("DATABASE FOUND. USING DATABASE")
	}

	db, err := sql.Open("mysql", config.GetDsnDb())
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database connection: %w", err)
	}
	return db, nil
}

func existsCheck(db *sql.DB, dbName string) error {
	var exists string
	row := db.QueryRow(config.ExistsQuery, dbName).Scan(&exists)
	if row == sql.ErrNoRows {
		return sql.ErrNoRows
	}

	return nil
}

func createDB(db *sql.DB, dbName string) error {
	query := fmt.Sprintf(config.CreateDBQuery, dbName)

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating database. Error: %s", err)
	}
	return nil
}
