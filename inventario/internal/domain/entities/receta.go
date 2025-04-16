package domain

type Receta struct {
	Id           string                `json:"id"`
	Nombre       string                `json:"nombre"`
	Ingredientes []IngredienteCantidad `json:"ingredientes"`
}

type IngredienteCantidad struct {
	IngredienteId string  `json:"ingrediente_id"`
	Cantidad      float64 `json:"cantidad"`
}
