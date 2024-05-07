package domain

import "time"

type ActorToAdd struct {
	Name       string    `json:"name"`
	Avatar     string    `json:"avatar"`
	Birthday   time.Time `json:"birthday"`
	Career     string    `json:"career"`
	Height     uint32    `json:"height"`
	BirthPlace string    `json:"birthPlace"`
	Spouse     string    `json:"spouse"`
}

type SearchActors struct {
	Actors []ActorData `json:"actors"`
	Count  uint32      `json:"count"`
}

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

type ActorPreview struct {
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
