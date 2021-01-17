--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1 (Debian 13.1-1.pgdg100+1)
-- Dumped by pg_dump version 13.1

-- Started on 2021-01-16 00:57:43

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
-- TOC entry 201 (class 1259 OID 32769)
-- Name: Grupos; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public."Grupos" (
    "ID" bigint NOT NULL,
    "FechaCreacion" timestamp without time zone,
    "Nombre" text,
    "Eliminado" boolean
);


ALTER TABLE public."Grupos" OWNER TO "user";

--
-- TOC entry 202 (class 1259 OID 32777)
-- Name: Grupos_Usuarios; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public."Grupos_Usuarios" (
    id_usuario bigint NOT NULL,
    id_grupo bigint NOT NULL
);


ALTER TABLE public."Grupos_Usuarios" OWNER TO "user";

--
-- TOC entry 200 (class 1259 OID 24583)
-- Name: usuarios; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.usuarios (
    "ID" bigint NOT NULL,
    "fechaCreacion" timestamp without time zone NOT NULL,
    "Nombre" text NOT NULL,
    "Eliminado" boolean NOT NULL
);


ALTER TABLE public.usuarios OWNER TO "user";

--
-- TOC entry 2816 (class 2606 OID 32781)
-- Name: Grupos_Usuarios Grupos_Usuarios_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public."Grupos_Usuarios"
    ADD CONSTRAINT "Grupos_Usuarios_pkey" PRIMARY KEY (id_usuario, id_grupo);


--
-- TOC entry 2814 (class 2606 OID 32776)
-- Name: Grupos Grupos_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public."Grupos"
    ADD CONSTRAINT "Grupos_pkey" PRIMARY KEY ("ID");


--
-- TOC entry 2812 (class 2606 OID 24590)
-- Name: usuarios usuarios_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.usuarios
    ADD CONSTRAINT usuarios_pkey PRIMARY KEY ("ID");


-- Completed on 2021-01-16 00:57:43

--
-- PostgreSQL database dump complete
--

