package internal

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"snapshop/models"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func DBSession() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		// Migrate the schema
		if err := db.AutoMigrate(&models.DBProduct{}); err != nil {
			log.Fatalf("autoMigrate failed: %v", err)
		}
	})
	return db, err
}

func ToDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.DBProduct{})
	if err != nil {
		return nil, fmt.Errorf("autoMigrate failed: %v", err)
	}
	return db, nil

	// Create
	//db.Create(&models.DBProduct{Code: "D42", Price: 100})

	// Read
	//var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)
}
