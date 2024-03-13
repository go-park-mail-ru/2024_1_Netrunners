package domain

import "time"

type User struct {
	Uuid         string
	Email        string `json:"login"`
	Name         string `json:"username"`
	Password     string `json:"password"`
	RegisteredAt time.Time
	Birthday     time.Time
	Version      uint8
	IsAdmin      bool
	Avatar       string
}

type UserSignUp struct {
	Email    string
	Name     string
	Password string
}

type UserPreview struct {
	Name   string
	Avatar string
}
