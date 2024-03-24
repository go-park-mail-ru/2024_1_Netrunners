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

func (storage *FilmsStorage) GetFilmByUuid(uuid string) (domain.FilmData, error) {
	var film domain.FilmData
	err := storage.pool.QueryRow(context.Background(),
		`select f.uuid, f.title, d.name, f.published_at, f.duration, avg(c.score), count(c.id)
			from films f
			left join comments c on f.id = c.film
			join directors d on f.director = d.id
			where f.uuid = $1
			group by f.uuid, f.title, d.name, f.duration;`, uuid).Scan(
		&film.Preview,
		&film.Title,
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

func (storage *FilmsStorage) Film(film domain.FilmDataToAdd) error {
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
		`select count(*)
			from directors
			where name = $1;`, film.Director).Scan(&directorFlag)
	if err != nil {
		return fmt.Errorf("error at inserting into data in AddFilm: %w",
			myerrors.ErrInternalServerError)
	}
	if directorFlag == 0 {
		_, err = storage.pool.Exec(context.Background(),
			`insert into directors (name) values ($1);`, film.Director)
		if err != nil {
			return fmt.Errorf("error at inserting into data in AddFilm: %w",
				myerrors.ErrInternalServerError)
		}
	}

	var directorID int
	err = storage.pool.QueryRow(context.Background(),
		`select id
			from directors
			where name = $1;`, film.Director).Scan(&directorID)
	if err != nil {
		return fmt.Errorf("error at inserting into data in AddFilm: %w",
			myerrors.ErrInternalServerError)
	}

	_, err = storage.pool.Exec(context.Background(),
		`insert into films (title, data, duration, director) 
    		values ($1, $2, $3, $4);`, film.Title, film.Data, film.Duration, directorID)
	if err != nil {
		return fmt.Errorf("error at inserting into data in AddFilm: %w",
			myerrors.ErrInternalServerError)
	}

	var filmID int
	err = storage.pool.QueryRow(context.Background(),
		`select id
			from films
			where title = $1;`, film.Title).Scan(&filmID)
	if err != nil {
		return fmt.Errorf("error at inserting into data in AddFilm: %w",
			myerrors.ErrInternalServerError)
	}

	ActorsCast := film.Actors
	for _, actor := range ActorsCast {
		var actorFlag int
		err = storage.pool.QueryRow(context.Background(),
			`select count(*)
			from actors
			where actors.name = $1;`, actor.Name).Scan(&actorFlag)
		if err != nil {
			return fmt.Errorf("error at inserting into data in AddFilm: %w",
				myerrors.ErrInternalServerError)
		}
		if actorFlag == 0 {
			_, err = storage.pool.Exec(context.Background(),
				`insert into actors (name, data) values ($1, $2);`, actor.Name, actor.Data)
			if err != nil {
				return fmt.Errorf("error at inserting into data in AddFilm: %w",
					myerrors.ErrInternalServerError)
			}
		}

		var actorID int
		err = storage.pool.QueryRow(context.Background(),
			`select id
			from actors
			where actors.name = $1;`, actor.Name).Scan(&actorID)
		if err != nil {
			return fmt.Errorf("error at inserting into data in AddFilm: %w",
				myerrors.ErrInternalServerError)
		}

		_, err = storage.pool.Exec(context.Background(),
			`insert into film_actors (film, actor) values ($1, $2);`, filmID, actorID)
		if err != nil {
			return fmt.Errorf("error at inserting into data in AddFilm: %w",
				myerrors.ErrInternalServerError)
		}
	}

	return nil
}

func (storage *FilmsStorage) RemoveFilm(uuid string) error {
	_, err := storage.pool.Exec(context.Background(),
		`delete from films
			where uuid = $1;`, uuid)
	if err != nil {
		return fmt.Errorf("error at inserting into data in RemoveFilm: %w", myerrors.ErrInternalServerError)
	}

	return nil
}

func (storage *FilmsStorage) GetFilmPreview(uuid string) (domain.FilmPreview, error) {
	var filmPreview domain.FilmPreview
	err := storage.pool.QueryRow(context.Background(),
		`select f.uuid, f.title, d.name, f.duration, avg(comments.score), count(comments.id)
			from films f
			left join comments c on f.id = c.film
			join directors d on f.director = d.id
			where f.uuid = $1
			group by f.uuid, f.title, d.name, f.duration;`, uuid).Scan(
		&filmPreview.Preview,
		&filmPreview.Title,
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
		`select f.uuid, f.title, f.director, f.duration, count(c.id)
			from films f
			left join comments c on f.id = c.film
			group by f.uuid, f.title, f.director, f.duration;`)
	if err != nil {
		return nil, fmt.Errorf("error at recieving data in GetAllFilmsPreviews: %w",
			myerrors.ErrInternalServerError)
	}

	films := make([]domain.FilmPreview, 0)
	var (
		FilmUuid     string
		FilmTitle    string
		FilmDirector string
		FilmDuration int
		// FilmScore    float32
		FilmRating int
	)
	_, err = pgx.ForEachRow(rows,
		[]any{&FilmUuid, &FilmTitle, &FilmDirector, &FilmDuration, &FilmRating}, func() error {
			film := domain.FilmPreview{
				Preview:  FilmUuid,
				Title:    FilmTitle,
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
		`select comments.uuid, users.name as author_name, comments.text, comments.score, comments.added_at
			from comments
			join users on comments.author = users.id
			join films on comments.film = films.id
			where film.uuid = $1;`, uuid)
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
		`select a.uuid, a.name
			from actors a
			join film_actors fa on a.id = fa.actor
			join films f on fa.film = f.id
			where f.uuid = $1;`, uuid)
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
