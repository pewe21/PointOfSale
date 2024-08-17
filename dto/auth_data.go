package dto

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

type SignData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
