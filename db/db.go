package db

import (
	"log"

	"github.com/ari1021/websocket/config"
	"github.com/ari1021/websocket/model"
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
	DSN := config.DSN()
	conn, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return conn, err
}

func Migrate(conn *gorm.DB) {
	if err := conn.AutoMigrate(
		&model.User{},
		&model.Room{},
		&model.Chat{},
	); err != nil {
		log.Println(err)
	}
}
