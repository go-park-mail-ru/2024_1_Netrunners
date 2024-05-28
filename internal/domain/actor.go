package domain

import "time"

//easyjson:json
type ActorToAdd struct {
	Name       string    `json:"name"`
	Avatar     string    `json:"avatar"`
	Birthday   time.Time `json:"birthday"`
	Career     string    `json:"career"`
	Height     uint32    `json:"height"`
	BirthPlace string    `json:"birthPlace"`
	Spouse     string    `json:"spouse"`
}

//easyjson:json
type SearchActors struct {
	Actors []ActorData `json:"actors"`
	Count  uint32      `json:"count"`
}

//easyjson:json
type ActorData struct {
	Uuid       string        `json:"uuid"`
	Name       string        `json:"name"`
	Avatar     string        `json:"avatar"`
	Birthday   time.Time     `json:"birthday"`
	Career     string        `json:"career"`
	Height     uint32        `json:"height"`
	BirthPlace string        `json:"birthPlace"`
	Spouse     string        `json:"spouse"`
	Films      []FilmPreview `json:"films"`
}

//easyjson:json
type ActorPreview struct {
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
