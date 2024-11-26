package dto

type CustomerData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	//Username string `json:"username"`
	//Email    string `json:"email"`
	Phone string `json:"phone"`
}

type CreateCustomerRequest struct {
	Name string `json:"name" validate:"required,min=3,max=50"`
	//Username string `json:"username"`
	//Password string `json:"password"`
	//Email    string `json:"email"`
	Phone string `json:"phone" validate:"required,min=10,max=15"`
}

type UpdateCustomerRequest struct {
	Name string `json:"name" validate:"required,min=5,max=50"`
	//Username string `json:"username"`
	//Phone string `json:"phone"`
	Phone   string `json:"phone" validate:"required,min=10,max=15"`
	Address string `json:"address"`
}
