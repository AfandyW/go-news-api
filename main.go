package main

import (
	"fmt"
	"go-news-api/config"
	tagshandler "go-news-api/handler/tags-handler"
	newsrepo "go-news-api/repository/mysql/news-repo"
	tagsrepo "go-news-api/repository/mysql/tags-repo"
	newsservice "go-news-api/service/news-service"
	tagsservice "go-news-api/service/tags-service"

	newsshandler "go-news-api/handler/news-handler"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

	// Register all repo
	tagsRepo := tagsrepo.NewRepository(db, cache)
	newsRepo := newsrepo.NewRepository(db, cache)

	// Register all service
	tagsService := tagsservice.NewService(tagsRepo)
	newsService := newsservice.NewService(newsRepo, tagsRepo)

	//Handler

	route := mux.NewRouter()

	// Test Route
	route.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte("pong"))
		return
	}).Methods("GET")

	apiroute := route.PathPrefix("/api").Subrouter()

	tagshandler.NewHandler(apiroute, tagsService)
	newsshandler.NewHandler(apiroute, newsService)

	endpoint := fmt.Sprintf("%s:%s", cfg.API.Host, cfg.API.Port)

	server := &http.Server{
		Handler: route,
		Addr:    endpoint,
	}

	fmt.Println("Server running on ", cfg.API.Host, ":", cfg.API.Port, cache, db)
	server.ListenAndServe()
}
