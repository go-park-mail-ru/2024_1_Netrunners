--
-- PostgreSQL database dump
--

-- Dumped from database version 15.7 (Debian 15.7-1.pgdg120+1)
-- Dumped by pg_dump version 15.7 (Debian 15.7-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: actor; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.actor (
    id integer NOT NULL,
    external_id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    avatar text DEFAULT 'https://shorturl.at/ewzP8'::text NOT NULL,
    birthday timestamp with time zone DEFAULT now() NOT NULL,
    career text DEFAULT ''::text NOT NULL,
    height integer DEFAULT 192 NOT NULL,
    birth_place text DEFAULT 'Russia, Angarsk'::text NOT NULL,
    spouse text DEFAULT 'Светлана Ходченкова'::text NOT NULL,
    CONSTRAINT actor_height_check CHECK ((height < 300))
);


ALTER TABLE public.actor OWNER TO postgres;

--
-- Name: actor_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.actor ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.actor_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: comment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comment (
    id integer NOT NULL,
    external_id uuid DEFAULT gen_random_uuid() NOT NULL,
    text text NOT NULL,
    score smallint NOT NULL,
    author_external_id uuid NOT NULL,
    film_external_id uuid NOT NULL,
    added_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT comment_score_check CHECK (((score >= 0) AND (score <= 10)))
);


ALTER TABLE public.comment OWNER TO postgres;

--
-- Name: comment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.comment ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.comment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: director; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.director (
    id integer NOT NULL,
    external_id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    avatar text DEFAULT 'https://shorturl.at/ewzP8'::text NOT NULL,
    birthday timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.director OWNER TO postgres;

--
-- Name: director_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.director ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.director_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: episode; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.episode (
    id integer NOT NULL,
    number integer NOT NULL,
    title text NOT NULL,
    s3_link text DEFAULT 'https://shorturl.at/jHIMO'::text NOT NULL
);


ALTER TABLE public.episode OWNER TO postgres;

--
-- Name: episode_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.episode ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.episode_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: favorite_film; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.favorite_film (
    id integer NOT NULL,
    film_external_id uuid NOT NULL,
    user_external_id uuid NOT NULL
);


ALTER TABLE public.favorite_film OWNER TO postgres;

--
-- Name: favorite_film_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.favorite_film ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.favorite_film_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: film; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.film (
    id integer NOT NULL,
    external_id uuid DEFAULT gen_random_uuid() NOT NULL,
    is_serial boolean DEFAULT false NOT NULL,
    title text NOT NULL,
    data text DEFAULT ''::text NOT NULL,
    banner text DEFAULT 'https://shorturl.at/akMR2'::text NOT NULL,
    s3_link text DEFAULT 'https://shorturl.at/jHIMO'::text NOT NULL,
    director integer,
    age_limit smallint DEFAULT 18 NOT NULL,
    duration smallint DEFAULT 143 NOT NULL,
    published_at timestamp with time zone DEFAULT now() NOT NULL,
    with_subscription boolean DEFAULT false NOT NULL,
    CONSTRAINT film_age_limit_check CHECK (((age_limit >= 0) AND (age_limit <= 18))),
    CONSTRAINT film_duration_check CHECK ((duration > 0))
);


ALTER TABLE public.film OWNER TO postgres;

--
-- Name: film_actor; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.film_actor (
    id integer NOT NULL,
    film integer NOT NULL,
    actor integer NOT NULL
);


ALTER TABLE public.film_actor OWNER TO postgres;

--
-- Name: film_actor_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.film_actor ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.film_actor_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: film_genres; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.film_genres (
    id integer NOT NULL,
    film_external_id uuid NOT NULL,
    genre_external_id uuid NOT NULL
);


ALTER TABLE public.film_genres OWNER TO postgres;

--
-- Name: film_genres_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.film_genres ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.film_genres_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: film_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.film ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.film_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: genre; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.genre (
    id integer NOT NULL,
    external_id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.genre OWNER TO postgres;

--
-- Name: genre_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.genre ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.genre_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: season; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.season (
    film_id integer NOT NULL,
    number integer NOT NULL,
    episode_id integer NOT NULL
);


ALTER TABLE public.season OWNER TO postgres;

--
-- Name: subscription; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.subscription (
    id integer NOT NULL,
    title text NOT NULL,
    amount numeric NOT NULL,
    description text NOT NULL,
    duration integer NOT NULL
);


ALTER TABLE public.subscription OWNER TO postgres;

--
-- Name: subscription_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.subscription ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.subscription_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    external_id uuid DEFAULT gen_random_uuid() NOT NULL,
    email text NOT NULL,
    avatar text DEFAULT '/9j/4AAQSkZJRgABAQEBLAEsAAD/4QCaRXhpZgAASUkqAAgAAAADAA4BAgBQAAAAMgAAABoBBQABAAAAggAAABsBBQABAAAAigAAAAAAAABEZWZhdWx0IEF2YXRhciBQcm9maWxlIEljb24gVmVjdG9yLiBTb2NpYWwgTWVkaWEgVXNlciBJbWFnZS4gVmVjdG9yIElsbHVzdHJhdGlvbiwBAAABAAAALAEAAAEAAAD/4QVraHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLwA8P3hwYWNrZXQgYmVnaW49Iu+7vyIgaWQ9Ilc1TTBNcENlaGlIenJlU3pOVGN6a2M5ZCI/Pgo8eDp4bXBtZXRhIHhtbG5zOng9ImFkb2JlOm5zOm1ldGEvIj4KCTxyZGY6UkRGIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53My5vcmcvMTk5OS8wMi8yMi1yZGYtc3ludGF4LW5zIyI+CgkJPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6cGhvdG9zaG9wPSJodHRwOi8vbnMuYWRvYmUuY29tL3Bob3Rvc2hvcC8xLjAvIiB4bWxuczpJcHRjNHhtcENvcmU9Imh0dHA6Ly9pcHRjLm9yZy9zdGQvSXB0YzR4bXBDb3JlLzEuMC94bWxucy8iICAgeG1sbnM6R2V0dHlJbWFnZXNHSUZUPSJodHRwOi8veG1wLmdldHR5aW1hZ2VzLmNvbS9naWZ0LzEuMC8iIHhtbG5zOmRjPSJodHRwOi8vcHVybC5vcmcvZGMvZWxlbWVudHMvMS4xLyIgeG1sbnM6cGx1cz0iaHR0cDovL25zLnVzZXBsdXMub3JnL2xkZi94bXAvMS4wLyIgIHhtbG5zOmlwdGNFeHQ9Imh0dHA6Ly9pcHRjLm9yZy9zdGQvSXB0YzR4bXBFeHQvMjAwOC0wMi0yOS8iIHhtbG5zOnhtcFJpZ2h0cz0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL3JpZ2h0cy8iIHBob3Rvc2hvcDpDcmVkaXQ9IkdldHR5IEltYWdlcyIgR2V0dHlJbWFnZXNHSUZUOkFzc2V0SUQ9IjEzMzcxNDQxNDYiIHhtcFJpZ2h0czpXZWJTdGF0ZW1lbnQ9Imh0dHBzOi8vd3d3LmlzdG9ja3Bob3RvLmNvbS9sZWdhbC9saWNlbnNlLWFncmVlbWVudD91dG1fbWVkaXVtPW9yZ2FuaWMmYW1wO3V0bV9zb3VyY2U9Z29vZ2xlJmFtcDt1dG1fY2FtcGFpZ249aXB0Y3VybCIgPgo8ZGM6Y3JlYXRvcj48cmRmOlNlcT48cmRmOmxpPk1hcmlhIFNoYXBpbG92YTwvcmRmOmxpPjwvcmRmOlNlcT48L2RjOmNyZWF0b3I+PGRjOmRlc2NyaXB0aW9uPjxyZGY6QWx0PjxyZGY6bGkgeG1sOmxhbmc9IngtZGVmYXVsdCI+RGVmYXVsdCBBdmF0YXIgUHJvZmlsZSBJY29uIFZlY3Rvci4gU29jaWFsIE1lZGlhIFVzZXIgSW1hZ2UuIFZlY3RvciBJbGx1c3RyYXRpb248L3JkZjpsaT48L3JkZjpBbHQ+PC9kYzpkZXNjcmlwdGlvbj4KPHBsdXM6TGljZW5zb3I+PHJkZjpTZXE+PHJkZjpsaSByZGY6cGFyc2VUeXBlPSdSZXNvdXJjZSc+PHBsdXM6TGljZW5zb3JVUkw+aHR0cHM6Ly93d3cuaXN0b2NrcGhvdG8uY29tL3Bob3RvL2xpY2Vuc2UtZ20xMzM3MTQ0MTQ2LT91dG1fbWVkaXVtPW9yZ2FuaWMmYW1wO3V0bV9zb3VyY2U9Z29vZ2xlJmFtcDt1dG1fY2FtcGFpZ249aXB0Y3VybDwvcGx1czpMaWNlbnNvclVSTD48L3JkZjpsaT48L3JkZjpTZXE+PC9wbHVzOkxpY2Vuc29yPgoJCTwvcmRmOkRlc2NyaXB0aW9uPgoJPC9yZGY6UkRGPgo8L3g6eG1wbWV0YT4KPD94cGFja2V0IGVuZD0idyI/Pgr/7QCWUGhvdG9zaG9wIDMuMAA4QklNBAQAAAAAAHocAlAAD01hcmlhIFNoYXBpbG92YRwCeABQRGVmYXVsdCBBdmF0YXIgUHJvZmlsZSBJY29uIFZlY3Rvci4gU29jaWFsIE1lZGlhIFVzZXIgSW1hZ2UuIFZlY3RvciBJbGx1c3RyYXRpb24cAm4ADEdldHR5IEltYWdlc//bAEMACgcHCAcGCggICAsKCgsOGBAODQ0OHRUWERgjHyUkIh8iISYrNy8mKTQpISIwQTE0OTs+Pj4lLkRJQzxINz0+O//CAAsIAmQCZAEBEQD/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'::text NOT NULL,
    name text DEFAULT 'user'::text NOT NULL,
    password text NOT NULL,
    registered_at timestamp with time zone DEFAULT now() NOT NULL,
    birthday timestamp with time zone DEFAULT now() NOT NULL,
    is_admin boolean DEFAULT false NOT NULL,
    subscription_end_date timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT users_email_check CHECK ((length(email) <= 30)),
    CONSTRAINT users_password_check CHECK ((length(password) <= 64))
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: actor actor_external_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.actor
    ADD CONSTRAINT actor_external_id_key UNIQUE (external_id);


--
-- Name: actor actor_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.actor
    ADD CONSTRAINT actor_name_key UNIQUE (name);


--
-- Name: actor actor_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.actor
    ADD CONSTRAINT actor_pkey PRIMARY KEY (id);


--
-- Name: comment comment_external_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_external_id_key UNIQUE (external_id);


--
-- Name: comment comment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_pkey PRIMARY KEY (id);


--
-- Name: director director_external_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.director
    ADD CONSTRAINT director_external_id_key UNIQUE (external_id);


--
-- Name: director director_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.director
    ADD CONSTRAINT director_name_key UNIQUE (name);


--
-- Name: director director_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.director
    ADD CONSTRAINT director_pkey PRIMARY KEY (id);


--
-- Name: episode episode_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.episode
    ADD CONSTRAINT episode_pkey PRIMARY KEY (id);


--
-- Name: favorite_film favorite_film_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorite_film
    ADD CONSTRAINT favorite_film_pkey PRIMARY KEY (id);


--
-- Name: film_actor film_actor_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_actor
    ADD CONSTRAINT film_actor_pkey PRIMARY KEY (id);


--
-- Name: film film_external_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film
    ADD CONSTRAINT film_external_id_key UNIQUE (external_id);


--
-- Name: film_genres film_genres_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_genres
    ADD CONSTRAINT film_genres_pkey PRIMARY KEY (id);


--
-- Name: film film_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film
    ADD CONSTRAINT film_pkey PRIMARY KEY (id);


--
-- Name: genre genre_external_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.genre
    ADD CONSTRAINT genre_external_id_key UNIQUE (external_id);


--
-- Name: genre genre_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.genre
    ADD CONSTRAINT genre_pkey PRIMARY KEY (id);


--
-- Name: subscription subscription_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subscription
    ADD CONSTRAINT subscription_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_external_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_external_id_key UNIQUE (external_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: comment comment_author_external_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_author_external_id_fkey FOREIGN KEY (author_external_id) REFERENCES public.users(external_id) ON DELETE SET NULL;


--
-- Name: comment comment_film_external_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_film_external_id_fkey FOREIGN KEY (film_external_id) REFERENCES public.film(external_id) ON DELETE CASCADE;


--
-- Name: favorite_film favorite_film_film_external_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorite_film
    ADD CONSTRAINT favorite_film_film_external_id_fkey FOREIGN KEY (film_external_id) REFERENCES public.film(external_id) ON DELETE CASCADE;


--
-- Name: favorite_film favorite_film_user_external_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorite_film
    ADD CONSTRAINT favorite_film_user_external_id_fkey FOREIGN KEY (user_external_id) REFERENCES public.users(external_id) ON DELETE CASCADE;


--
-- Name: film_actor film_actor_actor_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_actor
    ADD CONSTRAINT film_actor_actor_fkey FOREIGN KEY (actor) REFERENCES public.actor(id) ON DELETE CASCADE;


--
-- Name: film_actor film_actor_film_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_actor
    ADD CONSTRAINT film_actor_film_fkey FOREIGN KEY (film) REFERENCES public.film(id) ON DELETE CASCADE;


--
-- Name: film film_director_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film
    ADD CONSTRAINT film_director_fkey FOREIGN KEY (director) REFERENCES public.director(id) ON DELETE SET NULL;


--
-- Name: film_genres film_genres_film_external_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_genres
    ADD CONSTRAINT film_genres_film_external_id_fkey FOREIGN KEY (film_external_id) REFERENCES public.film(external_id) ON DELETE CASCADE;


--
-- Name: film_genres film_genres_genre_external_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_genres
    ADD CONSTRAINT film_genres_genre_external_id_fkey FOREIGN KEY (genre_external_id) REFERENCES public.genre(external_id) ON DELETE CASCADE;


--
-- Name: season season_episode_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.season
    ADD CONSTRAINT season_episode_id_fkey FOREIGN KEY (episode_id) REFERENCES public.episode(id) ON DELETE SET NULL;


--
-- Name: season season_film_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.season
    ADD CONSTRAINT season_film_id_fkey FOREIGN KEY (film_id) REFERENCES public.film(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--