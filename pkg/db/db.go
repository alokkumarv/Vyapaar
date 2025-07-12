package db

import (
	"log"

	"github.com/alokkumarv/Vyapaar/config"
	"github.com/alokkumarv/Vyapaar/internal/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDB() *DB {
	cfg := config.GetConfig()
	db, err := gorm.Open(sqlite.Open(cfg.DB), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	return &DB{
		db: db,
	}

}

func (db *DB) Migrate() {
	db.db.AutoMigrate(&user.User{})
}
