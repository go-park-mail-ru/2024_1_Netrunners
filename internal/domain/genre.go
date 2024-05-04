package domain

type GenreFilms struct {
	Name  string        `json:"genre"`
	Uuid  string        `json:"genreUuid"`
	Films []FilmPreview `json:"films"`
}

type Genre struct {
	Name string `json:"genreName"`
	Uuid string `json:"genreUuid"`
}
