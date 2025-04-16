package domain

type Ingrediente struct {
	Id           string  `json:"id"`
	Nombre       string  `json:"nombre"`
	Cantidad     float64 `json:"cantidad"`
	UnidadMedida string  `json:"unidad_medida"`
}
