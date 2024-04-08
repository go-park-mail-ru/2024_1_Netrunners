package domain

import "time"

type ActorToAdd struct {
	Name string
	Data string
}

type ActorData struct {
	Uuid       string    `json:"uuid"`
	Name       string    `json:"name"`
	Avatar     string    `json:"avatar"`
	Birthday   time.Time `json:"birthday"`
	Career     string
	Height     uint8
	BirthPlace string
	Genres     string
	Spouse     string
	Films      []FilmLink `json:"films"`
}

type ActorPreview struct {
	Uuid   string
	Name   string
	Avatar string
}
