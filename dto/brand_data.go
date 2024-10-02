package dto

type BrandData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateBrandRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateBrandRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
