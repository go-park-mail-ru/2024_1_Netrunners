package domain

import "time"

type FilmData struct {
	Uuid         string
	Preview      string
	Title        string
	Director     string
	AverageScore int
	ScoresCount  int
	Duration     int
	Date         time.Time
	Data         string
}

type FilmPreview struct {
	Uuid         string `json:"uuid"`
	Preview      string `json:"preview_data"`
	Title        string `json:"title"`
	Director     string `json:"author"`
	AverageScore int    `json:"average_score"`
	ScoresCount  int    `json:"scores_count"`
	Duration     int    `json:"duration"`
}

type FilmLink struct {
	Uuid  string
	Title string
}
