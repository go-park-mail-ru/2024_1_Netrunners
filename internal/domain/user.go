package domain

import "time"

//easyjson:json
type User struct {
	Uuid            string    `json:"uuid"`
	Email           string    `json:"login"`
	Name            string    `json:"username"`
	Password        string    `json:"password"`
	Version         uint32    `json:"version"`
	IsAdmin         bool      `json:"isAdmin"`
	Avatar          string    `json:"avatar"`
	RegisteredAt    time.Time `json:"registeredAt"`
	Birthday        time.Time `json:"birthday"`
	HasSubscription bool      `json:"hasSubscription"`
}

//easyjson:json
type UserSignUp struct {
	Email    string `json:"login"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

//easyjson:json
type UserPreview struct {
	Uuid   string
	Name   string
	Avatar string
}

//easyjson:json
type ProfileResponse struct {
	Status   int  `json:"status"`
	UserInfo User `json:"user"`
}

//easyjson:json
type ProfilePreviewResponse struct {
	Status      int         `json:"status"`
	UserPreview UserPreview `json:"user"`
}

type Subscription struct {
	Uuid        string  `json:"uuid"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
	Duration    uint32  `json:"duration"`
}
