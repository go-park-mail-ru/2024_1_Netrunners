CREATE DATABASE netrunnerflix;

CREATE TABLE IF NOT EXISTS users
(
	id            INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id   UUID UNIQUE DEFAULT gen_random_uuid()           NOT NULL,
	email         TEXT UNIQUE CHECK (LENGTH(email) <= 30)         NOT NULL,
	avatar        TEXT        DEFAULT 'https://shorturl.at/ewzP8' NOT NULL,
	name          TEXT        DEFAULT 'user'                      NOT NULL,
	password      TEXT CHECK (LENGTH(password) <= 64)             NOT NULL,
	registered_at TIMESTAMPTZ DEFAULT NOW()                       NOT NULL,
	birthday      TIMESTAMPTZ DEFAULT NOW()                       NOT NULL,
	is_admin      BOOLEAN     DEFAULT FALSE                       NOT NULL
);

CREATE TABLE IF NOT EXISTS actor
(
	id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id UUID UNIQUE                  DEFAULT gen_random_uuid()           NOT NULL,
	name        TEXT UNIQUE                                                      NOT NULL,
	avatar      TEXT                         DEFAULT 'https://shorturl.at/ewzP8' NOT NULL,
	birthday    TIMESTAMPTZ                  DEFAULT NOW()                       NOT NULL,
	career      TEXT                         DEFAULT ''                          NOT NULL,
	height      INTEGER CHECK (height < 300) DEFAULT 192                         NOT NULL,
	birth_place TEXT                         DEFAULT 'Russia, Angarsk'           NOT NULL,
	spouse      TEXT                         DEFAULT 'Светлана Ходченкова'       NOT NULL
);

CREATE TABLE IF NOT EXISTS director
(
	id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id UUID UNIQUE DEFAULT gen_random_uuid()           NOT NULL,
	name        TEXT UNIQUE                                     NOT NULL,
	avatar      TEXT        DEFAULT 'https://shorturl.at/ewzP8' NOT NULL,
	birthday    TIMESTAMPTZ DEFAULT NOW()                       NOT NULL
);

CREATE TABLE IF NOT EXISTS film
(
	id           INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id  UUID UNIQUE                                         DEFAULT gen_random_uuid()           NOT NULL,
<<<<<<< HEAD
	is_serial BOOLEAN NOT NULL,
=======
	is_serial    BOOLEAN                                                                                 NOT NULL,
>>>>>>> 3113de9e52b0931ff56bc01ba0ba64ad63c867a9
	title        TEXT                                                                                    NOT NULL,
	data         TEXT                                                DEFAULT ''                          NOT NULL,
	banner       TEXT                                                DEFAULT 'https://shorturl.at/akMR2' NOT NULL,
	s3_link      TEXT                                                DEFAULT 'https://shorturl.at/jHIMO' NOT NULL,
	director     INTEGER,
	age_limit    SMALLINT CHECK (age_limit >= 0 AND age_limit <= 18) DEFAULT 18                          NOT NULL,
	duration     SMALLINT CHECK (duration > 0)                       DEFAULT 143                         NOT NULL,
	published_at TIMESTAMPTZ                                         DEFAULT NOW()                       NOT NULL,
	FOREIGN KEY (director) REFERENCES director (id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS episode
(
<<<<<<< HEAD
	id           INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	number INTEGER NOT NULL,
	s3_link      TEXT                                                DEFAULT 'https://shorturl.at/jHIMO' NOT NULL,
=======
	id      INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	number  INTEGER                                  NOT NULL,
	s3_link TEXT DEFAULT 'https://shorturl.at/jHIMO' NOT NULL
>>>>>>> 3113de9e52b0931ff56bc01ba0ba64ad63c867a9
);

CREATE TABLE IF NOT EXISTS season
(
<<<<<<< HEAD
	film_id INTEGER NOT NULL,
	number INTEGER NOT NULL,
	episode_id INTEGER NOT NULL,
	FOREIGN KEY (film_id) REFERENCES film (id) ON DELETE SET NULL
	FOREIGN KEY (episode_id) REFERENCES episode (id) ON DELETE SET NULL
)
=======
	film_id    INTEGER NOT NULL,
	number     INTEGER NOT NULL,
	episode_id INTEGER NOT NULL,
	FOREIGN KEY (film_id) REFERENCES film (id) ON DELETE SET NULL,
	FOREIGN KEY (episode_id) REFERENCES episode (id) ON DELETE SET NULL
);
>>>>>>> 3113de9e52b0931ff56bc01ba0ba64ad63c867a9

CREATE TABLE IF NOT EXISTS comment
(
	id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id UUID UNIQUE DEFAULT gen_random_uuid()       NOT NULL,
	text        TEXT                                        NOT NULL,
	score       SMALLINT CHECK (score >= 0 AND score <= 10) NOT NULL,
	author      INTEGER,
	film        INTEGER                                     NOT NULL,
	added_at    TIMESTAMPTZ DEFAULT NOW()                   NOT NULL,
	FOREIGN KEY (author) REFERENCES users (id) ON DELETE SET NULL,
	FOREIGN KEY (film) REFERENCES film (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS film_actor
(
	id    INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	film  INTEGER NOT NULL,
	actor INTEGER NOT NULL,
	FOREIGN KEY (film) REFERENCES film (id) ON DELETE CASCADE,
	FOREIGN KEY (actor) REFERENCES actor (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS favorite_film
(
	id               INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	film_external_id UUID NOT NULL,
	user_external_id UUID NOT NULL,
	FOREIGN KEY (user_external_id) REFERENCES users (external_id) ON DELETE CASCADE,
	FOREIGN KEY (film_external_id) REFERENCES film (external_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS genre
(
	id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id UUID UNIQUE DEFAULT gen_random_uuid() NOT NULL,
	name        TEXT                                  NOT NULL
);

CREATE TABLE IF NOT EXISTS film_genres
(
	id                INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	film_external_id  UUID NOT NULL,
	genre_external_id UUID NOT NULL,
	FOREIGN KEY (genre_external_id) REFERENCES genre (external_id) ON DELETE CASCADE,
	FOREIGN KEY (film_external_id) REFERENCES film (external_id) ON DELETE CASCADE
);
