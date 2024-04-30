package domain

import "time"

type FilmData struct {
	Uuid         string    `json:"uuid"`
	IsSerial     bool      `json:"isSerial"`
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
	Genres       []string  `json:"genres"`
}

type CommonFilmData struct {
	Uuid         string    `json:"uuid"`
	IsSerial     bool      `json:"isSerial"`
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
	Seasons      []Season  `json:"seasons"`
}

type Episode struct {
	Link string `json:"link"`
}

type Season struct {
	Series []Episode `json:"series"`
}

type SerialData struct {
	Uuid         string    `json:"uuid"`
	IsSerial     bool      `json:"isSerial"`
	Preview      string    `json:"preview"`
	Title        string    `json:"title"`
	Seasons      []Season  `json:"seasons"`
	Director     string    `json:"director"`
	AverageScore float32   `json:"averageScore"`
	ScoresCount  uint64    `json:"scoresCount"`
	Duration     uint32    `json:"duration"`
	Date         time.Time `json:"date"`
	Data         string    `json:"data"`
	AgeLimit     uint32    `json:"ageLimit"`
}

type FilmDataToAdd struct {
	Title       string
	Preview     string
	Director    string
	Data        string
	AgeLimit    uint32
	Duration    uint32
	PublishedAt time.Time
	Actors      []ActorData
	Genres      []string
}

type FilmPreview struct {
	Uuid         string  `json:"uuid"`
	IsSerial     bool    `json:"isSerial"`
	Preview      string  `json:"preview_data"`
	Title        string  `json:"title"`
	Director     string  `json:"author"`
	AverageScore float32 `json:"average_score"`
	ScoresCount  uint64  `json:"scores_count"`
	Duration     uint32  `json:"duration"`
	AgeLimit     uint32  `json:"ageLimit"`
}
