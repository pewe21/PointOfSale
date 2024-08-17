package dto

type CustomerData struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type CreateCustomerRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UpdateCustomerRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
