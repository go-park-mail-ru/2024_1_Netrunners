package domain

import "time"

type User struct {
	Uuid         string
	Email        string `json:"login"`
	Name         string `json:"username"`
	Password     string `json:"password"`
	Version      uint8
	IsAdmin      bool
	Avatar       string
	RegisteredAt time.Time
	Birthday     time.Time
}

type UserSignUp struct {
	Email    string `json:"login"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

type UserPreview struct {
	Name   string
	Avatar string
}

func NewMockUser() User {
	return User{
		Uuid:         "1",
		Email:        "cakethefake@gmail.com",
		Name:         "Danya",
		Password:     "123456789",
		IsAdmin:      true,
		RegisteredAt: time.Now(),
		Birthday:     time.Now(),
	}
}

func NewMockUserSignUp() UserSignUp {
	return UserSignUp{
		Email:    "cakethefake@gmail.com",
		Name:     "Danya",
		Password: "123456789",
	}
}
