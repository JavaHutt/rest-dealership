package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JavaHutt/rest-dealership/internal/model"
	"github.com/JavaHutt/rest-dealership/internal/service"
	mock_service "github.com/JavaHutt/rest-dealership/internal/service/mocks"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_createVehicle(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockVehicle, vehicle model.Vehicle)

	testTable := []struct {
		name                 string
		inputVehicle         model.Vehicle
		inputBody            string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			inputVehicle: model.Vehicle{
				Brand:        "Subaru",
				VehicleModel: "Impreza",
				Price:        22500,
				Status:       1,
				Mileage:      0,
			},
			inputBody: `{
				"brand": "Subaru",
				"vehicle_model": "Impreza",
				"price": 22500,
				"status": 1,
				"mileage": 0
			}`,
			mockBehaviour: func(s *mock_service.MockVehicle, vehicle model.Vehicle) {
				s.EXPECT().Create(vehicle).Return(&model.Vehicle{
					Brand:        "Subaru",
					VehicleModel: "Impreza",
					Price:        22500,
					Status:       1,
					Mileage:      0,
				}, nil)
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponseBody: `{
				"id": 0,
				"brand": "Subaru",
				"vehicle_model": "Impreza",
				"price": 22500,
				"status": 1,
				"mileage": 0
			}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			vehicle := mock_service.NewMockVehicle(c)
			testCase.mockBehaviour(vehicle, testCase.inputVehicle)

			services := &service.Service{Vehicle: vehicle}
			handler := NewHandler(services)

			r := chi.NewRouter()
			r.Post("/api/v1/vehicles", handler.createVehicle)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/vehicles", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.JSONEq(t, testCase.expectedResponseBody, w.Body.String())
		})
	}

}
