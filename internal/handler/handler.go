package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JavaHutt/rest-dealership/internal/model"
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
	vehicles := h.services.GetAll()

	res, err := json.Marshal(&vehicles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to unmarshal JSON"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h Handler) getVehicleById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res := fmt.Sprintf("get by id: %s", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (h Handler) createVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle model.Vehicle

	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	created, err := h.services.Create(vehicle)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := json.Marshal(created)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to unmarshal JSON"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (h Handler) updateVehicle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("update"))
}

func (h Handler) deleteVehicle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("delete"))
}
