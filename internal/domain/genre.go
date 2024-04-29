package domain

type GenreFilms struct {
	Name  string        `json:"genre"`
	Uuid  string        `json:"genreUuid"`
	Films []FilmPreview `json:"films"`
}
