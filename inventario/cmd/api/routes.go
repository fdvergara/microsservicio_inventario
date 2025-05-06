package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func MapRoutes(app *gin.Engine) {
	// Controller de ingrediente
	app.GET(appControllers.ingredienteController.GetRouteIngrediente(), appControllers.ingredienteController.Get)
	app.POST(appControllers.ingredienteController.GetRouteAddIngrediente(), appControllers.ingredienteController.AddIngrediente)
	app.PUT(appControllers.ingredienteController.GetRouteCantidadIngrediente(), appControllers.ingredienteController.UpdateCantidadIngrediente)

	//Controller de recetas
	app.GET(appControllers.recetaController.RouteGetReceta(), appControllers.recetaController.GetReceta)
	app.POST(appControllers.recetaController.RouteAddReceta(), appControllers.recetaController.AddReceta)
	app.PUT(appControllers.recetaController.RouteRecetaIngrediente(), appControllers.recetaController.AddIngrediente)

	//Controller de inventario
	app.POST(appControllers.inventarioController.GetRouteInventario(), appControllers.inventarioController.GetAvailable)
}

func InitConsumer() {
	err := appControllers.consumer.Start()
	if err != nil {
		log.Fatalf("Error iniciando el consumidor: %v", err)
	}
}
