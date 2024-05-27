package domain

import "time"

type Comment struct {
	Uuid       string    `json:"uuid"`
	FilmUuid   string    `json:"filmUuid"`
	AuthorUuid string    `json:"authorUuid"`
	Author     string    `json:"author"`
	Text       string    `json:"text"`
	Score      uint32    `json:"score"`
	AddedAt    time.Time `json:"added_at"`
}

type CommentToAdd struct {
	FilmUuid   string `json:"filmUuid"`
	AuthorUuid string `json:"authorUuid"`
	Text       string `json:"text"`
	Score      uint32 `json:"score"`
}

type CommentToRemove struct {
	FilmUuid   string `json:"filmUuid"`
	AuthorUuid string `json:"authorUuid"`
}
