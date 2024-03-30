package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	"github.com/jackc/pgx/v5"
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

func (storage *ActorsStorage) GetAllActorsPreviews() ([]domain.ActorPreview, error) {
	rows, err := storage.pool.Query(context.Background(),
		`SELECT uuid, name
		FROM actor;`)
	if err != nil {
		return nil, fmt.Errorf("error at recieving data in GetAllActorsPreviews: %w", myerrors.ErrInternalServerError)
	}

	actors := make([]domain.ActorPreview, 0)
	var (
		ActorUuid string
		ActorName string
	)
	_, err = pgx.ForEachRow(rows, []any{&ActorUuid, &ActorName}, func() error {
		actor := domain.ActorPreview{
			Uuid: ActorUuid,
			Name: ActorName,
		}

		actors = append(actors, actor)

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error at recieving data in GetAllActorsPreviews: %w", myerrors.ErrInternalServerError)
	}

	return actors, nil
}

func (storage *ActorsStorage) GetActorByUuid(actorUuid string) (domain.ActorData, error) {
	tx, err := storage.pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return domain.ActorData{},
			fmt.Errorf("error at begin transaction in GetActorByUuid: %w", myerrors.ErrInternalServerError)
	}

	var actor domain.ActorData
	err = tx.QueryRow(context.Background(),
		`SELECT uuid, avatar, name, data, birthday
		FROM actor
		WHERE uuid = $1;`, actorUuid).Scan(
		&actor.Uuid,
		&actor.Avatar,
		&actor.Name,
		&actor.Data,
		&actor.Birthday)
	if err != nil {
		return domain.ActorData{},
			fmt.Errorf("error at recieving data in GetActorByUuid: %w", myerrors.ErrInternalServerError)
	}

	rows, err := tx.Query(context.Background(),
		`SELECT f.uuid, f.title
		FROM film f LEFT JOIN (film_actor fa LEFT JOIN actor a ON fa.actor = a.id) faa ON f.id = faa.film
		WHERE faa.uuid = $1`, actorUuid)
	if err != nil {
		return domain.ActorData{},
			fmt.Errorf("error at recieving films in GetActorByUuid: %w", myerrors.ErrInternalServerError)
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
		return domain.ActorData{},
			fmt.Errorf("error at recieving films in GetActorByUuid: %w", myerrors.ErrInternalServerError)
	}

	actor.Films = films

	return actor, nil
}

func (storage *ActorsStorage) GetActorsByFilm(filmUuid string) ([]domain.ActorPreview, error) {
	rows, err := storage.pool.Query(context.Background(),
		`SELECT a.uuid, a.name
		FROM actor a LEFT JOIN (film_actor fa LEFT JOIN film f ON fa.film = f.id) faf ON a.id = faf.actor
		WHERE faf.uuid = $1;`, filmUuid)
	if err != nil {
		return nil, fmt.Errorf("error at recieving data in GetAllActorsPreviews: %w", myerrors.ErrInternalServerError)
	}

	actors := make([]domain.ActorPreview, 0)
	var (
		ActorUuid string
		ActorName string
	)
	_, err = pgx.ForEachRow(rows, []any{&ActorUuid, &ActorName}, func() error {
		actor := domain.ActorPreview{
			Uuid: ActorUuid,
			Name: ActorName,
		}

		actors = append(actors, actor)

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error at recieving data in GetAllActorsPreviews: %w", myerrors.ErrInternalServerError)
	}

	return actors, nil
}
