package domain

import "time"

type ActorToAdd struct {
	Name string
	Data string
}

type ActorData struct {
	Uuid       string        `json:"uuid"`
	Name       string        `json:"name"`
	Avatar     string        `json:"avatar"`
	Birthday   time.Time     `json:"birthday"`
	Career     string        `json:"career"`
	Height     uint8         `json:"height"`
	BirthPlace string        `json:"birthPlace"`
	Genres     string        `json:"genres"`
	Spouse     string        `json:"spouse"`
	Films      []FilmPreview `json:"films"`
}

type ActorPreview struct {
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
