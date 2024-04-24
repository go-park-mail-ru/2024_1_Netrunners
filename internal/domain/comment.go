package domain

import "time"

type Comment struct {
	Uuid     string
	FilmUuid string
	Author   string
	Text     string
	Score    uint32
	AddedAt  time.Time
}
