package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=127.0.0.1 port=5432 user=root dbname=maindb sslmode=disable password=root"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// // Read
	var product Product
	// db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "price = ?",100) // 查找 code 字段值为 D42 的记录

	fmt.Print(product)

	// // Update - 将 product 的 price 更新为 200
	// db.Model(&product).Update("Price", 200)
	// // Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)
}
   