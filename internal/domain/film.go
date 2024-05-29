package domain

import "time"

//easyjson:json
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
	Genres       []Genre   `json:"genres"`
	WithSub      bool      `json:"withSubscription"`
}

//easyjson:json
type SearchFilms struct {
	Films []FilmData `json:"films"`
	Count uint32     `json:"count"`
}

//easyjson:json
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
	Genres       []Genre   `json:"genres"`
	WithSub      bool      `json:"withSubscription"`
}

//easyjson:json
type Episode struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

//easyjson:json
type Season struct {
	Series []Episode `json:"series"`
}

//easyjson:json
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
	Genres       []Genre   `json:"genres"`
	WithSub      bool      `json:"withSubscription"`
}

//easyjson:json
type FilmDataToAdd struct {
	Title       string    `json:"title"`
	IsSerial    bool      `json:"isSerial"`
	Preview     string    `json:"preview"`
	Director    string    `json:"director"`
	Data        string    `json:"data"`
	AgeLimit    uint32    `json:"ageLimit"`
	Duration    uint32    `json:"duration"`
	PublishedAt time.Time `json:"publishedAt"`
	Genres      []string  `json:"genres"`
	Link        string    `json:"link"`
	Seasons     []Season  `json:"seasons,omitempty"`
}

//easyjson:json
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

//easyjson:json
type FilmToAdd struct {
	FilmData      FilmDataToAdd `json:"filmData"`
	Actors        []ActorToAdd  `json:"actors"`
	DirectorToAdd DirectorToAdd `json:"directorToAdd"`
}

//easyjson:json
type TopFilm struct {
	Uuid     string `json:"uuid"`
	IsSerial bool   `json:"isSerial"`
	Title    string `json:"title"`
	Preview  string `json:"preview_data"`
	Data     string `json:"data"`
}

//easyjson:json
type ShortSearchResponse struct {
	Status int            `json:"status"`
	Films  []FilmPreview  `json:"films"`
	Actors []ActorPreview `json:"actors"`
}

//easyjson:json
type LongSearchResponse struct {
	Status int         `json:"status"`
	Films  []FilmData  `json:"films"`
	Actors []ActorData `json:"actors"`
	Count  int         `json:"searchResCount"`
}

//easyjson:json
type FilmsPreviewsResponse struct {
	Status int           `json:"status"`
	Films  []FilmPreview `json:"films"`
}

//easyjson:json
type FilmDataResponse struct {
	Status   int         `json:"status"`
	FilmData interface{} `json:"film"`
}

//easyjson:json
type FilmCommentsResponse struct {
	Status   int       `json:"status"`
	Comments []Comment `json:"comments"`
}

//easyjson:json
type FilmActorsResponse struct {
	Status int            `json:"status"`
	Actors []ActorPreview `json:"actors"`
}

//easyjson:json
type ActorResponse struct {
	Status int       `json:"status"`
	Actor  ActorData `json:"actor"`
}

//easyjson:json
type DataToFavorite struct {
	FilmUuid string `json:"filmUuid"`
	UserUuid string `json:"userUuid"`
}

//easyjson:json
type GenresResponse struct {
	Status      int          `json:"status"`
	GenresFilms []GenreFilms `json:"genres"`
}

//easyjson:json
type TopFilmsResponse struct {
	Status int       `json:"status"`
	Films  []TopFilm `json:"films"`
}
