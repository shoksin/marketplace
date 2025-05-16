package dto

type RegisterRequest struct {
	Username string
	Password string
	Email    string
}

type AdminRegisterRequest struct {
	Username string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}

type AdminLoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token string
}
