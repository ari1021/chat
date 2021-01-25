package model

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

var (
	DuplicateKeyError = errors.New("Duplicate key error")
	ForeignKeyError   = errors.New("Foreign key error")
	DatabaseError     = errors.New("Database error")
	UnknownError      = errors.New("Unknown error")
)

type APIError struct {
	StatusCode int
	Message    string
}

func NewAPIError(statusCode int, message string) *APIError {
	res := &APIError{
		StatusCode: statusCode,
		Message:    message,
	}
	return res
}

func CheckMySQLError(err error) error {
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		return UnknownError
	} else {
		switch me.Number {
		case 1062:
			return DuplicateKeyError
		case 1452:
			return ForeignKeyError
		default:
			return DatabaseError
		}
	}
}
