CREATE TABLE IF NOT EXISTS profile_question
(
	id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id UUID UNIQUE DEFAULT gen_random_uuid() NOT NULL,
	question    TEXT                                  NOT NULL
);

CREATE TABLE IF NOT EXISTS additional_question
(
	id                   INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id          UUID UNIQUE DEFAULT gen_random_uuid() NOT NULL,
	question_external_id UUID                                  NOT NULL,
	question             TEXT                                  NOT NULL
);

CREATE TABLE IF NOT EXISTS additional_answer
(
	id                   INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id          UUID UNIQUE DEFAULT gen_random_uuid() NOT NULL,
	id_inside_question   INTEGER                               NOT NULL,
	question_external_id UUID                                  NOT NULL,
	variant              TEXT                                  NOT NULL
);


CREATE TABLE IF NOT EXISTS profile_stat
(
	id                   INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	question_external_id UUID    NOT NULL,
	score                INTEGER NOT NULL,
	is_additional_score  BOOL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS film_data_question
(
	id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id UUID UNIQUE DEFAULT gen_random_uuid() NOT NULL,
	question    TEXT                                  NOT NULL
);

CREATE TABLE IF NOT EXISTS film_data_stat
(
	id                   INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	question_external_id UUID    NOT NULL,
	score                INTEGER NOT NULL,
	is_additional_score  BOOL DEFAULT FALSE

);

CREATE TABLE IF NOT EXISTS actor_question
(
	id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id UUID UNIQUE DEFAULT gen_random_uuid() NOT NULL,
	question    TEXT                                  NOT NULL
);

CREATE TABLE IF NOT EXISTS actor_stat
(
	id                   INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	question_external_id UUID    NOT NULL,
	score                INTEGER NOT NULL,
	is_additional_score  BOOL DEFAULT FALSE

);

CREATE TABLE IF NOT EXISTS film_question
(
	id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	external_id UUID UNIQUE DEFAULT gen_random_uuid() NOT NULL,
	question    TEXT                                  NOT NULL
);

CREATE TABLE IF NOT EXISTS film_stat
(
	id                   INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	question_external_id UUID    NOT NULL,
	score                INTEGER NOT NULL,
	is_additional_score  BOOL DEFAULT FALSE
);