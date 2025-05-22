package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file. Error: %s", err)
	}
	log.Print(".env loaded successfully")
	return nil
}

func GetMysqlUser() string {
	return os.Getenv("MYSQL_USER")
}

func GetMysqlPwd() string {
	return os.Getenv("MYSQL_PASSWORD")
}

func GetMysqlHost() string {
	return os.Getenv("MYSQL_HOST")
}

func GetMysqlPort() string {
	return os.Getenv("MYSQL_PORT")
}

func GetDB_Name() string {
	return os.Getenv("MYSQL_DB_NAME")
}

func GetDsnRoot() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", GetMysqlUser(), GetMysqlPwd(), GetMysqlHost(), GetMysqlPort())

	return dsn
}

func GetDsnDb() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", GetMysqlUser(), GetMysqlPwd(), GetMysqlHost(), GetMysqlPort(), GetDB_Name())

	return dsn
}
