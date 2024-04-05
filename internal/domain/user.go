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
