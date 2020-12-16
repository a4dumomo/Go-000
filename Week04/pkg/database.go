package pkg

import (
	"Week04/internal/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func NewDb(config *Config) {
	var err error
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: config.Db.DNS,
	}), &gorm.Config{})
	Db.AutoMigrate(&model.Article{})
	if err != nil {
		log.Fatal("db init fail,err:", err)
	}

}
