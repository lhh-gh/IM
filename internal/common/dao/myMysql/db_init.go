package myMysql

import (
	"fmt"
	"sync"

	models2 "github/lhh-gh/IM/internal/common/dao/myMysql/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	client  *gorm.DB
	rwMutex sync.RWMutex
}

func MustNewMySQL(url string, opts ...gorm.Option) *DB {
	db, err := gorm.Open(mysql.Open(url), opts...)
	if err != nil {
		panic("Failed to connect to MySQL: " + err.Error())
	}
	fmt.Println("Connected to MySQL")
	return &DB{db, sync.RWMutex{}}
}

func (db *DB) Migrate() error {
	err := db.client.AutoMigrate(
		&models2.User{},
		&models2.UserConfig{},
		&models2.Friends{},
		&models2.Group{},
		&models2.GroupMember{},
	)
	return err
}
