package db

import (
	"database/sql"
	"fmt"
)

type DBTable struct {
	db        *sql.DB
	tableName string
}

func NewDBInstance(dbName string) *DBTable {
	fmt.Println("NewDBInstance", dbName)
	return &DBTable{
		db:        nil,
		tableName: dbName,
	}
}
