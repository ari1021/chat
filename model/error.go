package model

import (
	"github.com/go-sql-driver/mysql"
)

type APIError struct {
	StatusCode int
	Message    string
}

func IsDuplicateKeyError(err error) bool {
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		return false
	}
	if me.Number == 1062 {
		return true
	}
	return false
}

func IsForeignKeyError(err error) bool {
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		return false
	}
	if me.Number == 1452 {
		return true
	}
	return false
}
