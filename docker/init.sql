--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2
-- Dumped by pg_dump version 15.2

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

--
-- Name: alto; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE alto WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE alto OWNER TO postgres;

\connect alto

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

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    "createdAt" date DEFAULT '2024-06-30'::date NOT NULL,
    "updatedAt" date DEFAULT '2024-06-30'::date NOT NULL,
    "deletedAt" date,
    "firstName" character varying NOT NULL,
    role numeric NOT NULL,
    "middleName" character varying NOT NULL,
    "lastName" character varying NOT NULL,
    email character varying NOT NULL,
    "passwordHash" character varying NOT NULL,
    salt character varying NOT NULL,
    active boolean NOT NULL,
    "blockedAt" date
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (id, "createdAt", "updatedAt", "deletedAt", "firstName", role, "middleName", "lastName", email, "passwordHash", salt, active, "blockedAt") FROM stdin;
3bdd29d4-c344-4910-8b85-78ba4815b99f	2024-06-30	2024-06-30	\N	Dmitrii	0	Sergeevich	Totskii	dimatocky@gmail.com	asd	asd	f	\N
cc64d22d-8f7e-4d5d-8376-a24c407db626	2024-06-30	2024-06-30	\N	Dmitrii	0	Sergeevich	Totskii	dimatocky@gmail.com	asd	asd	f	\N
1c24fdc0-6131-4f51-b7d0-f6afe0c17fbb	2024-06-30	2024-06-30	\N	Dmitrii	0	Sergeevich	Totskii	dimatocky@gmail.com	asd	asd	f	\N
\.


--
-- Name: user PK_cace4a159ff9f2512dd42373760; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT "PK_cace4a159ff9f2512dd42373760" PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

