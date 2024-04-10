package domain

import "time"

type FilmData struct {
	Uuid         string    `json:"uuid"`
	Preview      string    `json:"preview"`
	Title        string    `json:"title"`
	Director     string    `json:"director"`
	AverageScore float32   `json:"averageScore"`
	ScoresCount  int       `json:"scoresCount"`
	Duration     int       `json:"duration"`
	Date         time.Time `json:"date"`
	Data         string    `json:"data"`
	AgeLimit     uint8     `json:"ageLimit"`
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
