package controllers

import (
	"cmd/api/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventarioController struct {
	service services.InventarioService
}

func NewInventarioController(service services.InventarioService) InventarioController {
	return InventarioController{
		service: service,
	}
}

func (h *InventarioController) GetRouteInventario() string {
	return "/inventario/disponibilidad"
}

func (h *InventarioController) GetAvailable(c *gin.Context) {
	var request map[string]float64 // Id de ingrediente y cantidad requerida
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	disponible, faltantes, err := h.service.VerificarDisponibilidad(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"disponible": disponible, "faltantes": faltantes})
}

func (h *InventarioController) UpdateInventario(c *gin.Context) {
	var request map[string]float64 // Id de ingrediente y cantidad requerida
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	disponible, faltantes, err := h.service.VerificarDisponibilidad(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"disponible": disponible, "faltantes": faltantes})
}
