package dtos

import "time"

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Model
}

type RegisterResponse struct {
	Username string `json:"username"`
	Success  bool   `json:"success"`
	AboutResponse
	Model
}

type Users struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedBy time.Time `json:"created_by"`
	Password  string    `json:"password"`
	Model
}

func (RegisterRequest) TableName() string {
	return "users"
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Model
}

func (LoginRequest) TableName() string {
	return "users"
}
