package test

import (
	"fmt"
	"go-news-api/domain/entities"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
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
	err = godotenv.Load(os.ExpandEnv("../../../.env"))
	fmt.Println(122, os.Getenv("TEST_DB_HOST"))

	param := DBConfig{
		Host:     os.Getenv("TEST_DB_HOST"),
		Port:     os.Getenv("TEST_DB_PORT"),
		User:     os.Getenv("TEST_DB_USER"),
		Password: os.Getenv("TEST_DB_PASSWORD"),
		Name:     os.Getenv("TEST_DB_NAME"),
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
