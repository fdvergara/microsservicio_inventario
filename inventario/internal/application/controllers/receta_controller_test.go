package controllers

import (
	"bytes"
	domain "cmd/api/internal/domain/entities"
	"cmd/api/internal/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetReceta(t *testing.T) {
	receta := domain.Receta{
		Id:     "12345",
		Nombre: "test",
		Ingredientes: []domain.IngredienteCantidad{
			{
				IngredienteId: "12345",
				Cantidad:      20.0,
			},
		},
	}
	service := mocks.RecetaServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Get", "12345").Return(receta, nil)

	router := gin.Default()
	controller := NewrecetaController(&service)
	router.GET(controller.RouteGetReceta(), controller.GetReceta)

	req, _ := http.NewRequest(http.MethodGet, "/receta/12345", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetRecetaNotFound(t *testing.T) {
	service := mocks.RecetaServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Get", "12345").Return(nil, errors.New("not_found"))

	router := gin.Default()
	controller := NewrecetaController(&service)
	router.GET(controller.RouteGetReceta(), controller.GetReceta)

	req, _ := http.NewRequest(http.MethodGet, "/receta/12345", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestGetRecetaError(t *testing.T) {
	service := mocks.RecetaServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Get", "12345").Return(nil, errors.New("error"))

	router := gin.Default()
	controller := NewrecetaController(&service)
	router.GET(controller.RouteGetReceta(), controller.GetReceta)

	req, _ := http.NewRequest(http.MethodGet, "/receta/12345", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func TestAddReceta(t *testing.T) {
	recetaSaved := `{"nombre": "12345","ingredientes":[{"ingrediente_id":"12345","cantidad":0}]}`
	service := mocks.RecetaServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Create", mock.Anything).Return("12345", nil)

	router := gin.Default()
	controller := NewrecetaController(&service)
	router.POST(controller.RouteAddReceta(), controller.AddReceta)

	req, _ := http.NewRequest(http.MethodPost, "/receta", bytes.NewBufferString(recetaSaved))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestAddRecetaBadRequest(t *testing.T) {
	recetaSaved := ``
	service := mocks.RecetaServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Create", mock.Anything).Return("12345", nil)

	router := gin.Default()
	controller := NewrecetaController(&service)
	router.POST(controller.RouteAddReceta(), controller.AddReceta)

	req, _ := http.NewRequest(http.MethodPost, "/receta", bytes.NewBufferString(recetaSaved))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestAddRecetaError(t *testing.T) {
	recetaSaved := `{"nombre": "12345","ingredientes":[{"ingrediente_id":"12345","cantidad":0}]}`
	service := mocks.RecetaServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Create", mock.Anything).Return("12345", errors.New("error"))

	router := gin.Default()
	controller := NewrecetaController(&service)
	router.POST(controller.RouteAddReceta(), controller.AddReceta)

	req, _ := http.NewRequest(http.MethodPost, "/receta", bytes.NewBufferString(recetaSaved))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func TestAddIngredienteReceta(t *testing.T) {
	ingrediente := `{"ingrediente_id":"12345","cantidad":40}`
	service := mocks.RecetaServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Get", "12345").Return(mock.Anything, nil)
	service.On("AddIngrediente", mock.Anything, mock.Anything).Return(nil)

	router := gin.Default()
	controller := NewrecetaController(&service)
	router.PUT(controller.RouteRecetaIngrediente(), controller.AddIngrediente)

	req, _ := http.NewRequest(http.MethodPut, "/receta/12345/ingrediente", bytes.NewBufferString(ingrediente))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestAddIngredienteRecetaBadRequest(t *testing.T) {
	ingrediente := ``
	service := mocks.RecetaServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Get", "12345").Return(mock.Anything, nil)
	service.On("AddIngrediente", mock.Anything, mock.Anything).Return(nil)

	router := gin.Default()
	controller := NewrecetaController(&service)
	router.PUT(controller.RouteRecetaIngrediente(), controller.AddIngrediente)

	req, _ := http.NewRequest(http.MethodPut, "/receta/12345/ingrediente", bytes.NewBufferString(ingrediente))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestAddIngredienteRecetaError(t *testing.T) {
	ingrediente := `{"ingrediente_id":"12345","cantidad":40}`
	service := mocks.RecetaServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Get", "12345").Return(mock.Anything, nil)
	service.On("AddIngrediente", mock.Anything, mock.Anything).Return(errors.New("error"))

	router := gin.Default()
	controller := NewrecetaController(&service)
	router.PUT(controller.RouteRecetaIngrediente(), controller.AddIngrediente)

	req, _ := http.NewRequest(http.MethodPut, "/receta/12345/ingrediente", bytes.NewBufferString(ingrediente))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}
