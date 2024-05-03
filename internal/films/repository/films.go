package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type PgxIface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close()
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

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

const getAmountOfFilmsByName = `
		SELECT COUNT(title) FROM film WHERE title = $1;`

const getAmountOfDirectorsByName = `
		SELECT COUNT(*)
		FROM director
		WHERE name = $1;`

const getDirectorIDByName = `
		SELECT id FROM director WHERE name = $1;`

const insertDirector = `
		INSERT INTO director (name, avatar, birthday) VALUES ($1, $2, $3) RETURNING id;`

const insertFilm = `
		INSERT INTO film (title, banner, director, data, age_limit, duration, published_at, s3_link) 
    	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, external_id;`

const getAmountOfActorsByName = `
		SELECT COUNT(*)
		FROM actor
		WHERE actor.name = $1;`

const insertActor = `
		INSERT INTO actor (name, avatar, career, birthday, birth_place, height, spouse) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

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

const getActorDataByUuid = `
		SELECT external_id, name, avatar, birthday, career, height, birth_place, spouse
		FROM actor
		WHERE external_id = $1;`

const getFilmsByActor = `
		SELECT f.external_id, f.title, f.banner, d.name, f.duration,
        	COALESCE(AVG(c.score), 0) AS avg_score, COALESCE(COUNT(c.id), 0) AS comment_count, f.age_limit
		FROM film f
		LEFT JOIN (film_actor fa LEFT JOIN actor a ON fa.actor = a.id) faa ON f.id = faa.film
		LEFT JOIN comment c ON f.id = c.film
		JOIN director d ON f.director = d.id
		WHERE faa.external_id = $1
		GROUP BY f.external_id, f.title, f.banner, d.name, f.duration, f.age_limit;`

const getActorsByFilm = `
		SELECT a.external_id, a.name, a.avatar
		FROM actor a 
		LEFT JOIN (film_actor fa LEFT JOIN film f ON fa.film = f.id) faf ON a.id = faf.actor
		WHERE faf.external_id = $1;`

const putFavoriteFilm = `
		INSERT INTO favorite_film (film_external_id, user_external_id) VALUES ($1, $2);`

const removeFavoriteFilm = `
		DELETE FROM favorite_film
		WHERE film_external_id = $1 AND user_external_id = $2;`

const getAmountOfUserByUuid = `
		SELECT COUNT(id)
		FROM users
		WHERE users.external_id = $1;`

const getAmountOfFilmByUuid = `
		SELECT COUNT(id)
		FROM film
		WHERE film.external_id = $1;`

const getAllFavoriteFilms = `
		SELECT f.external_id, f.title, f.banner, d.name, f.duration, 
		       COALESCE(AVG(c.score), 0) AS avg_score, COALESCE(COUNT(c.id), 0) AS comment_count, f.age_limit
		FROM film f
		INNER JOIN favorite_film fav ON f.external_id = fav.film_external_id
		LEFT JOIN comment c ON f.id = c.film
		JOIN director d ON f.director = d.id
		WHERE fav.user_external_id = $1
		GROUP BY f.external_id, f.title, f.banner, d.name, f.duration, f.age_limit;`

const getOneFavoriteByUuids = `
		SELECT film_external_id, user_external_id 
		FROM favorite_film 
		WHERE film_external_id = $1 AND user_external_id = $2;`

const getAllFilmsByGenreUuid = `
		SELECT f.external_id, f.title, f.banner, d.name, f.duration,
        	COALESCE(AVG(c.score), 0) AS avg_score, COALESCE(COUNT(c.id), 0) AS comment_count, f.age_limit
		FROM film f
		LEFT JOIN film_genres fg ON f.external_id = fg.film_external_id
		LEFT JOIN comment c ON f.id = c.film
		JOIN director d ON f.director = d.id
		WHERE fg.genre_external_id = $1
		GROUP BY f.external_id, f.title, f.banner, d.name, f.duration, f.age_limit;`

const getAllGenres = `
		SELECT g.external_id, g.name
		FROM genre g;`

const getGenresByFilm = `
		SELECT g.name, g.external_id
		FROM genre g
		LEFT JOIN film_genres fg ON fg.genre_external_id = g.external_id
		WHERE fg.film_external_id = $1;`

const getAmountOfGenresByName = `
		SELECT COUNT(id) 
		FROM genre 
		WHERE name = $1`

const insertGenre = `
		INSERT INTO genre (name) VALUES ($1) RETURNING external_id;`

const getGenreUuidByName = `
		SELECT external_id FROM genre WHERE name = $1;`

const insertFilmGenre = `
		INSERT INTO film_genres (film_external_id, genre_external_id) VALUES ($1, $2)`

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
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.FilmData{}, fmt.Errorf("%w", myerrors.ErrNotFound)
	}
	if err != nil {
		return domain.FilmData{}, fmt.Errorf("failed to get film data by uuid: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}

	genresRows, err := storage.pool.Query(context.Background(), getGenresByFilm, uuid)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.FilmData{}, fmt.Errorf("%w", myerrors.ErrNotFound)
	}
	if err != nil {
		return domain.FilmData{}, fmt.Errorf("failed to get film data by uuid: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}
	var genres []domain.Genre
	for genresRows.Next() {
		var genre domain.Genre
		err = genresRows.Scan(&genre.Name, &genre.Uuid)
		if err != nil {
			return domain.FilmData{}, fmt.Errorf("failed to get film data by uuid: %w: %w", err,
				myerrors.ErrFailInQueryRow)
		}
		genres = append(genres, genre)
	}
	film.Genres = genres

	return film, nil
}

func (storage *FilmsStorage) AddFilm(film domain.FilmToAdd) error {
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

	var (
		directorFlag int
		directorID   int
		filmID       int
		filmUuid     string
		filmFlag     int
	)
	err = tx.QueryRow(context.Background(), getAmountOfFilmsByName, film.FilmData.Title).Scan(&filmFlag)
	if err != nil {
		return fmt.Errorf("failed to get amount of directors: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}
	if filmFlag != 0 {
		return fmt.Errorf("failed to insert film: %w: %w", err,
			myerrors.ErrFilmAlreadyExists)
	}

	err = tx.QueryRow(context.Background(), getAmountOfDirectorsByName, film.DirectorToAdd.Name).Scan(&directorFlag)
	if err != nil {
		return fmt.Errorf("failed to get amount of directors: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}
	if directorFlag == 0 {
		err = tx.QueryRow(context.Background(), insertDirector,
			film.DirectorToAdd.Name, film.DirectorToAdd.Avatar, film.DirectorToAdd.Birthday).Scan(&directorID)
		if err != nil {
			return fmt.Errorf("failed to insert director: %w: %w", err,
				myerrors.ErrFailInExec)
		}
	} else {
		err = tx.QueryRow(context.Background(), getDirectorIDByName,
			film.DirectorToAdd.Name).Scan(&directorID)
		if err != nil {
			return fmt.Errorf("failed to get director id by name: %w: %w", err,
				myerrors.ErrFailInExec)
		}
	}

	err = tx.QueryRow(context.Background(), insertFilm, film.FilmData.Title, film.FilmData.Preview, directorID,
		film.FilmData.Data, film.FilmData.AgeLimit, film.FilmData.Duration,
		film.FilmData.PublishedAt, film.FilmData.Link).Scan(&filmID, &filmUuid)
	if err != nil {
		return fmt.Errorf("failed to insert film: %w: %w", err,
			myerrors.ErrFailInExec)
	}
	for _, genre := range film.FilmData.Genres {
		var (
			genreFlag int
			genreUuid string
		)
		err = tx.QueryRow(context.Background(), getAmountOfGenresByName, genre).Scan(&genreFlag)
		if err != nil {
			return fmt.Errorf("failed to amount of genres: %w: %w", err,
				myerrors.ErrFailInExec)
		}
		if genreFlag == 0 {
			err = tx.QueryRow(context.Background(), insertGenre, genre).Scan(&genreUuid)
			if err != nil {
				return fmt.Errorf("failed to insert genre: %w: %w", err,
					myerrors.ErrFailInExec)
			}
		} else {
			err = tx.QueryRow(context.Background(), getGenreUuidByName, genre).Scan(&genreUuid)
			if err != nil {
				return fmt.Errorf("failed to get genre uuid: %w: %w", err,
					myerrors.ErrFailInQueryRow)
			}
		}

		_, err = tx.Exec(context.Background(), insertFilmGenre, filmUuid, genreUuid)
		if err != nil {
			return fmt.Errorf("failed to insert film genre: %w: %w", err,
				myerrors.ErrFailInQueryRow)
		}
	}

	ActorsCast := film.Actors
	for _, actor := range ActorsCast {
		var actorFlag int
		err = tx.QueryRow(context.Background(), getAmountOfActorsByName, actor.Name).Scan(&actorFlag)
		if err != nil {
			return fmt.Errorf("failed to get amount of actors: %w: %w", err,
				myerrors.ErrFailInQueryRow)
		}
		var actorID int
		if actorFlag == 0 {
			err = tx.QueryRow(context.Background(), insertActor, actor.Name, actor.Avatar, actor.Career, actor.Birthday,
				actor.BirthPlace, actor.Height, actor.Spouse).Scan(&actorID)
			if err != nil {
				return fmt.Errorf("failed to insert actor: %w: %w", err,
					myerrors.ErrFailInExec)
			}
		} else {
			err = tx.QueryRow(context.Background(), getActorId, actor.Name).Scan(&actorID)
			if err != nil {
				return fmt.Errorf("failed to get actor id: %w: %w", err,
					myerrors.ErrFailInQueryRow)
			}
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
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.FilmPreview{}, fmt.Errorf("%w", myerrors.ErrNotFound)
	}
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
		filmDuration uint32
		filmScore    float32
		filmRating   uint64
		filmAgeLimit uint32
	)
	for rows.Next() {
		var film domain.FilmPreview
		err = rows.Scan(&filmUuid, &filmTitle, &filmPreview, &filmDirector, &filmDuration, &filmScore, &filmRating,
			&filmAgeLimit)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w", myerrors.ErrNotFound)
		}
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w", myerrors.ErrNotFound)
		}
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
		CommentScore    uint32
		CommentAddedAt  time.Time
	)
	for rows.Next() {
		err = rows.Scan(&CommentUuid, &CommentFilmUuid, &CommentAuthor, &CommentText, &CommentScore, &CommentAddedAt)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w", myerrors.ErrNotFound)
		}
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

func (storage *FilmsStorage) GetActorsByFilm(filmUuid string) ([]domain.ActorPreview, error) {
	rows, err := storage.pool.Query(context.Background(), getActorsByFilm, filmUuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get actors by film: %w: %w", err,
			myerrors.ErrFailInQuery)
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
		return nil, fmt.Errorf("failed to save actors by film: %w: %w", err,
			myerrors.ErrFailInForEachRow)
	}

	return actors, nil
}

func (storage *FilmsStorage) GetActorByUuid(actorUuid string) (domain.ActorData, error) {
	var actor = domain.ActorData{}
	err := storage.pool.QueryRow(context.Background(), getActorDataByUuid, actorUuid).Scan(
		&actor.Uuid,
		&actor.Name,
		&actor.Avatar,
		&actor.Birthday,
		&actor.Career,
		&actor.Height,
		&actor.BirthPlace,
		&actor.Spouse)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ActorData{}, fmt.Errorf("%w", myerrors.ErrNotFound)
	}
	if err != nil {
		return domain.ActorData{}, fmt.Errorf("failed to get actor by uuid: %w: %w", err,
			myerrors.ErrFailInQueryRow)
	}

	rows, err := storage.pool.Query(context.Background(), getFilmsByActor, actorUuid)
	if err != nil {
		return domain.ActorData{}, fmt.Errorf("failed to get actor's films: %w: %w", err,
			myerrors.ErrFailInQuery)
	}

	films := make([]domain.FilmPreview, 0)
	var (
		filmUuid     string
		filmPreview  string
		filmTitle    string
		filmDirector string
		filmDuration uint32
		filmScore    float32
		filmRating   uint64
		filmAgeLimit uint32
	)
	for rows.Next() {
		var film domain.FilmPreview
		err = rows.Scan(&filmUuid, &filmTitle, &filmPreview, &filmDirector, &filmDuration, &filmScore, &filmRating,
			&filmAgeLimit)
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ActorData{}, fmt.Errorf("%w", myerrors.ErrNotFound)
		}
		if err != nil {
			return domain.ActorData{}, err
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
	actor.Films = films

	return actor, nil
}

func (storage *FilmsStorage) PutFavoriteFilm(filmUuid string, userUuid string) error {
	var (
		amountOfUsers   int
		amountOfFilms   int
		filmUuidExisted string
		userUuidExisted string
	)
	err := storage.pool.QueryRow(context.Background(), getAmountOfUserByUuid, userUuid).Scan(&amountOfUsers)
	if err != nil {
		return err
	}
	if amountOfUsers == 0 {
		return fmt.Errorf("%w", myerrors.ErrNoSuchUser)
	}

	err = storage.pool.QueryRow(context.Background(), getAmountOfFilmByUuid, filmUuid).Scan(&amountOfFilms)
	if err != nil {
		return err
	}
	if amountOfFilms == 0 {
		return fmt.Errorf("%w", myerrors.ErrNoSuchUser)
	}

	err = storage.pool.QueryRow(context.Background(), getOneFavoriteByUuids, filmUuid, userUuid).Scan(&filmUuidExisted,
		&userUuidExisted)
	if errors.Is(err, pgx.ErrNoRows) {
		_, err = storage.pool.Exec(context.Background(), putFavoriteFilm, filmUuid, userUuid)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("%w", myerrors.ErrFavoriteAlreadyExists)
}

func (storage *FilmsStorage) RemoveFavoriteFilm(filmUuid string, userUuid string) error {
	var (
		amountOfUsers int
		amountOfFilms int
	)
	err := storage.pool.QueryRow(context.Background(), getAmountOfUserByUuid, userUuid).Scan(&amountOfUsers)
	if err != nil {
		return err
	}
	if amountOfUsers == 0 {
		return fmt.Errorf("%w", myerrors.ErrNoSuchUser)
	}

	err = storage.pool.QueryRow(context.Background(), getAmountOfFilmByUuid, filmUuid).Scan(&amountOfFilms)
	if err != nil {
		return err
	}
	if amountOfFilms == 0 {
		return fmt.Errorf("%w", myerrors.ErrNoSuchUser)
	}

	err = storage.pool.QueryRow(context.Background(), getAmountOfFilmByUuid, filmUuid).Scan(&amountOfFilms)
	if err != nil {
		return err
	}
	if amountOfFilms == 0 {
		return fmt.Errorf("%w", myerrors.ErrNoSuchUser)
	}
	_, err = storage.pool.Exec(context.Background(), removeFavoriteFilm, filmUuid, userUuid)
	if err != nil {
		return err
	}
	return nil
}

func (storage *FilmsStorage) GetAllFavoriteFilms(userUuid string) ([]domain.FilmPreview, error) {
	var (
		amountOfUsers int
	)
	err := storage.pool.QueryRow(context.Background(), getAmountOfUserByUuid, userUuid).Scan(&amountOfUsers)
	if err != nil {
		return nil, err
	}
	if amountOfUsers == 0 {
		return nil, fmt.Errorf("%w", myerrors.ErrNoSuchUser)
	}

	rows, err := storage.pool.Query(context.Background(), getAllFavoriteFilms, userUuid)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("%w, %s", myerrors.ErrNotFound, userUuid)
	} else if err != nil {
		return nil, err
	}

	films := make([]domain.FilmPreview, 0)
	var (
		FilmUuid     string
		FilmPreview  string
		FilmTitle    string
		FilmDirector string
		FilmDuration uint32
		FilmScore    float32
		FilmRating   uint64
		FilmAgeLimit uint32
	)
	for rows.Next() {
		var film domain.FilmPreview
		err = rows.Scan(&FilmUuid, &FilmTitle, &FilmPreview, &FilmDirector, &FilmDuration, &FilmScore, &FilmRating,
			&FilmAgeLimit)
		if err != nil {
			return nil, err
		}

		film.Uuid = FilmUuid
		film.Title = FilmTitle
		film.Preview = FilmPreview
		film.Director = FilmDirector
		film.Duration = FilmDuration
		film.ScoresCount = FilmRating
		film.AverageScore = FilmScore
		film.AgeLimit = FilmAgeLimit

		films = append(films, film)
	}
	return films, nil
}

func (storage *FilmsStorage) GetAllFilmsByGenre(genreUuid string) ([]domain.FilmPreview, error) {
	rows, err := storage.pool.Query(context.Background(), getAllFilmsByGenreUuid, genreUuid)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("%w, %s", myerrors.ErrNotFound, genreUuid)
	}
	if err != nil {
		return nil, err
	}

	var (
		films        []domain.FilmPreview
		FilmUuid     string
		FilmPreview  string
		FilmTitle    string
		FilmDirector string
		FilmDuration uint32
		FilmScore    float32
		FilmRating   uint64
		FilmAgeLimit uint32
	)
	for rows.Next() {
		var film domain.FilmPreview
		err = rows.Scan(&FilmUuid, &FilmTitle, &FilmPreview, &FilmDirector, &FilmDuration, &FilmScore, &FilmRating,
			&FilmAgeLimit)
		if err != nil {
			return nil, err
		}
		film.Uuid = FilmUuid
		film.Title = FilmTitle
		film.Preview = FilmPreview
		film.Director = FilmDirector
		film.Duration = FilmDuration
		film.ScoresCount = FilmRating
		film.AverageScore = FilmScore
		film.AgeLimit = FilmAgeLimit

		films = append(films, film)
	}
	return films, nil
}

func (storage *FilmsStorage) GetAllGenres() ([]domain.GenreFilms, error) {
	genreRows, err := storage.pool.Query(context.Background(), getAllGenres)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, myerrors.ErrNoGenres
	}
	if err != nil {
		return nil, err
	}

	var genresFilms []domain.GenreFilms
	for genreRows.Next() {
		var genreFilms domain.GenreFilms
		err = genreRows.Scan(&genreFilms.Uuid, &genreFilms.Name)
		if err != nil {
			return nil, err
		}
		filmsRows, err := storage.pool.Query(context.Background(), getAllFilmsByGenreUuid, genreFilms.Uuid)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w, %s", myerrors.ErrNotFound, genreFilms.Uuid)
		}
		if err != nil {
			return nil, err
		}
		films := make([]domain.FilmPreview, 0)
		var (
			FilmUuid     string
			FilmPreview  string
			FilmTitle    string
			FilmDirector string
			FilmDuration uint32
			FilmScore    float32
			FilmRating   uint64
			FilmAgeLimit uint32
		)
		i := 0
		for filmsRows.Next() && i < 4 {
			var film domain.FilmPreview
			err = filmsRows.Scan(&FilmUuid, &FilmTitle, &FilmPreview, &FilmDirector, &FilmDuration, &FilmScore,
				&FilmRating, &FilmAgeLimit)
			if err != nil {
				return nil, err
			}

			film.Uuid = FilmUuid
			film.Title = FilmTitle
			film.Preview = FilmPreview
			film.Director = FilmDirector
			film.Duration = FilmDuration
			film.ScoresCount = FilmRating
			film.AverageScore = FilmScore
			film.AgeLimit = FilmAgeLimit

			films = append(films, film)
			i++
		}
		genreFilms.Films = films
		genresFilms = append(genresFilms, genreFilms)
	}
	return genresFilms, nil
}
