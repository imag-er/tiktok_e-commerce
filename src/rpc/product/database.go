package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dsn string = "host=127.0.0.1 port=5432 user=root password=root dbname=maindb sslmode=disable"

func InitGORM() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// 自动迁移表结构
	db.AutoMigrate(&Product{})

	return db
}
