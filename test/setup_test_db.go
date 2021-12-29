package test

import (
	"fmt"
	"go-news-api/domain/entities"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB
	syncDBOnce sync.Once
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func NewDB() (*gorm.DB, error) {
	var err error
	// err = godotenv.Load()

	param := DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "root",
		Name:     "go_news_test",
	}

	syncDBOnce.Do(func() {
		if db == nil {
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
				param.User, param.Password, param.Host, param.Port, param.Name)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

			if err != nil {
				return
			}
			sqlDb, _ := db.DB()
			err = sqlDb.Ping()

			if err != nil {
				return
			}

			sqlDb.SetMaxIdleConns(10)
			sqlDb.SetMaxOpenConns(100)
			sqlDb.SetConnMaxLifetime(time.Hour)
		}
	})
	Migrate(db)
	return db, err
}


func Migrate(db *gorm.DB) error {
	err := db.Migrator().DropTable(&entities.News{}, &entities.Tags{})
	err = db.AutoMigrate(&entities.News{}, &entities.Tags{})

	return err
}
