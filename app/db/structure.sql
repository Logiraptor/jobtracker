--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: goose_db_version; Type: TABLE; Schema: public; Owner: jobtracker; Tablespace: 
--

CREATE TABLE goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);


ALTER TABLE goose_db_version OWNER TO jobtracker;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE; Schema: public; Owner: jobtracker
--

CREATE SEQUENCE goose_db_version_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE goose_db_version_id_seq OWNER TO jobtracker;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jobtracker
--

ALTER SEQUENCE goose_db_version_id_seq OWNED BY goose_db_version.id;


--
-- Name: reset_tokens; Type: TABLE; Schema: public; Owner: jobtracker; Tablespace: 
--

CREATE TABLE reset_tokens (
    user_id integer NOT NULL,
    value text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE reset_tokens OWNER TO jobtracker;

--
-- Name: sessions; Type: TABLE; Schema: public; Owner: jobtracker; Tablespace: 
--

CREATE TABLE sessions (
    id integer NOT NULL,
    user_id integer,
    token text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE sessions OWNER TO jobtracker;

--
-- Name: sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: jobtracker
--

CREATE SEQUENCE sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sessions_id_seq OWNER TO jobtracker;

--
-- Name: sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jobtracker
--

ALTER SEQUENCE sessions_id_seq OWNED BY sessions.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: jobtracker; Tablespace: 
--

CREATE TABLE users (
    id integer NOT NULL,
    email text NOT NULL,
    password_hash text NOT NULL,
    current_token text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE users OWNER TO jobtracker;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: jobtracker
--

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users_id_seq OWNER TO jobtracker;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jobtracker
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: jobtracker
--

ALTER TABLE ONLY goose_db_version ALTER COLUMN id SET DEFAULT nextval('goose_db_version_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: jobtracker
--

ALTER TABLE ONLY sessions ALTER COLUMN id SET DEFAULT nextval('sessions_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: jobtracker
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Name: goose_db_version_pkey; Type: CONSTRAINT; Schema: public; Owner: jobtracker; Tablespace: 
--

ALTER TABLE ONLY goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);


--
-- Name: reset_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: jobtracker; Tablespace: 
--

ALTER TABLE ONLY reset_tokens
    ADD CONSTRAINT reset_tokens_pkey PRIMARY KEY (user_id);


--
-- Name: reset_tokens_value_key; Type: CONSTRAINT; Schema: public; Owner: jobtracker; Tablespace: 
--

ALTER TABLE ONLY reset_tokens
    ADD CONSTRAINT reset_tokens_value_key UNIQUE (value);


--
-- Name: sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: jobtracker; Tablespace: 
--

ALTER TABLE ONLY sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- Name: users_email_key; Type: CONSTRAINT; Schema: public; Owner: jobtracker; Tablespace: 
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users_pkey; Type: CONSTRAINT; Schema: public; Owner: jobtracker; Tablespace: 
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: reset_tokens_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jobtracker
--

ALTER TABLE ONLY reset_tokens
    ADD CONSTRAINT reset_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id);


--
-- Name: sessions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jobtracker
--

ALTER TABLE ONLY sessions
    ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

