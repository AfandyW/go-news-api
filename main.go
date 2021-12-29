package main

import (
	"fmt"
	"go-news-api/config"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load configuration")
	}

	db, err := config.InitDB(cfg.DB)

	if err != nil {
		log.Fatal(err.Error())
	}

	sqlDB, err := db.DB()
	defer sqlDB.Close()

	cache := config.InitClient(cfg.Cache)

	// Migrate
	err = config.Migrate(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Handler

	endpoint := fmt.Sprintf("%s:%s", cfg.API.Host, cfg.API.Port)

	server := &http.Server{
		Addr: endpoint,
	}

	fmt.Println("Server running on ", cfg.API.Host, ":", cfg.API.Port, cache, db)
	server.ListenAndServe()
}
