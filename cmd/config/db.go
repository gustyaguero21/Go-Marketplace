package config

// database queries
const (
	ExistsQuery   = "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME=?"
	CreateDBQuery = "CREATE DATABASE IF NOT EXISTS %s"
)

// mysql queries
const (
	UsersTableQuery = "CREATE TABLE users (id VARCHAR(36) PRIMARY KEY,name VARCHAR(100),last_name VARCHAR(100),username VARCHAR(50),email VARCHAR(100),password VARCHAR(100))"
)
