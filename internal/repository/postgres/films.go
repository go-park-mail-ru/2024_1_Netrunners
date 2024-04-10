package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type FilmsStorage struct {
	pool PgxIface
}

func NewFilmsStorage(pool PgxIface) (*FilmsStorage, error) {
	return &FilmsStorage{
		pool: pool,
	}, nil
}

const getFilmDataByUuid = `
		SELECT f.uuid, f.title, f.banner, d.name, f.data, f.duration, f.published_at, AVG(c.score), COUNT(c.id)
		FROM film f
		LEFT JOIN comment c ON f.id = c.film
		JOIN director d ON f.director = d.id
		WHERE f.uuid = $1
		GROUP BY f.uuid, f.title,  f.banner, d.name, f.published_at, f.duration, f.data;`

const getAmountOfDirectorsByName = `
		SELECT COUNT(*)
		FROM director
		WHERE name = $1;`

const insertDirector = `
		INSERT INTO director (name) VALUES ($1);`

const getDirectorsIdByName = `
		SELECT id
		FROM director
		WHERE name = $1;`

const insertFilm = `
		INSERT INTO film (title, banner, director, data, age_limit, duration, published_at) 
    	VALUES ($1, $2, $3, $4, $5, $6, $7);`

const getFilmIdByTitle = `
		SELECT id
		FROM film
		WHERE title = $1;`

const getAmountOfActorsByName = `
		SELECT COUNT(*)
		FROM actor
		WHERE actor.name = $1;`

const insertActor = `
		INSERT INTO actor (name) VALUES ($1);`

const getActorId = `
		SELECT id
		FROM actor
		WHERE actor.name = $1;`

const insertIntoFilmActors = `
		INSERT INTO film_actor (film, actor) VALUES ($1, $2);`

const deleteFilm = `
		DELETE FROM film
		WHERE uuid = $1;`

const getFilmPreview = `
		SELECT f.uuid, f.title, f.banner, d.name, f.duration, AVG(c.score), COUNT(c.id)
		FROM film f
		LEFT JOIN comment c ON f.id = c.film
		JOIN director d ON f.director = d.id
		WHERE f.uuid = $1
		GROUP BY f.uuid, f.title, f.banner, d.name, f.duration;`

const getAllFilmsPreviews = `
		SELECT f.uuid, f.title, f.banner, f.director, f.duration, AVG(c.score), COUNT(c.id)
		FROM film f
		LEFT JOIN comment c ON f.id = c.film
		GROUP BY f.uuid, f.title, f.banner, f.director, f.duration;`

const getAllFilmComments = `
		SELECT comment.uuid, film.uuid, users.name AS author_name, comment.text, comment.score, comment.added_at
		FROM comment
		JOIN users ON comment.author = users.id
		JOIN film ON comment.film = film.id
		WHERE film.uuid = $1;`

const getAllFilmActors = `
		SELECT a.uuid, a.name, a.avatar
		FROM actor a
		JOIN film_actor fa ON a.id = fa.actor
		JOIN film f ON fa.film = f.id
		WHERE f.uuid = $1;`

func (storage *FilmsStorage) GetFilmDataByUuid(uuid string) (domain.FilmData, error) {
	var film domain.FilmData
	err := storage.pool.QueryRow(context.Background(), getFilmDataByUuid, uuid).Scan(
		&film.Uuid,
		&film.Title,
		&film.Preview,
		&film.Director,
		&film.Data,
		&film.Duration,
		&film.Date,
		&film.AverageScore,
		&film.ScoresCount)
	if err != nil {
		return domain.FilmData{}, myerrors.ErrInternalServerError
	}
	return film, nil
}

func (storage *FilmsStorage) AddFilm(film domain.FilmDataToAdd) error {
	tx, err := storage.pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to begin transaction to add film: %w",
			myerrors.ErrInternalServerError)
	}
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			fmt.Printf("failed to rollback transaction to add film: %v",
				myerrors.ErrInternalServerError)
		}
	}()

	var directorFlag int
	err = storage.pool.QueryRow(context.Background(), getAmountOfDirectorsByName, film.Director).Scan(&directorFlag)
	if err != nil {
		return fmt.Errorf("failed to get amount of directors: %w",
			myerrors.ErrInternalServerError)
	}
	if directorFlag == 0 {
		_, err = storage.pool.Exec(context.Background(), insertDirector, film.Director)
		if err != nil {
			return fmt.Errorf("failed to insert director: %w",
				myerrors.ErrInternalServerError)
		}
	}

	var directorID int
	err = storage.pool.QueryRow(context.Background(), getDirectorsIdByName, film.Director).Scan(&directorID)
	if err != nil {
		return fmt.Errorf("failed to get directors id: %w",
			myerrors.ErrInternalServerError)
	}

	_, err = storage.pool.Exec(context.Background(), insertFilm, film.Title, film.Preview, directorID, film.Data,
		film.AgeLimit, film.Duration, film.PublishedAt)
	if err != nil {
		return fmt.Errorf("failed to insert film: %w",
			myerrors.ErrInternalServerError)
	}

	var filmID int
	err = storage.pool.QueryRow(context.Background(), getFilmIdByTitle, film.Title).Scan(&filmID)
	if err != nil {
		return fmt.Errorf("failed to get film id: %w",
			myerrors.ErrInternalServerError)
	}

	ActorsCast := film.Actors
	for _, actor := range ActorsCast {
		var actorFlag int
		err = storage.pool.QueryRow(context.Background(), getAmountOfActorsByName, actor.Name).Scan(&actorFlag)
		if err != nil {
			return fmt.Errorf("failed to get amount of actors: %w",
				myerrors.ErrInternalServerError)
		}
		if actorFlag == 0 {
			_, err = storage.pool.Exec(context.Background(), insertActor, actor.Name)
			if err != nil {
				return fmt.Errorf("failed to insert actor: %w",
					myerrors.ErrInternalServerError)
			}
		}

		var actorID int
		err = storage.pool.QueryRow(context.Background(), getActorId, actor.Name).Scan(&actorID)
		if err != nil {
			return fmt.Errorf("failed to get actor id: %w",
				myerrors.ErrInternalServerError)
		}

		_, err = storage.pool.Exec(context.Background(), insertIntoFilmActors, filmID, actorID)
		if err != nil {
			return fmt.Errorf("failed to insert film actors: %w",
				myerrors.ErrInternalServerError)
		}
	}

	return nil
}

