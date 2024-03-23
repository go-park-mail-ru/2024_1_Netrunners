package mockdb

import (
	"fmt"
	"sync"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type FilmsMockDB struct {
	storage map[string]domain.FilmPreview
	mutex   *sync.RWMutex
}

func InitFilmsMockDB() *FilmsMockDB {
	return &FilmsMockDB{
		storage: make(map[string]domain.FilmPreview),
		mutex:   &sync.RWMutex{},
	}
}

func (db *FilmsMockDB) AddFilm(film domain.FilmPreview) error {
	db.mutex.RLock()
	_, ok := db.storage[film.Uuid]
	db.mutex.RUnlock()

	if ok {
		return fmt.Errorf("film with given Id already exists: %w", myerrors.ErrInternalServerError)
	}

	db.mutex.Lock()
	db.storage[film.Uuid] = film
	db.mutex.Unlock()

	return nil
}

func (db *FilmsMockDB) RemoveFilm(id string) error {
	db.mutex.RLock()
	_, ok := db.storage[id]
	db.mutex.RUnlock()

	if !ok {
		return fmt.Errorf("film with given Id doesn't exists: %w", myerrors.ErrInternalServerError)
	}

	db.mutex.Lock()
	delete(db.storage, id)
	db.mutex.Unlock()

	return nil
}

func (db *FilmsMockDB) GetFilmPreview(id string) (domain.FilmPreview, error) {
	db.mutex.RLock()
	film, ok := db.storage[id]
	db.mutex.RUnlock()

	if !ok {
		return domain.FilmPreview{}, fmt.Errorf("film with given Id doesn't exists: %w", myerrors.ErrNoSuchFilm)
	}

	return film, nil
}

func (db *FilmsMockDB) GetAllFilmsPreviews() []domain.FilmPreview {
	films := make([]domain.FilmPreview, 0, len(db.storage))
	db.mutex.RLock()
	for _, filmPreview := range db.storage {
		films = append(films, filmPreview)
	}
	db.mutex.RUnlock()

	return films
}
