package sql

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/internal/config"
	"github.com/kimxuanhong/user-manager-go/internal/infra/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Database struct {
	*gorm.DB
}

var instanceDatabase *Database
var databaseOnce sync.Once

func InitDB(cfg *config.Config) *Database {
	databaseOnce.Do(func() {
		host := cfg.Database.Host
		port := cfg.Database.Port
		user := cfg.Database.Username
		password := cfg.Database.Password
		dbname := cfg.Database.Name
		scheme := cfg.Database.Scheme

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
		fmt.Println(dsn)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		// Thiết lập schema mặc định
		db.Exec(fmt.Sprintf("SET search_path TO %s", scheme))

		// Migrate schema
		err = db.AutoMigrate(&entity.User{})
		if err != nil {
			panic("failed to create user table")
		}

		sqlDB, err := db.DB()
		if err != nil {
			panic("failed to create user table")
		}

		// Cấu hình connection pool
		sqlDB.SetMaxOpenConns(20)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)

		instanceDatabase = &Database{
			DB: db,
		}
	})
	return instanceDatabase
}
