package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JavaHutt/rest-dealership/internal/model"
	"github.com/JavaHutt/rest-dealership/internal/service"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
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
		r.Put("/{id}", h.updateVehicle)
		r.Delete("/{id}", h.deleteVehicle)
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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect ID"))
		return
	}

	vehicle, err := h.services.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicle)
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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Location", fmt.Sprintf("/api/v1/vehicles/%s", strconv.Itoa(int(created.ID))))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h Handler) updateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle model.Vehicle

	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect ID"))
		return
	}

	updated, err := h.services.Update(id, vehicle)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Location", fmt.Sprintf("/api/v1/vehicles/%s", strconv.Itoa(int(updated.ID))))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

func (h Handler) deleteVehicle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect ID"))
		return
	}

	if err = h.services.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
