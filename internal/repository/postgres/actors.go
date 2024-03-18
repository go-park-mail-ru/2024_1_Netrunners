package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type ActorsStorage struct {
	pool *pgxpool.Pool
}

func NewActorsStorage() (*ActorsStorage, error) {
	pool, err := pgxpool.New(context.Background(), fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"root1234",
		"netrunnerflix",
	))
	if err != nil {
		return nil, fmt.Errorf("error at connecting to database: %w", myerrors.ErrInternalServerError)
	}

	return &ActorsStorage{
		pool: pool,
	}, nil
}

func (storage *ActorsStorage) GetAllActorsPreviews() ([]domain.ActorPreview, error) {
	rows, err := storage.pool.Query(context.Background(),
		`select uuid, name
		from actors;`)
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
		return domain.ActorData{}, fmt.Errorf("error at begin transaction in GetActorByUuid: %w", myerrors.ErrInternalServerError)
	}
	defer tx.Rollback(context.Background())

	var actor domain.ActorData
	err = tx.QueryRow(context.Background(),
		`select uuid, name, data, birthday
		from actors
		where uuid = $1;`, actorUuid).Scan(
		&actor.Uuid,
		&actor.Name,
		&actor.Data,
		&actor.Birthday)
	if err != nil {
		return domain.ActorData{}, fmt.Errorf("error at recieving data in GetActorByUuid: %w", myerrors.ErrInternalServerError)
	}

	rows, err := storage.pool.Query(context.Background(),
		`select f.uuid, f.title
		from films f left join (film_actors fa left join actors a on fa.actor = a.id) faa on f.id = faa.film
		where faa.uuid = $1`, actorUuid)
	if err != nil {
		return domain.ActorData{}, fmt.Errorf("error at recieving films in GetActorByUuid: %w", myerrors.ErrInternalServerError)
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
		return domain.ActorData{}, fmt.Errorf("error at recieving films in GetActorByUuid: %w", myerrors.ErrInternalServerError)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return domain.ActorData{}, fmt.Errorf("error at commit transaction in GetActorByUuid: %w", myerrors.ErrInternalServerError)
	}

	actor.Films = films

	return actor, nil
}

func (storage *ActorsStorage) GetActorsByFilm(filmUuid string) ([]domain.ActorPreview, error) {
	rows, err := storage.pool.Query(context.Background(),
		`select a.uuid, a.name
		from actors a left join (film_actors fa left join films f on fa.film = f.id) faf on a.id = faf.actor
		where faf.uuid = $1;`, filmUuid)
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
