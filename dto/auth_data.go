package dto

type SignInRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

type SignData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
