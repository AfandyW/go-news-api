package route

import (
	"go-news-api/domain/news"
	"go-news-api/domain/tags"
	news_handler "go-news-api/route/news"
	tags_handler "go-news-api/route/tags"

	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	NewsService news.IService
	TagsService tags.IService
}

func NewHandler(newsService news.IService, tagsService tags.IService) *Handler {
	return &Handler{
		NewsService: newsService,
		TagsService: tagsService,
	}
}

func (h Handler) NewRoute() *mux.Router {
	//Handler
	route := mux.NewRouter()

	// Test Route
	route.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte("pong"))
		return
	}).Methods("GET")

	apiroute := route.PathPrefix("/api").Subrouter()

	tags_handler.NewHandler(apiroute, h.TagsService)
	news_handler.NewHandler(apiroute, h.NewsService)

	return route
}
