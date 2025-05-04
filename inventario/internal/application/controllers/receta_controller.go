package controllers

import (
	domain "cmd/api/internal/domain/entities"
	"cmd/api/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecetaController struct {
	service services.RecetaService
}

func NewrecetaController(service services.RecetaService) RecetaController {
	return RecetaController{
		service: service,
	}
}

func (h RecetaController) RouteGetReceta() string {
	return "/receta/:id_receta"
}

func (h RecetaController) RouteAddReceta() string {
	return "/receta"
}

func (h RecetaController) RouteRecetaIngrediente() string {
	return "/receta/:id_receta/ingrediente"
}

func (h RecetaController) RouteRemoveIngrediente() string {
	return "/receta/:id_receta/ingrediente/:id_ingrediente"
}

func (h RecetaController) GetReceta(c *gin.Context) {
	id_receta := c.Param("id_receta")
	receta, err := h.service.Get(c, id_receta)
	if err != nil {
		if err.Error() == "not_found" {
			c.String(http.StatusNotFound, "")
			return
		}
		c.String(http.StatusInternalServerError, "Error to get receta")
		return
	}
	c.JSON(http.StatusOK, receta)
}

func (h RecetaController) AddReceta(c *gin.Context) {
	var receta domain.Receta
	if err := c.BindJSON(&receta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.Create(c, receta)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error to create receta")
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (h RecetaController) AddIngrediente(c *gin.Context) {
	var ingredienteCantidad domain.IngredienteCantidad
	id_receta := c.Param("id_receta")
	if err := c.BindJSON(&ingredienteCantidad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.AddIngrediente(c, id_receta, ingredienteCantidad)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error to add ingrediente")
		return
	}

	c.JSON(http.StatusOK, "OK")
}
