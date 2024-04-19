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
		SELECT f.external_id, f.title, f.banner, f.s3_link, d.name, f.data, f.duration, f.published_at, 
		       COALESCE(AVG(c.score), 0) AS avg_score, COALESCE(COUNT(c.id), 0) AS comment_count, age_limit
		FROM film f
		LEFT JOIN comment c ON f.id = c.film
		JOIN director d ON f.director = d.id
		WHERE f.external_id = $1
		GROUP BY f.external_id, f.title,  f.banner, d.name, f.published_at, f.s3_link, f.data, f.duration, f.age_limit;`

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
		WHERE external_id = $1;`

const getFilmPreview = `
		SELECT f.external_id, f.title, f.banner, d.name, f.duration,
        	COALESCE(AVG(c.score), 0) AS avg_score, COALESCE(COUNT(c.id), 0) AS comment_count, f.age_limit
		FROM film f
		LEFT JOIN comment c ON f.id = c.film
		JOIN director d ON f.director = d.id
		WHERE f.external_id = $1
		GROUP BY f.external_id, f.title, f.banner, d.name, f.duration, f.age_limit;`

const getAllFilmsPreviews = `
    SELECT f.external_id, f.title, f.banner, d.name, f.duration,
        COALESCE(AVG(c.score), 0) AS avg_score, COALESCE(COUNT(c.id), 0) AS comment_count, f.age_limit
    FROM film f
    LEFT JOIN comment c ON f.id = c.film
    JOIN director d ON f.director = d.id
    GROUP BY f.external_id, f.title, f.banner, d.name, f.duration, f.age_limit;`

const getAllFilmComments = `
		SELECT c.external_id, f.external_id, u.name AS author_name, c.text, c.score, 
		       c.added_at
		FROM comment c
		JOIN users u ON c.author = u.id
		JOIN film f ON c.film = f.id
		WHERE f.external_id = $1;`

const getAllFilmActors = `
		SELECT a.external_id, a.name, a.avatar
		FROM actor a
		JOIN film_actor fa ON a.id = fa.actor
		JOIN film f ON fa.film = f.id
		WHERE f.external_id = $1;`

