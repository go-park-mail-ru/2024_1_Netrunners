package domain

import "time"

type FilmData struct {
	Uuid         string    `json:"uuid"`
	Preview      string    `json:"preview"`
	Title        string    `json:"title"`
	Link         string    `json:"link"`
	Director     string    `json:"director"`
	AverageScore float32   `json:"averageScore"`
	ScoresCount  uint64    `json:"scoresCount"`
	Duration     uint32    `json:"duration"`
	Date         time.Time `json:"date"`
	Data         string    `json:"data"`
	AgeLimit     uint32    `json:"ageLimit"`
	Genres       []Genre   `json:"genres"`
}

type FilmDataToAdd struct {
	Title       string    `json:"title"`
	Preview     string    `json:"preview"`
	Director    string    `json:"director"`
	Data        string    `json:"data"`
	AgeLimit    uint32    `json:"ageLimit"`
	Duration    uint32    `json:"duration"`
	PublishedAt time.Time `json:"publishedAt"`
	Genres      []string  `json:"genres"`
	Link        string    `json:"link"`
}

type FilmPreview struct {
	Uuid         string  `json:"uuid"`
	Preview      string  `json:"preview_data"`
	Title        string  `json:"title"`
	Director     string  `json:"author"`
	AverageScore float32 `json:"average_score"`
	ScoresCount  uint64  `json:"scores_count"`
	Duration     uint32  `json:"duration"`
	AgeLimit     uint32  `json:"ageLimit"`
}

type FilmToAdd struct {
	FilmData      FilmDataToAdd `json:"filmData"`
	Actors        []ActorToAdd  `json:"actors"`
	DirectorToAdd DirectorToAdd `json:"directorToAdd"`
}
