package handler

import (
	"fmt"
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
	r.Route("/api/v1/vehicles", func(r chi.Router) {
		r.Get("/", h.getAllVehicles)
		r.Get("/{id}", h.getVehicleById)
		r.Post("/", h.createVehicle)
		r.Put("/", h.updateVehicle)
		r.Delete("/", h.deleteVehicle)
	})
	return r
}

func (h Handler) getAllVehicles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("getAll method"))
}

func (h Handler) getVehicleById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res := fmt.Sprintf("get by id: %s", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (h Handler) createVehicle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("create"))
}

func (h Handler) updateVehicle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("update"))
}

func (h Handler) deleteVehicle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("delete"))
}
