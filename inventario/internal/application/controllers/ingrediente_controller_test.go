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
)

func TestGetIngrediente(t *testing.T) {
	ingrediente := domain.Ingrediente{
		Id:           "12345",
		Nombre:       "test",
		Cantidad:     20,
		UnidadMedida: "unidad",
	}
	service := mocks.IngredienteServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Get", "12345").Return(ingrediente, nil)

	router := gin.Default()
	controller := NewIngredienteController(&service)
	router.GET(controller.GetRouteIngrediente(), controller.Get)

	req, _ := http.NewRequest(http.MethodGet, "/ingrediente/12345", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetIngredienteError(t *testing.T) {
	service := mocks.IngredienteServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Get", "12345").Return(nil, errors.New("error"))

	router := gin.Default()
	controller := NewIngredienteController(&service)
	router.GET(controller.GetRouteIngrediente(), controller.Get)

	req, _ := http.NewRequest(http.MethodGet, "/ingrediente/12345", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)

}

func TestGetIngredienteNotFound(t *testing.T) {
	service := mocks.IngredienteServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Get", "12345").Return(nil, errors.New("not_found"))

	router := gin.Default()
	controller := NewIngredienteController(&service)
	router.GET(controller.GetRouteIngrediente(), controller.Get)

	req, _ := http.NewRequest(http.MethodGet, "/ingrediente/12345", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

}

func TestAddIngrediente(t *testing.T) {
	ingrediente := `{"id":"12345","nombre":"test","cantidad":20,"unidad_medida": "unidad"}`
	ingredientes := domain.Ingrediente{
		Id:           "12345",
		Nombre:       "test",
		Cantidad:     20,
		UnidadMedida: "unidad",
	}
	service := mocks.IngredienteServiceMock{}
	gin.SetMode(gin.TestMode)

	id := "12345"
	service.On("Create", ingredientes).Return(&id, nil)

	router := gin.Default()
	controller := NewIngredienteController(&service)
	router.POST(controller.GetRouteAddIngrediente(), controller.AddIngrediente)

	req, _ := http.NewRequest(http.MethodPost, "/ingrediente", bytes.NewBufferString(ingrediente))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestAddIngredienteError(t *testing.T) {
	ingrediente := `{"id":"12345","nombre":"test","cantidad":20,"unidad_medida": "unidad"}`
	ingredientes := domain.Ingrediente{
		Id:           "12345",
		Nombre:       "test",
		Cantidad:     20,
		UnidadMedida: "unidad",
	}
	service := mocks.IngredienteServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Create", ingredientes).Return(nil, errors.New("error"))

	router := gin.Default()
	controller := NewIngredienteController(&service)
	router.POST(controller.GetRouteAddIngrediente(), controller.AddIngrediente)

	req, _ := http.NewRequest(http.MethodPost, "/ingrediente", bytes.NewBufferString(ingrediente))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func TestAddIngredienteBadRequest(t *testing.T) {
	ingrediente := ``
	ingredientes := domain.Receta{}
	service := mocks.IngredienteServiceMock{}
	gin.SetMode(gin.TestMode)

	service.On("Create", ingredientes).Return(nil, errors.New("error"))

	router := gin.Default()
	controller := NewIngredienteController(&service)
	router.POST(controller.GetRouteAddIngrediente(), controller.AddIngrediente)

	req, _ := http.NewRequest(http.MethodPost, "/ingrediente", bytes.NewBufferString(ingrediente))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestUpdateCantidadIngrediente(t *testing.T) {
	cantidad := `{"cantidad":30}`
	service := mocks.IngredienteServiceMock{}
	gin.SetMode(gin.TestMode)

	id := "12345"
	service.On("Update", id, float64(30)).Return(nil)

	router := gin.Default()
	controller := NewIngredienteController(&service)
	router.PUT(controller.GetRouteCantidadIngrediente(), controller.UpdateCantidadIngrediente)

	req, _ := http.NewRequest(http.MethodPut, "/ingrediente/12345/cantidad", bytes.NewBufferString(cantidad))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateCantidadIngredienteBadRequest(t *testing.T) {
	cantidad := ``
	service := mocks.IngredienteServiceMock{}
	gin.SetMode(gin.TestMode)

	id := "12345"
	service.On("Update", id, float64(30)).Return(nil)

	router := gin.Default()
	controller := NewIngredienteController(&service)
	router.PUT(controller.GetRouteCantidadIngrediente(), controller.UpdateCantidadIngrediente)

	req, _ := http.NewRequest(http.MethodPut, "/ingrediente/12345/cantidad", bytes.NewBufferString(cantidad))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestUpdateIngredienteError(t *testing.T) {
	cantidad := `{"cantidad":30}`
	service := mocks.IngredienteServiceMock{}
	gin.SetMode(gin.TestMode)

	id := "12345"
	service.On("Update", id, float64(30)).Return(errors.New("error"))

	router := gin.Default()
	controller := NewIngredienteController(&service)
	router.PUT(controller.GetRouteCantidadIngrediente(), controller.UpdateCantidadIngrediente)

	req, _ := http.NewRequest(http.MethodPut, "/ingrediente/12345/cantidad", bytes.NewBufferString(cantidad))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func TestUpdateCantidadIngredienteNotFound(t *testing.T) {
	cantidad := `{"cantidad":30}`
	service := mocks.IngredienteServiceMock{}
	gin.SetMode(gin.TestMode)

	id := "12345"
	service.On("Update", id, float64(30)).Return(errors.New("not_found"))

	router := gin.Default()
	controller := NewIngredienteController(&service)
	router.PUT(controller.GetRouteCantidadIngrediente(), controller.UpdateCantidadIngrediente)

	req, _ := http.NewRequest(http.MethodPut, "/ingrediente/12345/cantidad", bytes.NewBufferString(cantidad))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

}
