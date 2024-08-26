package dto

type TypeData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateTypeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateTypeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
