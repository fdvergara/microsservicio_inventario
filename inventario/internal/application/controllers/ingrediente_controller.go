package controllers

import (
	domain "cmd/api/internal/domain/entities"
	"cmd/api/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IngredienteController struct {
	service services.IngredienteService
}

type cantidadRequest struct {
	Cantidad int `json:"cantidad"`
}

func NewIngredienteController(service services.IngredienteService) IngredienteController {
	return IngredienteController{
		service: service,
	}
}

func (c *IngredienteController) GetRouteIngrediente() string {
	return "/ingrediente/:id_ingrediente"
}

func (c *IngredienteController) GetRouteCantidadIngrediente() string {
	return "/ingrediente/:id_ingrediente/cantidad"
}

func (h *IngredienteController) GetRouteAddIngrediente() string {
	return "/ingrediente"
}

func (h *IngredienteController) Get(c *gin.Context) {
	id_ingrediente := c.Param("id_ingrediente")
	ingrediente, err := h.service.Get(c, id_ingrediente)
	if err != nil {
		if err.Error() == "not_found" {
			c.String(http.StatusNotFound, "")
			return
		}
		c.String(http.StatusInternalServerError, "Error to get ingrediente")
		return
	}
	c.JSON(http.StatusOK, ingrediente)
}

func (h *IngredienteController) AddIngrediente(c *gin.Context) {
	var ingrediente domain.Ingrediente
	if err := c.BindJSON(&ingrediente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.Create(c, ingrediente)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error to create ingrediente")
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (h *IngredienteController) UpdateCantidadIngrediente(c *gin.Context) {
	var cantidad cantidadRequest
	id_ingrediente := c.Param("id_ingrediente")
	if err := c.BindJSON(&cantidad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Update(c, id_ingrediente, float64(cantidad.Cantidad))
	if err != nil {
		if err.Error() == "not_found" {
			c.String(http.StatusNotFound, "")
			return
		}
		c.String(http.StatusInternalServerError, "Error to update cantidad")
		return
	}

	c.String(http.StatusOK, "OK")
}
