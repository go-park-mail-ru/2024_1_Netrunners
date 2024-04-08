package domain

import "time"

type FilmData struct {
	Uuid         string
	Preview      string
	Title        string
	Director     string
	AverageScore float32
	ScoresCount  int
	Duration     int
	Date         time.Time
	Data         string
	AgeLimit     uint8
}

type FilmDataToAdd struct {
	Title       string
	Preview     string
	Director    string
	Data        string
	AgeLimit    uint8
	Duration    int
	PublishedAt time.Time
	Actors      []ActorData
}

type FilmPreview struct {
	Uuid         string  `json:"uuid"`
	Preview      string  `json:"preview_data"`
	Title        string  `json:"title"`
	Director     string  `json:"author"`
	AverageScore float32 `json:"average_score"`
	ScoresCount  int     `json:"scores_count"`
	Duration     int     `json:"duration"`
}

type FilmLink struct {
	Uuid  string `json:"uuid"`
	Title string `json:"title"`
}
