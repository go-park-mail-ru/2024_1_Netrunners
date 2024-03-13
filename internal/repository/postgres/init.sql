<<<<<<< HEAD
create database netrunnerflix

create table if not exists users (
    id serial primary key unique not null,
    uuid uuid default gen_random_uuid() unique not null,
    email varchar(30) unique not null,
    name text not null,
    password varchar(64) not null,
    registered_at timestamp default now() not null,
    birthday timestamp default now() not null,
    is_admin boolean default false not null,
    version smallint default 0
);

create table if not exists actors (
    id serial primary key unique not null,
    uuid uuid default gen_random_uuid() unique not null,
    name text not null,
    data text not null,
    birthday timestamp default now() not null
);

create table if not exists directors (
    id serial primary key unique not null,
    name text not null,
    birthday timestamp default now() not null
);

create table if not exists films (
    id serial primary key unique not null,
    uuid uuid default gen_random_uuid() unique not null,
    title text not null,
    director integer not null,
    data text not null,
    age_limit integer check (age_limit >= 0 and age_limit <= 18) not null,
    duration smallint check (duration > 0) not null,
    published_at timestamp default now() not null,
    foreign key (director) references directors (id)
);

create table if not exists comments (
    id serial primary key unique not null,
    uuid uuid default gen_random_uuid() unique not null,
    text text not null,
    score smallint check (score >= 0 and score <= 10) not null,
    author integer,
    film integer not null,
    added_at timestamp default now() not null,
    foreign key (author) references users (id) on delete set null,
    foreign key (film) references films (id) on delete cascade
);

create table if not exists film_actors (
    film integer not null,
    actor integer not null,
    foreign key (film) references films (id) on delete cascade,
    foreign key (actor) references actors (id) on delete cascade
);
=======
create table if not exists users
(
    id            serial primary key unique                  not null,
    uuid          uuid      default gen_random_uuid() unique not null,
    email         varchar(30) unique                         not null,
    name          text                                       not null,
    password      varchar(64)                                not null,
    registered_at timestamp default now()                    not null,
    birthday      timestamp default now()                    not null,
    is_admin      boolean   default false                    not null
);

create table if not exists actors
(
    id       serial primary key unique             not null,
    uuid     uuid default gen_random_uuid() unique not null,
    name     text                                  not null,
    data     text                                  not null,
    birthday timestamp default now()                         not null
);

create table if not exists directors
(
    id       serial primary key unique not null,
    name     text                      not null,
    birthday timestamp default now()   not null
);

create table if not exists films
(
    id           serial primary key unique                  not null,
    uuid         uuid      default gen_random_uuid() unique not null,
    title        text                                       not null,
    director     integer                                    not null,
    data         text                                       not null,
    duration     smallint check (duration > 0)              not null,
    published_at timestamp default now()                    not null,
    foreign key (director) references directors (id)
);

create table if not exists comments
(
    id       serial primary key unique                   not null,
    uuid     uuid      default gen_random_uuid() unique  not null,
    text     text                                        not null,
    score    smallint check (score >= 0 and score <= 10) not null,
    author   integer                                     not null,
    film     integer                                     not null,
    added_at timestamp default now()                     not null,
    foreign key (author) references users (id),
    foreign key (film) references films (id)
);

create table if not exists film_actors
(
    film  integer not null,
    actor integer not null,
    foreign key (film) references films (id),
    foreign key (actor) references actors (id)
);

ALTER TABLE actors
    ALTER COLUMN birthday SET DEFAULT NOW();
>>>>>>> 0b3f692 (users and films done)
