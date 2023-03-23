package configs

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewDB() *gorm.DB {
	host := "localhost"
	user := "postgres"
	password := "root"
	dbName := "belajar_golang_restful_api"
	port := "5432"
	sslmode := "disable"
	timeZone := "UTC"
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host,
		user,
		password,
		dbName,
		port,
		sslmode,
		timeZone,
	)

	file, err2 := os.Create("gorm-log.txt")
	if err2 != nil {
		panic(err2)
	}

	newLogger := logger.New(
		log.New(file, "", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	return db
}
