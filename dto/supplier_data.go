package dto

type CreateSupplierRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UpdateSupplierRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type SupplierData struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
