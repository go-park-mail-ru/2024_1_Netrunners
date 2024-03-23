package domain

type User struct {
	Email    string `json:"login"`
	Name     string `json:"username"`
	Password string `json:"password"`
	Status   string
	Version  uint8
	IsAdmin  bool
	Avatar   string
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
