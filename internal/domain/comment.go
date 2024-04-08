package domain

import "time"

type Comment struct {
	Uuid     string
	FilmUuid string
	Author   string
	Text     string
	Score    int
	AddedAt  time.Time
}
