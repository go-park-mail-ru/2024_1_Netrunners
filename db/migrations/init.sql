CREATE DATABASE netrunnerflix;

CREATE SEQUENCE IF NOT EXISTS users_id_seq START 1;

CREATE TABLE IF NOT EXISTS users
(
	id            INTEGER PRIMARY KEY UNIQUE DEFAULT NEXTVAL('users_id_seq')     NOT NULL,
	uuid          UUID UNIQUE                DEFAULT gen_random_uuid()           NOT NULL,
	email         TEXT UNIQUE CHECK (LENGTH(email) <= 30)                        NOT NULL,
	avatar        TEXT                       DEFAULT 'https://shorturl.at/ewzP8' NOT NULL,
	name          TEXT                       DEFAULT 'user'                      NOT NULL,
	password      TEXT CHECK (LENGTH(email) <= 64)                               NOT NULL,
	registered_at TIMESTAMP                  DEFAULT NOW()                       NOT NULL,
	birthday      TIMESTAMP                  DEFAULT NOW()                       NOT NULL,
	is_admin      BOOLEAN                    DEFAULT FALSE                       NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS actor_id_seq START 1;

CREATE TABLE IF NOT EXISTS actor
(
	id          INTEGER PRIMARY KEY UNIQUE   DEFAULT NEXTVAL('actor_id_seq')     NOT NULL,
	uuid        UUID UNIQUE                  DEFAULT gen_random_uuid()           NOT NULL,
	name        TEXT UNIQUE                                                      NOT NULL,
	avatar      TEXT                         DEFAULT 'https://shorturl.at/ewzP8' NOT NULL,
	birthday    TIMESTAMP                    DEFAULT NOW()                       NOT NULL,
	career      TEXT                         DEFAULT ''                          NOT NULL,
	height      INTEGER CHECK (height < 300) DEFAULT 192                         NOT NULL,
	birth_place TEXT                         DEFAULT 'Russia, Angarsk'           NOT NULL,
	genres      TEXT                         DEFAULT 'Riddim'                    NOT NULL,
	spouse      TEXT                         DEFAULT 'Светлана Ходченкова'       NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS director_id_seq START 1;

CREATE TABLE IF NOT EXISTS director
(
	id       INTEGER PRIMARY KEY UNIQUE DEFAULT NEXTVAL('director_id_seq')  NOT NULL,
	uuid     UUID UNIQUE                DEFAULT gen_random_uuid()           NOT NULL,
	name     TEXT UNIQUE                                                    NOT NULL,
	avatar   TEXT                       DEFAULT 'https://shorturl.at/ewzP8' NOT NULL,
	birthday TIMESTAMP                  DEFAULT NOW()                       NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS film_id_seq START 1;

CREATE TABLE IF NOT EXISTS film
(
	id           INTEGER PRIMARY KEY UNIQUE                          DEFAULT NEXTVAL('film_id_seq')          NOT NULL,
	uuid         UUID UNIQUE                                         DEFAULT gen_random_uuid()               NOT NULL,
	title        TEXT                                                                                        NOT NULL,
	data         TEXT                                                DEFAULT ''                              NOT NULL,
	banner       TEXT                                                DEFAULT 'https://shorturl.at/akMR2'     NOT NULL,
	s3_link      TEXT                                                DEFAULT 'https://daimnefilm.hb.ru-msk.' ||
																			 'vkcs.cloud/Rick%20Roll.ia.mp4' NOT NULL,
	director     INTEGER,
	age_limit    SMALLINT CHECK (age_limit >= 0 AND age_limit <= 18) DEFAULT 18                              NOT NULL,
	duration     SMALLINT CHECK (duration > 0)                       DEFAULT 143                             NOT NULL,
	published_at TIMESTAMP                                           DEFAULT NOW()                           NOT NULL,
	FOREIGN KEY (director) REFERENCES director (id) ON DELETE SET NULL
);

CREATE SEQUENCE IF NOT EXISTS comment_id_seq START 1;

CREATE TABLE IF NOT EXISTS comment
(
	id       INTEGER PRIMARY KEY UNIQUE DEFAULT NEXTVAL('comment_id_seq') NOT NULL,
	uuid     UUID UNIQUE                DEFAULT gen_random_uuid()         NOT NULL,
	text     TEXT                                                         NOT NULL,
	score    SMALLINT CHECK (score >= 0 AND score <= 10)                  NOT NULL,
	author   INTEGER,
	film     INTEGER                                                      NOT NULL,
	added_at TIMESTAMP                  DEFAULT NOW()                     NOT NULL,
	FOREIGN KEY (author) REFERENCES users (id) ON DELETE SET NULL,
	FOREIGN KEY (film) REFERENCES film (id) ON DELETE CASCADE
);

CREATE SEQUENCE IF NOT EXISTS film_actor_id_seq START 1;

CREATE TABLE IF NOT EXISTS film_actor
(
	id    INTEGER PRIMARY KEY UNIQUE DEFAULT NEXTVAL('film_actor_id_seq') NOT NULL,
	film  INTEGER                                                         NOT NULL,
	actor INTEGER                                                         NOT NULL,
	FOREIGN KEY (film) REFERENCES film (id) ON DELETE CASCADE,
	FOREIGN KEY (actor) REFERENCES actor (id) ON DELETE CASCADE
);