func (storage *FilmsStorage) GetFilmDataByUuid(uuid string) (domain.FilmData, error) {
	var film domain.FilmData
	err := storage.pool.QueryRow(context.Background(), getFilmDataByUuid, uuid).Scan(
		&film.Uuid,
		&film.Title,
		&film.Preview,
		&film.Link,
		&film.Director,
		&film.Data,
		&film.Duration,
		&film.Date,
		&film.AverageScore,
		&film.ScoresCount,
		&film.AgeLimit)
	if err != nil {
		return domain.FilmData{}, fmt.Errorf("failed to get film data by uuid: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}

	return film, nil
}

func (storage *FilmsStorage) AddFilm(film domain.FilmDataToAdd) error {
	tx, err := storage.pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to begin transaction to add film: %w: %w", err,
			myerrors.ErrFailedToBeginTransaction)
	}
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			fmt.Printf("failed to rollback transaction to add film: %v", err)
		}
	}()

	var directorFlag int
	err = tx.QueryRow(context.Background(), getAmountOfDirectorsByName, film.Director).Scan(&directorFlag)
	if err != nil {
		return fmt.Errorf("failed to get amount of directors: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}
	if directorFlag == 0 {
		_, err = tx.Exec(context.Background(), insertDirector, film.Director)
		if err != nil {
			return fmt.Errorf("failed to insert director: %w: %w", err,
				myerrors.ErrFailInExec)
		}
	}

	var directorID int
	err = tx.QueryRow(context.Background(), getDirectorsIdByName, film.Director).Scan(&directorID)
	if err != nil {
		return fmt.Errorf("failed to get directors id: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}

	_, err = tx.Exec(context.Background(), insertFilm, film.Title, film.Preview, directorID, film.Data,
		film.AgeLimit, film.Duration, film.PublishedAt)
	if err != nil {
		return fmt.Errorf("failed to insert film: %w: %w", err,
			myerrors.ErrFailInExec)
	}

	var filmID int
	err = tx.QueryRow(context.Background(), getFilmIdByTitle, film.Title).Scan(&filmID)
	if err != nil {
		return fmt.Errorf("failed to get film id: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}

	ActorsCast := film.Actors
	for _, actor := range ActorsCast {
		var actorFlag int
		err = tx.QueryRow(context.Background(), getAmountOfActorsByName, actor.Name).Scan(&actorFlag)
		if err != nil {
			return fmt.Errorf("failed to get amount of actors: %w: %w", err,
				myerrors.ErrFailInQueryRow)
		}
		if actorFlag == 0 {
			_, err = tx.Exec(context.Background(), insertActor, actor.Name)
			if err != nil {
				return fmt.Errorf("failed to insert actor: %w: %w", err,
					myerrors.ErrFailInExec)
			}
		}

		var actorID int
		err = tx.QueryRow(context.Background(), getActorId, actor.Name).Scan(&actorID)
		if err != nil {
			return fmt.Errorf("failed to get actor id: %w: %w", err,
				myerrors.ErrFailInQueryRow)
		}

		_, err = tx.Exec(context.Background(), insertIntoFilmActors, filmID, actorID)
		if err != nil {
			return fmt.Errorf("failed to insert film actors: %w: %w", err,
				myerrors.ErrFailInExec)
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w: %w", err,
			myerrors.ErrFailedToCommitTransaction)
	}

	return nil
}

func (storage *FilmsStorage) RemoveFilm(uuid string) error {
	_, err := storage.pool.Exec(context.Background(), deleteFilm, uuid)
	if err != nil {
		return fmt.Errorf("failed to remove film: %w: %w", err,
			myerrors.ErrFailInExec)
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
		&filmPreview.ScoresCount,
		&filmPreview.AgeLimit)
	if err != nil {
		return domain.FilmPreview{}, fmt.Errorf("failed to get film's preview: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}
	return filmPreview, nil
}

func (storage *FilmsStorage) GetAllFilmsPreviews() ([]domain.FilmPreview, error) {
	rows, err := storage.pool.Query(context.Background(), getAllFilmsPreviews)
	if err != nil {
		return nil, fmt.Errorf("failed to get all films' previews: %w: %w", err,
			myerrors.ErrInternalServerError)
	}

	films := make([]domain.FilmPreview, 0)
	var (
		filmUuid     string
		filmPreview  string
		filmTitle    string
		filmDirector string
		filmDuration int
		filmScore    float32
		filmRating   int
		filmAgeLimit uint8
	)
	for rows.Next() {
		var film domain.FilmPreview
		err = rows.Scan(&filmUuid, &filmTitle, &filmPreview, &filmDirector, &filmDuration, &filmScore, &filmRating,
			&filmAgeLimit)
		if err != nil {
			return nil, err
		}

		film.Uuid = filmUuid
		film.Title = filmTitle
		film.Preview = filmPreview
		film.Director = filmDirector
		film.Duration = filmDuration
		film.ScoresCount = filmRating
		film.AverageScore = filmScore
		film.AgeLimit = filmAgeLimit

		films = append(films, film)
	}

	return films, nil
}

func (storage *FilmsStorage) GetAllFilmActors(uuid string) ([]domain.ActorPreview, error) {
	rows, err := storage.pool.Query(context.Background(), getAllFilmActors, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get all film's actors: %w: %w", err,
			myerrors.ErrFailInQuery)
	}

	actors := make([]domain.ActorPreview, 0)
	var (
		ActorUuid   string
		ActorName   string
		ActorAvatar string
	)

	for rows.Next() {
		var actor domain.ActorPreview
		err = rows.Scan(&ActorUuid, &ActorName, &ActorAvatar)
		if err != nil {
			return nil, err
		}

		actor.Uuid = ActorUuid
		actor.Name = ActorName
		actor.Avatar = ActorAvatar

		actors = append(actors, actor)
	}

	return actors, nil
}

func (storage *FilmsStorage) GetAllFilmComments(uuid string) ([]domain.Comment, error) {
	rows, err := storage.pool.Query(context.Background(), getAllFilmComments, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get all film's comments: %w: %w", err,
			myerrors.ErrFailInQuery)
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
	for rows.Next() {
		err = rows.Scan(&CommentUuid, &CommentFilmUuid, &CommentAuthor, &CommentText, &CommentScore, &CommentAddedAt)
		if err != nil {
			return nil, err
		}
		var comment domain.Comment

		comment.Uuid = CommentUuid
		comment.Author = CommentAuthor
		comment.FilmUuid = CommentFilmUuid
		comment.Text = CommentText
		comment.Score = CommentScore
		comment.AddedAt = CommentAddedAt

		comments = append(comments, comment)
	}

	return comments, nil
}
