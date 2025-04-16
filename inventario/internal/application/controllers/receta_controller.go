package controllers

import (
	domain "cmd/api/internal/domain/entities"
	"cmd/api/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecetaController interface {
	RouteGetReceta() string
	RouteAddReceta() string
	RouteRecetaIngrediente() string
	RouteRemoveIngrediente() string
	GetReceta(c *gin.Context)
	AddReceta(c *gin.Context)
	AddIngrediente(c *gin.Context)
	RemoveIngrediente(c *gin.Context)
}

type recetaController struct {
	service services.RecetaService
}

func NewrecetaController(service services.RecetaService) RecetaController {
	return &recetaController{
		service: service,
	}
}

func (h recetaController) RouteGetReceta() string {
	return "/receta/:id_receta"
}

func (h recetaController) RouteAddReceta() string {
	return "/receta"
}

func (h recetaController) RouteRecetaIngrediente() string {
	return "/receta/:id_receta/ingrediente"
}

func (h recetaController) RouteRemoveIngrediente() string {
	return "/receta/:id_receta/ingrediente/:id_ingrediente"
}

func (h recetaController) GetReceta(c *gin.Context) {
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

func (h recetaController) AddReceta(c *gin.Context) {
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

func (h recetaController) AddIngrediente(c *gin.Context) {
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

func (h recetaController) RemoveIngrediente(c *gin.Context) {
	id_receta := c.Param("id_receta")
	id_ingrediente := c.Param("id_ingrediente")

	err := h.service.RemoveIngrediente(c, id_receta, id_ingrediente)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error to remove ingrediente")
		return
	}

	c.JSON(http.StatusOK, "OK")
}
