package domain

import "time"

type User struct {
	Uuid         string `json:"uuid"`
	Email        string `json:"login"`
	Name         string `json:"username"`
	Password     string `json:"password"`
	Version      uint8
	IsAdmin      bool      `json:"isAdmin"`
	Avatar       string    `json:"avatar"`
	RegisteredAt time.Time `json:"registeredAt"`
	Birthday     time.Time `json:"birthday"`
}

type UserSignUp struct {
	Email    string `json:"login"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

type UserPreview struct {
	Uuid   string
	Name   string
	Avatar string
}
