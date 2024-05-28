package domain

import "time"

//easyjson:json
type Comment struct {
	Uuid       string    `json:"uuid"`
	FilmUuid   string    `json:"filmUuid"`
	AuthorUuid string    `json:"authorUuid"`
	Author     string    `json:"author"`
	Text       string    `json:"text"`
	Score      uint32    `json:"score"`
	AddedAt    time.Time `json:"added_at"`
}

//easyjson:json
type CommentToAdd struct {
	FilmUuid   string `json:"filmUuid"`
	AuthorUuid string `json:"authorUuid"`
	Text       string `json:"text"`
	Score      uint32 `json:"score"`
}

//easyjson:json
type CommentToRemove struct {
	FilmUuid   string `json:"filmUuid"`
	AuthorUuid string `json:"authorUuid"`
}
