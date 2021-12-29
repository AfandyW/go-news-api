package main

import (
	"context"
	"fmt"
	"go-news-api/config"
	tagshandler "go-news-api/handler/tags-handler"
	newsrepo "go-news-api/repository/mysql/news-repo"
	tagsrepo "go-news-api/repository/mysql/tags-repo"
	"go-news-api/repository/redisc"
	newsservice "go-news-api/service/news-service"
	tagsservice "go-news-api/service/tags-service"

	newsshandler "go-news-api/handler/news-handler"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
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

	rC := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(`%s:%s`, cfg.Cache.Host, cfg.Cache.Port),
		Password: "",
		DB:       int(cfg.Cache.Db),
	})

	rC.FlushDB(context.Background())

	sqlDB, err := db.DB()
	defer sqlDB.Close()

	// Migrate
	err = config.Migrate(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Register all repo
	tagsRepo := tagsrepo.NewRepository(db)
	newsRepo := newsrepo.NewRepository(db)

	// Register cache
	cacheTags := redisc.NewRedisCache(cfg.Cache)
	cacheNews := redisc.NewRedisCach(cfg.Cache)

	// Register all service
	tagsService := tagsservice.NewService(tagsRepo, cacheTags)
	newsService := newsservice.NewService(newsRepo, tagsRepo, cacheNews)

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

	fmt.Println("Server running on ", cfg.API.Host, ":", cfg.API.Port)

	server.ListenAndServe()
}
