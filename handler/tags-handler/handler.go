package tagshandler

import (
	"context"
	"go-news-api/domain/entities"
	"go-news-api/domain/tags"
	"go-news-api/handler"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service tags.IService
}

func NewHandler(r *mux.Router, serv tags.IService) *Handler {
	h := &Handler{
		service: serv,
	}

	r.HandleFunc("/tags", h.Create).Methods("POST")
	r.HandleFunc("/tags", h.List).Methods("GET")
	r.HandleFunc("/tags/{tags_id}", h.Update).Methods("PUT")
	r.HandleFunc("/tags/{tags_id}", h.Delete).Methods("DELETE")

	return h
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var reqBody CreateNewTags

	err := handler.Decode(r, &reqBody)

	if err != nil {
		handler.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	payload := &entities.Tags{
		Name: reqBody.Name,
	}

	tags, err := h.service.CreateNewTags(ctx, payload)

	if err != nil {
		handler.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendWithData(w, http.StatusCreated, "Success Create New Tags", tags)
	return
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	tags, err := h.service.ListTags(ctx)
	if err != nil {
		handler.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendWithData(w, http.StatusOK, "Success List Tags", tags)
	return
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var reqBody CreateNewTags

	tagsId := handler.GetParams(r, "tags_id")

	err := handler.Decode(r, &reqBody)

	if err != nil {
		handler.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	tags, err := h.service.UpdateTags(ctx, tagsId, reqBody.Name)

	if err != nil {
		handler.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendWithData(w, http.StatusCreated, "Success Update Tags", tags)
	return
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	tagsId := handler.GetParams(r, "tags_id")

	err := h.service.DeleteTags(ctx, tagsId)

	if err != nil {
		handler.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendNoData(w, http.StatusOK, "Success Delete Tags")
	return
}
