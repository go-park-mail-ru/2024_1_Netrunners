package domain

type User struct {
	Login    string `json:"login"`
	Name     string `json:"username"`
	Password string `json:"password"`
	Status   string
	Version  uint8
}
