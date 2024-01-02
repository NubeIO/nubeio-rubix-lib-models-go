package dto

type User struct {
	Username string `json:"username" binding:"required" `
	Password string `json:"password" binding:"required"`
}
type ChangePassword struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
