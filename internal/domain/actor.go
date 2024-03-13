package domain

import "time"

type ActorToAdd struct {
	Name string
	Data string
}

type ActorData struct {
	Uuid     string
	Name     string
	Data     string
	Avatar   string
	Birthday time.Time
	Films    []FilmLink
}

type ActorPreview struct {
	Uuid   string
	Name   string
	Avatar string
}
