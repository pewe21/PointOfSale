package dto

type RoleData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type CreateRoleRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type UpdateRoleRequest struct {
	DisplayName string `json:"display_name"`
}

type AddCustomerRoleRequest struct {
	RoleId     string `json:"role_id"`
	CustomerId string `json:"customer_id"`
}

type UpdateCustomerRoleRequest struct {
	RoleId string `json:"role_id"`
}
