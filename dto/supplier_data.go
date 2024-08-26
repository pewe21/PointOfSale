package dto

type CreateSupplierRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

type UpdateSupplierRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

type SupplierData struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
