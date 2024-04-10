package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"time"

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
		SELECT uuid, name, avatar, birthday, career, height, birth_place, genres, spouse
		FROM actor
		WHERE uuid = $1;
`

const getActorsFilms = `
		SELECT f.uuid, f.title
		FROM film f LEFT JOIN (film_actor fa LEFT JOIN actor a ON fa.actor = a.id) faa ON f.id = faa.film
		WHERE faa.uuid = $1;
`

const getActorsByFilm = `
		SELECT a.uuid, a.name, a.avatar
		FROM actor a LEFT JOIN (film_actor fa LEFT JOIN film f ON fa.film = f.id) faf ON a.id = faf.actor
		WHERE faf.uuid = $1;
`

func (storage *ActorsStorage) GetActorByUuid(actorUuid string) (domain.ActorData, error) {
	var actor = domain.ActorData{
		Uuid:       "3e712cfc-29c2-490b-aafd-e24be2141ebc",
		Name:       "Стас Ярушин",
		Avatar:     "https://www.film.ru/sites/default/files/people/1587604-1727100.jpg",
		Birthday:   time.Now(),
		Career:     "Универ",
		Height:     192,
		BirthPlace: "Ангарск",
		Spouse:     "Светлана Ходченков",
		Genres:     "Дабстеп",
	}

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
