package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connection情報を保持するstruct
type DBConnection struct {
	Conn *gorm.DB
}

var DB = DBConnection{}

// connectionを返すmethod
func (db DBConnection) GetConnection() *gorm.DB {
	return db.Conn
}

// connectionを生成するfunction
func NewConnection() (*gorm.DB, error) {
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp"
	MYSQL_ADDRESS := "127.0.0.1"
	MYSQL_PORT := "3306"
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	dsn := MYSQL_USER + ":" + MYSQL_PASSWORD + "@" + PROTOCOL + "(" + MYSQL_ADDRESS + ":" + MYSQL_PORT + ")/" + MYSQL_DATABASE
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return conn, err
}
