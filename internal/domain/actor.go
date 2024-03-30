package domain

import "time"

type ActorToAdd struct {
	Name string
	Data string
}

type ActorData struct {
	Uuid     string     `json:"uuid"`
	Name     string     `json:"name"`
	Data     string     `json:"data"`
	Avatar   string     `json:"avatar"`
	Birthday time.Time  `json:"birthday"`
	Films    []FilmLink `json:"films"`
}

type ActorPreview struct {
	Uuid   string
	Name   string
	Avatar string
}

func NewMockActor() *ActorData {
	return &ActorData{
		Uuid: "1",
		Name: "Danya",
		Data: "Pizza",
	}
}
