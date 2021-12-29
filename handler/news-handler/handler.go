package newshandler

import (
	"context"
	"go-news-api/domain/news"
	"go-news-api/handler"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service news.IService
}

func NewHandler(r *mux.Router, nService news.IService) *Handler {
	h := &Handler{
		service: nService,
	}

	r.HandleFunc("/news", h.Create).Methods("POST")
	r.HandleFunc("/news", h.List).Methods("GET")
	r.HandleFunc("/news/{news_id}", h.Update).Methods("PUT")
	r.HandleFunc("/news/{news_id}", h.Delete).Methods("DELETE")

	return h
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var reqBody CreateNewNews

	err := handler.Decode(r, &reqBody)

	if err != nil {
		handler.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	tags, err := h.service.CreateNewNews(ctx, reqBody.Tags, reqBody.Name, reqBody.Status)

	if err != nil {
		handler.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendWithData(w, http.StatusCreated, "Success Create New Tags", tags)
	return
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	topic := r.URL.Query()["topic"]

	if topic != nil {
		news, err := h.service.ListNewsByTopic(ctx, topic[0])
		if err != nil {
			handler.SendNoData(w, http.StatusInternalServerError, err.Error())
			return
		}

		handler.SendWithData(w, http.StatusOK, "Success List News By Topic", news)
		return
	}

	status := r.URL.Query()["status"]

	if status != nil {
		news, err := h.service.ListNewsByStatus(ctx, status[0])
		if err != nil {
			handler.SendNoData(w, http.StatusInternalServerError, err.Error())
			return
		}

		handler.SendWithData(w, http.StatusOK, "Success List News By Status", news)
		return
	}

	news, err := h.service.ListNews(ctx)
	if err != nil {
		handler.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendWithData(w, http.StatusOK, "Success List News", news)
	return
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var reqBody CreateNewNews

	newsId := handler.GetParams(r, "news_id")

	err := handler.Decode(r, &reqBody)

	if err != nil {
		handler.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	news, err := h.service.UpdateNews(ctx, newsId, reqBody.Tags, reqBody.Name, reqBody.Status)

	if err != nil {
		handler.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendWithData(w, http.StatusCreated, "Success Update News", news)
	return
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	newsId := handler.GetParams(r, "news_id")

	err := h.service.DeleteNews(ctx, newsId)

	if err != nil {
		handler.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendNoData(w, http.StatusOK, "Success Delete News")
	return
}
