package user

import (
	"database/sql"
	"fmt"
)

type MysqlUsers struct {
	DB *sql.DB
}

func NewMysqlUsers(db *sql.DB) *MysqlUsers {
	return &MysqlUsers{DB: db}
}

func (mu *MysqlUsers) CreateUsersTable(query string) error {
	_, err := mu.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating table. Error: %s", err)
	}
	return nil
}
