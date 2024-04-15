package database

import (
	"context"

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

type ActorsStorage struct {
	pool PgxIface
}

func NewActorsStorage(pool PgxIface) (*ActorsStorage, error) {
	return &ActorsStorage{
		pool: pool,
	}, nil
}

const getActorDataByUuid = `
		SELECT external_id, name, avatar, birthday, career, height, birth_place, genres, spouse
		FROM actor
		WHERE external_id = $1;`

const getActorsFilms = `
		SELECT f.external_id, f.title
		FROM film f LEFT JOIN (film_actor fa LEFT JOIN actor a ON fa.actor = a.id) faa ON f.id = faa.film
		WHERE faa.external_id = $1;`

const getActorsByFilm = `
		SELECT a.external_id, a.name, a.avatar
		FROM actor a LEFT JOIN (film_actor fa LEFT JOIN film f ON fa.film = f.id) faf ON a.id = faf.actor
		WHERE faf.external_id = $1;`

func (storage *ActorsStorage) GetActorByUuid(actorUuid string) (domain.ActorData, error) {
	var actor domain.ActorData
	err := storage.pool.QueryRow(context.Background(), getActorDataByUuid, actorUuid).Scan(
		&actor.Uuid,
		&actor.Name,
		&actor.Avatar,
		&actor.Birthday,
		&actor.Career,
		&actor.Height,
		&actor.BirthPlace,
		&actor.Genres,
		&actor.Spouse)
	if err != nil {
		return domain.ActorData{}, myerrors.ErrInternalServerError
	}

	rows, err := storage.pool.Query(context.Background(), getActorsFilms, actorUuid)
	if err != nil {
		return domain.ActorData{}, myerrors.ErrInternalServerError
	}

	films := make([]domain.FilmLink, 0)
	var (
		filmUuid  string
		filmTitle string
	)
	_, err = pgx.ForEachRow(rows, []any{&filmUuid, &filmTitle}, func() error {
		film := domain.FilmLink{
			Uuid:  filmUuid,
			Title: filmTitle,
		}

		films = append(films, film)

		return nil
	})
	if err != nil {
		return domain.ActorData{}, myerrors.ErrInternalServerError
	}

	actor.Films = films

	return actor, nil
}

func (storage *ActorsStorage) GetActorsByFilm(filmUuid string) ([]domain.ActorPreview, error) {
	rows, err := storage.pool.Query(context.Background(), getActorsByFilm, filmUuid)
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
