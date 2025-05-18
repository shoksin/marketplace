package dto

type RegisterRequest struct {
	Username string
	Password string
	Email    string
	Birthday string
}

type AdminRegisterRequest struct {
	Username string
	Password string
}

type RegisterResponse struct {
	ID int64
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
