package main

import (
	"context"
	"fmt"
	"go-news-api/config"
	news_mysql "go-news-api/repository/mysql/news"
	tags_mysql "go-news-api/repository/mysql/tags"
	"go-news-api/repository/redisc"
	"go-news-api/route"
	news_service "go-news-api/service/news"
	tags_service "go-news-api/service/tags"

	"log"

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

	redis := config.InitRedis(cfg.Cache)
	defer redis.Close()
	redis.FlushDB(context.Background())

	sqlDB, err := db.DB()
	defer sqlDB.Close()

	// Migrate
	err = config.Migrate(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Register all repo
	tagsRepo := tags_mysql.NewRepository(db)
	newsRepo := news_mysql.NewRepository(db)

	// Register cache
	cacheTags := redisc.NewRedisTagsCache(redis, cfg.Cache.Expires)
	cacheNews := redisc.NewRedisNewsCache(redis, cfg.Cache.Expires)

	// Register all service
	tagsService := tags_service.NewService(tagsRepo, cacheTags)
	newsService := news_service.NewService(newsRepo, tagsRepo, cacheNews)

	//Handler Route
	handler := route.NewHandler(newsService, tagsService)
	route := handler.NewRoute()

	endpoint := fmt.Sprintf("%s:%s", cfg.API.Host, cfg.API.Port)

	Start(route, endpoint)
}