func (storage *FilmsStorage) RemoveFilm(uuid string) error {
	_, err := storage.pool.Exec(context.Background(), deleteFilm, uuid)
	if err != nil {
		return myerrors.ErrInternalServerError
	}

	return nil
}

func (storage *FilmsStorage) GetFilmPreview(uuid string) (domain.FilmPreview, error) {
	var filmPreview domain.FilmPreview
	err := storage.pool.QueryRow(context.Background(), getFilmPreview, uuid).Scan(
		&filmPreview.Uuid,
		&filmPreview.Title,
		&filmPreview.Preview,
		&filmPreview.Director,
		&filmPreview.Duration,
		&filmPreview.AverageScore,
		&filmPreview.ScoresCount)
	if err != nil {
		return domain.FilmPreview{}, myerrors.ErrInternalServerError
	}
	return filmPreview, nil
}

func (storage *FilmsStorage) GetAllFilmsPreviews() ([]domain.FilmPreview, error) {
	rows, err := storage.pool.Query(context.Background(), getAllFilmsPreviews)
	if err != nil {
		return nil, myerrors.ErrInternalServerError
	}

	films := make([]domain.FilmPreview, 0)
	var (
		FilmUuid     string
		FilmPreview  string
		FilmTitle    string
		FilmDirector string
		FilmDuration int
		FilmScore    float32
		FilmRating   int
	)
	_, err = pgx.ForEachRow(rows,
		[]any{&FilmUuid, &FilmTitle, &FilmPreview, &FilmDirector, &FilmDuration, &FilmScore, &FilmRating}, func() error {
			film := domain.FilmPreview{
				Uuid:         FilmUuid,
				Title:        FilmTitle,
				Preview:      FilmPreview,
				Director:     FilmDirector,
				Duration:     FilmDuration,
				AverageScore: FilmScore,
				ScoresCount:  FilmRating,
			}

			films = append(films, film)

			return nil
		})
	if err != nil {
		return nil, myerrors.ErrInternalServerError
	}

	return films, nil
}

func (storage *FilmsStorage) GetAllFilmActors(uuid string) ([]domain.ActorPreview, error) {
	rows, err := storage.pool.Query(context.Background(), getAllFilmActors, uuid)
	if err != nil {
		return nil, myerrors.ErrInternalServerError
	}

	actors := make([]domain.ActorPreview, 0)
	var (
		ActorUuid   string
		ActorName   string
		ActorAvatar string
	)
	_, err = pgx.ForEachRow(rows, []any{&ActorUuid, &ActorName, &ActorAvatar}, func() error {
		actor := domain.ActorPreview{
			Uuid:   ActorUuid,
			Name:   ActorName,
			Avatar: ActorAvatar,
		}

		actors = append(actors, actor)

		return nil
	})
	if err != nil {
		return nil, myerrors.ErrInternalServerError
	}

	return actors, nil
}

func (storage *FilmsStorage) GetAllFilmComments(uuid string) ([]domain.Comment, error) {
	rows, err := storage.pool.Query(context.Background(), getAllFilmComments, uuid)
	if err != nil {
		return nil, myerrors.ErrInternalServerError
	}

	comments := make([]domain.Comment, 0)
	var (
		CommentUuid     string
		CommentFilmUuid string
		CommentAuthor   string
		CommentText     string
		CommentScore    int
		CommentAddedAt  time.Time
	)
	_, err = pgx.ForEachRow(rows,
		[]any{&CommentUuid, &CommentFilmUuid, &CommentAuthor, &CommentText, &CommentScore, &CommentAddedAt}, func() error {
			comment := domain.Comment{
				Uuid:     CommentUuid,
				FilmUuid: CommentFilmUuid,
				Author:   CommentAuthor,
				Text:     CommentText,
				Score:    CommentScore,
				AddedAt:  CommentAddedAt,
			}

			comments = append(comments, comment)

			return nil
		})
	if err != nil {
		return nil, myerrors.ErrInternalServerError
	}

	return comments, nil
}
