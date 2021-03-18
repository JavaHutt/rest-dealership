package handler

import (
	"net/http"

	"github.com/JavaHutt/rest-dealership/internal/service"
	"github.com/go-chi/chi"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	return r
}
