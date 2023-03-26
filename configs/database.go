package configs

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"time"
)

func dsn(dbName string) string {
	host := "localhost"
	user := "postgres"
	password := "root"
	port := "5432"
	sslmode := "disable"
	timeZone := "UTC"
	dsnResult := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host,
		user,
		password,
		dbName,
		port,
		sslmode,
		timeZone,
	)

	return dsnResult
}

func NewDB() *gorm.DB {
	file, err := os.Create("gorm-log.txt")
	if err != nil {
		panic(err)
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

	db, err := gorm.Open(postgres.Open(dsn("belajar_golang_restful_api")), &gorm.Config{
		Logger: newLogger,
	})

	err = db.Use(
		dbresolver.
			Register(
				dbresolver.Config{
					Sources: []gorm.Dialector{postgres.Open(dsn("dvdrental"))},
				},
				//"film", &domain.Film{}, "secondary",
				"film", "actor",
			).
			SetMaxIdleConns(10).
			SetMaxOpenConns(100),
	)
	if err != nil {
		panic(err)
	}

	return db
}
