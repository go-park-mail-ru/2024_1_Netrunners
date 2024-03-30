package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type FilmsStorage struct {
	pool *pgxpool.Pool
}

func NewFilmsStorage(pool *pgxpool.Pool) (*FilmsStorage, error) {
	return &FilmsStorage{
		pool: pool,
	}, nil
}

func (storage *FilmsStorage) GetFilmDataByUuid(uuid string) (domain.FilmData, error) {
	var film domain.FilmData
	err := storage.pool.QueryRow(context.Background(),
		`SELECT f.uuid, f.title, f.avatar, d.name, f.published_at, f.duration, AVG(c.score), COUNT(c.id)
			FROM film f
			LEFT JOIN comment c ON f.id = c.film
			JOIN director d ON f.director = d.id
			WHERE f.uuid = $1
			GROUP BY f.uuid, f.title,  f.avatar, d.name, f.published_at, f.duration;`, uuid).Scan(
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
		return domain.FilmData{},
			fmt.Errorf("error at begin transaction in GetFilmByUuid: %w", myerrors.ErrInternalServerError)
	}

	return film, nil
}

func (storage *FilmsStorage) AddFilm(film domain.FilmDataToAdd) error {
	tx, err := storage.pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return fmt.Errorf("error at begin transaction in AddFilm: %w",
			myerrors.ErrInternalServerError)
	}
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			fmt.Printf("error at rollback transaction in AddFilm: %v",
				myerrors.ErrInternalServerError)
		}
	}()

	var directorFlag int
	err = storage.pool.QueryRow(context.Background(),
		`SELECT COUNT(*)
			FROM director
			WHERE name = $1;`, film.Director).Scan(&directorFlag)
	if err != nil {
		return fmt.Errorf("error at inserting into data in AddFilm: %w",
			myerrors.ErrInternalServerError)
	}
	if directorFlag == 0 {
		_, err = storage.pool.Exec(context.Background(),
			`INSERT INTO director (name) VALUES ($1);`, film.Director)
		if err != nil {
			return fmt.Errorf("error at inserting into data in AddFilm: %w",
				myerrors.ErrInternalServerError)
		}
	}

	var directorID int
	err = storage.pool.QueryRow(context.Background(),
		`SELECT id
			FROM director
			WHERE name = $1;`, film.Director).Scan(&directorID)
	if err != nil {
		return fmt.Errorf("error at inserting into data in AddFilm: %w",
			myerrors.ErrInternalServerError)
	}

	_, err = storage.pool.Exec(context.Background(),
		`INSERT INTO film (title, avatar, director, data, age_limit, duration, published_at) 
    		VALUES ($1, $2, $3, $4, $5, $6, $7);`, film.Title, film.Preview, directorID, film.Data, film.AgeLimit,
		film.Duration, film.PublishedAt)
	if err != nil {
		return fmt.Errorf("error at inserting into data in AddFilm: %w",
			myerrors.ErrInternalServerError)
	}

	var filmID int
	err = storage.pool.QueryRow(context.Background(),
		`SELECT id
			FROM film
			WHERE title = $1;`, film.Title).Scan(&filmID)
	if err != nil {
		return fmt.Errorf("error at inserting into data in AddFilm: %w",
			myerrors.ErrInternalServerError)
	}

	ActorsCast := film.Actors
	for _, actor := range ActorsCast {
		var actorFlag int
		err = storage.pool.QueryRow(context.Background(),
			`SELECT COUNT(*)
			FROM actor
			WHERE actor.name = $1;`, actor.Name).Scan(&actorFlag)
		if err != nil {
			return fmt.Errorf("error at inserting into data in AddFilm: %w",
				myerrors.ErrInternalServerError)
		}
		if actorFlag == 0 {
			_, err = storage.pool.Exec(context.Background(),
				`INSERT INTO actor (name, data) VALUES ($1, $2);`, actor.Name, actor.Data)
			if err != nil {
				return fmt.Errorf("error at inserting into data in AddFilm: %w",
					myerrors.ErrInternalServerError)
			}
		}

		var actorID int
		err = storage.pool.QueryRow(context.Background(),
			`SELECT id
			FROM actor
			WHERE actor.name = $1;`, actor.Name).Scan(&actorID)
		if err != nil {
			return fmt.Errorf("error at inserting into data in AddFilm: %w",
				myerrors.ErrInternalServerError)
		}

		_, err = storage.pool.Exec(context.Background(),
			`INSERT INTO film_actor (film, actor) VALUES ($1, $2);`, filmID, actorID)
		if err != nil {
			return fmt.Errorf("error at inserting into data in AddFilm: %w",
				myerrors.ErrInternalServerError)
		}
	}

	return nil
}

func (storage *FilmsStorage) RemoveFilm(uuid string) error {
	_, err := storage.pool.Exec(context.Background(),
		`DELETE FROM film
			WHERE uuid = $1;`, uuid)
	if err != nil {
		return fmt.Errorf("error at inserting into data in RemoveFilm: %w", myerrors.ErrInternalServerError)
	}

	return nil
}

