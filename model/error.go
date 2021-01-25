package model

import (
	"errors"
	"net/http"

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

func checkMySQLError(err error) error {
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

func NewAPIResponse(err error) (statusCode int, res *APIError) {
	switch checkMySQLError(err) {
	case ForeignKeyError:
		res := NewAPIError(400, "foreign key error")
		return http.StatusBadRequest, res
	case DuplicateKeyError:
		res := NewAPIError(400, "duplicate key error")
		return http.StatusBadRequest, res
	case DatabaseError:
		res := NewAPIError(500, "database error")
		return http.StatusInternalServerError, res
	default:
		res := NewAPIError(500, "unknown error")
		return http.StatusInternalServerError, res
	}
}
