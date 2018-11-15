SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;

--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';



SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: clients; Type: TABLE; Schema: public; Owner: artesia
--
CREATE TABLE public.clients (
   id bigint NOT NULL,
   external_id text NOT NULL,
   secret text NOT NULL,
   redirect_uri text NOT NULL,
   user_data text NOT NULL
);

ALTER TABLE public.clients OWNER TO artesia;

--
-- Name: clients_id_seq; Type: SEQUENCE; Schema: public; Owner: artesia
--

CREATE SEQUENCE public.clients_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.clients_id_seq OWNER TO artesia;

--
-- Name: clients_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artesia
--

ALTER SEQUENCE public.clients_id_seq OWNED BY public.clients.id;

--
-- Name: clients id; Type: DEFAULT; Schema: public; Owner: artesia
--

ALTER TABLE ONLY public.clients ALTER COLUMN id SET DEFAULT nextval('public.clients_id_seq'::regclass);


--
-- Name: clients clients_pkey; Type: CONSTRAINT; Schema: public; Owner: artesia
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_pkey PRIMARY KEY (id);


--
-- Name: uq_external_id; Type: INDEX; Schema: public; Owner: artesia
--

CREATE UNIQUE INDEX uq_external_id ON public.clients USING btree (external_id);


--
-- Name: authorizations; Type: TABLE; Schema: public; Owner: artesia
--
CREATE TABLE public.authorizations (
   id bigint NOT NULL,
   client_id bigint NOT NULL,
   code text NOT NULL,
   expiration integer NOT NULL,
   scope text NOT NULL,
   redirect_uri text,
   external_id text NOT NULL,
   state_data text NOT NULL,
   secret text NOT NULL,
   created_at timestamp with time zone NOT NULL,
   user_data text NOT NULL
);

ALTER TABLE public.authorizations OWNER TO artesia;

--
-- Name: authorizations_id_seq; Type: SEQUENCE; Schema: public; Owner: artesia
--

CREATE SEQUENCE public.authorizations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authorizations_id_seq OWNER TO artesia;

--
-- Name: authorizations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artesia
--

ALTER SEQUENCE public.authorizations_id_seq OWNED BY public.authorizations.id;

--
-- Name: authorizations id; Type: DEFAULT; Schema: public; Owner: artesia
--

ALTER TABLE ONLY public.authorizations ALTER COLUMN id SET DEFAULT nextval('public.authorizations_id_seq'::regclass);


--
-- Name: uq_external_id; Type: INDEX; Schema: public; Owner: artesia
--

CREATE UNIQUE INDEX uq_external_id ON public.authorizations USING btree (external_id);


--
-- Name: access_tokens; Type: TABLE; Schema: public; Owner: artesia
--
CREATE TABLE public.access_tokens (
   id bigint NOT NULL,
   client_id bigint NOT NULL,
   authorize_id text NOT NULL,
   token text NOT NULL,
   refresh_token text NOT NULL,
   expiration integer NOT NULL,
   scope text NOT NULL,
   redirect_uri text NOT NULL,
   created_at timestamp with time zone NOT NULL
);

ALTER TABLE public.access_tokens OWNER TO artesia;

--
-- Name: access_tokens_id_seq; Type: SEQUENCE; Schema: public; Owner: artesia
--

CREATE SEQUENCE public.access_tokens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.access_tokens_id_seq OWNER TO artesia;

--
-- Name: access_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artesia
--

ALTER SEQUENCE public.access_tokens_id_seq OWNED BY public.access_tokens.id;

--
-- Name: access_tokens id; Type: DEFAULT; Schema: public; Owner: artesia
--

ALTER TABLE ONLY public.access_tokens ALTER COLUMN id SET DEFAULT nextval('public.access_tokens_id_seq'::regclass);


--
-- PostgreSQL database dump complete
--
