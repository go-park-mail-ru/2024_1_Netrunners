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

func NewMockActor() ActorData {
	return ActorData{
		Uuid:       "1",
		Name:       "Danya",
		Avatar:     "http://avatar",
		Birthday:   time.Now(),
		Career:     "career",
		Height:     192,
		BirthPlace: "Angarsk",
		Genres:     "Riddim",
		Spouse:     "Дабстеп",
		Films: []FilmLink{
			{"1", "Fast n Furious 1"},
			{"2", "Fast n Furious 2"},
			{"3", "Fast n Furious 3"},
		},
	}
}

func NewMockActorPreview() []ActorPreview {
	return []ActorPreview{
		{"1", "Fast n Furious 1", "avatar"},
		{"2", "Fast n Furious 2", "avatar"},
		{"3", "Fast n Furious 3", "avatar"},
	}
}
