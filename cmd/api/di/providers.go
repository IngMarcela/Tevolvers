package di

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"os"
)

const dbHostEnvVar = "DB_HOST"

func connectionBDProvider() *gorm.DB {
	dbHost := os.Getenv(dbHostEnvVar)
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	dsn := "user:password@tcp(" + dbHost + ":3306)/employees?timeout=15s&charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return db
}
