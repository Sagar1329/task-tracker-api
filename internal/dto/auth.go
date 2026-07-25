package dto

type SignupRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"`

	Email string `json:"email" binding:"required,email,max=255"`

	Password string `json:"password" binding:"required,min=8,max=100"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required,email,max=255"`

	Password string `json:"password" binding:"required,min=8,max=100"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}