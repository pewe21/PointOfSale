package dto

type RoleData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type CreateRoleRequest struct {
	Name        string `json:"name" validate:"required"`
	DisplayName string `json:"display_name"`
}

type UpdateRoleRequest struct {
	DisplayName string `json:"display_name" validate:"required"`
}

type AddCustomerRoleRequest struct {
	RoleId     string `json:"role_id" validate:"required,uuid"`
	CustomerId string `json:"customer_id" validate:"required,uuid"`
}

type UpdateCustomerRoleRequest struct {
	RoleId string `json:"role_id" validate:"required,uuid"`
}
