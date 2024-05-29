package domain

//easyjson:json
type GenreFilms struct {
	Name  string        `json:"genre"`
	Uuid  string        `json:"genreUuid"`
	Films []FilmPreview `json:"films"`
}

//easyjson:json
type Genre struct {
	Name string `json:"genreName"`
	Uuid string `json:"genreUuid"`
}