func (storage *FilmsStorage) GetFilmPreview(uuid string) (domain.FilmPreview, error) {
	var filmPreview domain.FilmPreview
	err := storage.pool.QueryRow(context.Background(),
		`SELECT f.uuid, f.title, f.avatar, d.name, f.duration, AVG(c.score), COUNT(c.id)
			FROM film f
			LEFT JOIN comment c ON f.id = c.film
			JOIN director d ON f.director = d.id
			WHERE f.uuid = $1
			GROUP BY f.uuid, f.title, f.avatar, d.name, f.duration;`, uuid).Scan(
		&filmPreview.Uuid,
		&filmPreview.Title,
		&filmPreview.Preview,
		&filmPreview.Director,
		&filmPreview.Duration,
		&filmPreview.AverageScore,
		&filmPreview.ScoresCount)
	if err != nil {
		return domain.FilmPreview{},
			fmt.Errorf("error at begin transaction in GetFilmPreview: %w", myerrors.ErrInternalServerError)
	}

	return filmPreview, nil
}

func (storage *FilmsStorage) GetAllFilmsPreviews() ([]domain.FilmPreview, error) {
	rows, err := storage.pool.Query(context.Background(),
		`SELECT f.uuid, f.title, f.avatar, f.director, f.duration, COUNT(c.id)
			FROM film f
			LEFT JOIN comment c ON f.id = c.film
			GROUP BY f.uuid, f.title, f.avatar, f.director, f.duration;`)
	if err != nil {
		return nil, fmt.Errorf("error at recieving data in GetAllFilmsPreviews: %w",
			myerrors.ErrInternalServerError)
	}

	films := make([]domain.FilmPreview, 0)
	var (
		FilmUuid     string
		FilmPreview  string
		FilmTitle    string
		FilmDirector string
		FilmDuration int
		// FilmScore    float32
		FilmRating int
	)
	_, err = pgx.ForEachRow(rows,
		[]any{&FilmUuid, &FilmTitle, &FilmPreview, &FilmDirector, &FilmDuration, &FilmRating}, func() error {
			film := domain.FilmPreview{
				Uuid:     FilmUuid,
				Title:    FilmTitle,
				Preview:  FilmPreview,
				Director: FilmDirector,
				Duration: FilmDuration,
				// AverageScore: FilmScore,
				ScoresCount: FilmRating,
			}

			films = append(films, film)

			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("error at recieving data in GetAllFilmsPreviews: %w",
			myerrors.ErrInternalServerError)
	}

	return films, nil
}

func (storage *FilmsStorage) GetAllFilmComments(uuid string) ([]domain.Comment, error) {
	rows, err := storage.pool.Query(context.Background(),
		`SELECT comment.uuid, users.name AS author_name, comment.text, comment.score, comment.added_at
			FROM comment
			JOIN users ON comment.author = users.id
			JOIN film ON comment.film = film.id
			WHERE film.uuid = $1;`, uuid)
	if err != nil {
		return nil,
			fmt.Errorf("error at recieving data in GetAllFilmComments: %w", myerrors.ErrInternalServerError)
	}

	comments := make([]domain.Comment, 0)
	var (
		CommentUuid    string
		CommentAuthor  string
		CommentText    string
		CommentScore   int
		CommentAddedAt time.Time
	)
	_, err = pgx.ForEachRow(rows,
		[]any{&CommentUuid, &CommentAuthor, &CommentText, &CommentScore, &CommentAddedAt}, func() error {
			comment := domain.Comment{
				Uuid:    CommentUuid,
				Author:  CommentAuthor,
				Text:    CommentText,
				Score:   CommentScore,
				AddedAt: CommentAddedAt,
			}

			comments = append(comments, comment)

			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("error at inserting into data in GetAllFilmComments: %w",
			myerrors.ErrInternalServerError)
	}

	return comments, nil
}

func (storage *FilmsStorage) GetAllFilmActors(uuid string) ([]domain.ActorPreview, error) {
	rows, err := storage.pool.Query(context.Background(),
		`SELECT a.uuid, a.name
			FROM actor a
			JOIN film_actor fa ON a.id = fa.actor
			JOIN film f ON fa.film = f.id
			WHERE f.uuid = $1;`, uuid)
	if err != nil {
		return nil,
			fmt.Errorf("error at recieving data in GetAllFilmComments: %w", myerrors.ErrInternalServerError)
	}

	actors := make([]domain.ActorPreview, 0)
	var (
		ActorUuid string
		ActorName string
	)
	_, err = pgx.ForEachRow(rows, []any{&ActorUuid, &ActorName}, func() error {
		actor := domain.ActorPreview{
			Uuid:   ActorUuid,
			Name:   ActorName,
			Avatar: ActorUuid,
		}

		actors = append(actors, actor)

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error at inserting into data in GetAllFilmActors: %w",
			myerrors.ErrInternalServerError)
	}

	return actors, nil
}
