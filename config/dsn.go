package config

import (
	"fmt"
	"os"
)

// DSN(Data Source Name)をreturn
func DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	) + "?parseTime=true&collation=utf8mb4_bin"
}
