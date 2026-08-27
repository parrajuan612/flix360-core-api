package domain

// PaginatedResponse es la estructura estándar para devolver listas
type PaginatedResponse struct {
	Total  int64       `json:"total"`
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
	Data   interface{} `json:"data"` // Aquí meteremos el arreglo de productos, o de activos, etc.
}
