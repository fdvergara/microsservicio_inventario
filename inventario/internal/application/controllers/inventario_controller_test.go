package controllers

import (
	"bytes"
	"cmd/api/internal/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAvailable(t *testing.T) {
	ingrediente := `{"1607a668-2213-4bdb-b605-4c496d155dac": 14}`
	returnIngrediente := map[string]float64{
		"faltantes": 0,
	}

	service := mocks.InventarioServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("VerificarDisponibilidad", mock.Anything).Return(true, returnIngrediente, nil)

	router := gin.Default()
	controller := NewInventarioController(&service)
	router.POST(controller.GetRouteInventario(), controller.GetAvailable)

	req, _ := http.NewRequest(http.MethodPost, "/inventario/disponibilidad", bytes.NewBufferString(ingrediente))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetAvailableError(t *testing.T) {
	ingrediente := `{"1607a668-2213-4bdb-b605-4c496d155dac": 14}`

	service := mocks.InventarioServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("VerificarDisponibilidad", mock.Anything).Return(nil, nil, errors.New("error"))

	router := gin.Default()
	controller := NewInventarioController(&service)
	router.POST(controller.GetRouteInventario(), controller.GetAvailable)

	req, _ := http.NewRequest(http.MethodPost, "/inventario/disponibilidad", bytes.NewBufferString(ingrediente))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)

}

func TestGetAvailableBadRequest(t *testing.T) {
	ingrediente := ``

	service := mocks.InventarioServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("VerificarDisponibilidad", mock.Anything).Return(true, nil, nil)

	router := gin.Default()
	controller := NewInventarioController(&service)
	router.POST(controller.GetRouteInventario(), controller.GetAvailable)

	req, _ := http.NewRequest(http.MethodPost, "/inventario/disponibilidad", bytes.NewBufferString(ingrediente))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

}
