package model

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
