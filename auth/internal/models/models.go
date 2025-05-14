package models

type LoginResponse struct {
	Token  string `json:"token"`
	Status string `json:"status"`
	Error  string `json:"error"`
}
