--
-- PostgreSQL database dump
--

\restrict 54NxGSlzbcdwDWzwk4BZn5hNi5C3hQBO9vRCS905mZrJqRfqEuPaUi3FjDlYtpr

-- Dumped from database version 17.10
-- Dumped by pg_dump version 17.10

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- Name: application_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.application_status AS ENUM (
    'Applied',
    'Interview',
    'Rejected',
    'Offer'
);


ALTER TYPE public.application_status OWNER TO postgres;

--
-- Name: job_portal; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.job_portal AS ENUM (
    'LinkedIn',
    'Naukri',
    'Indeed',
    'Wellfound',
    'Company Careers',
    'Other'
);


ALTER TYPE public.job_portal OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: job_applications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.job_applications (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    company_name character varying(255) NOT NULL,
    job_title character varying(255) NOT NULL,
    job_link text,
    job_portal public.job_portal NOT NULL,
    location character varying(255),
    status public.application_status DEFAULT 'Applied'::public.application_status NOT NULL,
    applied_date date NOT NULL,
    notes text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.job_applications OWNER TO postgres;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: job_applications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.job_applications (id, user_id, company_name, job_title, job_link, job_portal, location, status, applied_date, notes, created_at, updated_at) FROM stdin;
dbc2afc3-1f25-45fe-9715-ef05a5b74e12	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Amazon	Backend Engineer	https://careers.amazon.com/jobs/backend-engineer	LinkedIn	Bangalore	Applied	2026-07-24	Applied through LinkedIn referral	2026-07-25 20:45:24.88328	2026-07-25 20:45:24.88328
07e738a0-4593-4208-93c8-a3bbfd6efd1b	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Cisco	Backend Engineer	https://careers.cisco.com/jobs/backend-engineer	LinkedIn	Bangalore	Offer	2026-07-24	Applied through LinkedIn referral	2026-07-25 22:55:40.95869	2026-07-29 14:14:10.763595
41f15497-f7de-4c06-a186-b4539679c9a2	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Fixezy	Backend Engineer	https://careers.cisco.com/jobs/backend-engineer	LinkedIn	Bangalore	Rejected	2026-07-24	Applied through LinkedIn referral	2026-07-28 22:55:26.312508	2026-07-29 14:20:45.085282
13423053-83ce-49a4-bc92-00a5f66dee60	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Bosh	Architect	https://careers.cisco.com/jobs/backend-engineer	Naukri	Bengaluru	Interview	2026-07-29	we	2026-07-29 14:29:13.593178	2026-07-29 14:29:28.29803
e84fad10-0126-45aa-a6c7-7b5c79f59bc8	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	MYSTARTUPWAVE PRIVATE LIMITED	Software Developer	https://in.indeed.com/viewjob?from=app-tracker-post_apply-appcard&hl=en&jk=2143a02f2f87adb2&tk=1jupj0vpfgilv800	Indeed	Bengaluru	Applied	2026-07-30	Salary range - ₹40,000 - ₹1,20,000 a month	2026-07-30 18:53:25.805932	2026-07-30 18:53:25.805932
deeb4c9d-099e-4bbc-8b67-0b99a76b86a3	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Vanilla Networks Pvt Ltd	Front-End Javascript Developer	https://vanillanetworks.co.in/current-vacancies.html	Indeed	Bengaluru	Applied	2026-07-30	Applied on carrer site	2026-07-30 19:02:47.449072	2026-07-30 19:02:47.449072
52154081-5ce2-4791-8079-07766f1cc798	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Netsoft	Flutter Developer	https://www.nestsoft.com/jobs-internship-kochi/jobs-flutter-developer-work-from-home#	Indeed	Remote	Applied	2026-07-30	Part time role	2026-07-30 19:09:31.306451	2026-07-30 19:09:31.306451
33de8793-4292-45a1-80c1-42a275e5db02	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	BIG HAPPY	Front End Developer	https://in.indeed.com/viewjob?jk=28aca5ad4926eecf&from=shareddesktop_copy	Indeed	Bengaluru	Applied	2026-07-30	Delhi alli idde	2026-07-30 19:13:15.328711	2026-07-30 19:13:15.328711
52ade182-247b-480d-b76e-c0aa97e460cd	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Tapetide	Claude Code Operator	https://in.indeed.com/viewjob?jk=8de265db228e870d&from=shareddesktop_copy	Indeed	Bengaluru	Applied	2026-07-30	Yep	2026-07-30 19:16:07.122126	2026-07-30 19:16:07.122126
cf32faf5-26bb-40dc-abac-8c3c596da9b9	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	PinnacleU	Front End Developer	https://in.indeed.com/viewjob?jk=7c46b9f4cd317045&from=shareddesktop_copy	Indeed	Bengaluru	Applied	2026-07-30	Salary range mentioned \n₹9,00,000 - ₹10,00,000 a year	2026-07-30 19:18:17.952694	2026-07-30 19:18:17.952694
bc38d359-a400-4999-9e4a-0237ab84713a	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Hanumant Technology IT Company & Training Institute	Full Stack Developer	https://in.indeed.com/viewjob?jk=3067b73b8b930c21&from=shareddesktop_copy	Indeed	Bengaluru	Applied	2026-07-30	Location - Lucknow, Uttar Pradesh	2026-07-30 19:20:27.387974	2026-07-30 19:20:27.387974
64ec1d9c-5b82-419b-8876-fb0f0c3a82ec	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Talent Support India	Fullstack Developer	https://www.naukri.com/job-listings-290726005626	Naukri	Bengaluru	Applied	2026-07-30	YEp	2026-07-30 22:49:31.890399	2026-07-30 22:49:31.890399
e6ec7027-1ef0-456f-ac72-387ba0c4aece	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	K12 Techno Services	Fullstack Engineer	https://www.naukri.com/job-listings-170726020563	Naukri	Bengaluru	Applied	2026-07-31	Python \nMachine learning	2026-07-31 14:57:56.578774	2026-07-31 14:57:56.578774
f70ca840-414b-4e04-8df0-dfb9a45c689e	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Mobignosis	Full Stack Developer	https://www.naukri.com/job-listings-210726015984	Naukri	Bengaluru	Applied	2026-07-31	CTC up to 6 LPA	2026-07-31 15:00:58.863218	2026-07-31 15:00:58.863218
2a159742-9ca2-44d7-9189-6ea201457c18	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Cumulations	React Developer	https://www.cumulations.com/job_listing/frontend-developer/	Naukri	Bengaluru	Applied	2026-07-31	Applied on company site	2026-07-31 15:15:33.838492	2026-07-31 15:15:33.838492
9b664c6d-0e54-4e61-bafe-c221c63bb897	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Determinant Studios	Back End Developer	https://www.naukri.com/job-listings-220226003742	Naukri	Bengaluru	Applied	2026-07-31	6-10 Lacs P.A.	2026-07-31 15:17:30.265635	2026-07-31 15:17:30.265635
e704bd61-68f4-4030-ba52-0fc43e796ea0	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Honeywell	Software Engineer I	https://careers.honeywell.com/en/sites/Honeywell/job/154771/?utm_medium=jobshare&utm_source=External+Job+Share	LinkedIn	Bengaluru	Applied	2026-08-03	Applied on Carrers site	2026-08-03 12:09:06.861838	2026-08-03 12:09:06.861838
16ca741f-0124-444f-874b-b0688c39a8a4	7227fdc1-b682-4fd4-95f0-7450a6ad2cca	iVA TEK PRIVATE LIMITED	Full-Stack / Backend Developer (Remote)	https://in.indeed.com/viewjob?jk=228ac13ed5174689&from=shareddesktop_copy	Indeed	Remote	Rejected	2026-07-30	Yep dude	2026-07-30 19:29:22.826142	2026-08-03 12:50:41.762559
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schema_migrations (version, dirty) FROM stdin;
2	f
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, password_hash, created_at, updated_at) FROM stdin;
7227fdc1-b682-4fd4-95f0-7450a6ad2cca	Sagu	sagar@example.com	$2a$10$FHU8PZy0S6Cp/SfROCwAEOwnD5J7lSZVJLBKd00yOy9qkGHVxQomq	2026-07-22 23:39:26.318227	2026-07-22 23:39:26.318227
3ae2898f-0cb6-4ea4-85a2-cb749a067641	Sagu2	sagar2@example.com	$2a$10$rPCaKNL4AEX3pFiJhqFUluk18YvOemMGwlN4weaw3Fql.PB7BjblG	2026-07-26 22:29:31.88273	2026-07-26 22:29:31.88273
\.


--
-- Name: job_applications job_applications_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.job_applications
    ADD CONSTRAINT job_applications_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_job_applications_applied_date; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_job_applications_applied_date ON public.job_applications USING btree (applied_date);


--
-- Name: idx_job_applications_company_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_job_applications_company_name ON public.job_applications USING btree (company_name);


--
-- Name: idx_job_applications_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_job_applications_status ON public.job_applications USING btree (status);


--
-- Name: idx_job_applications_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_job_applications_user_id ON public.job_applications USING btree (user_id);


--
-- Name: job_applications job_applications_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.job_applications
    ADD CONSTRAINT job_applications_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict 54NxGSlzbcdwDWzwk4BZn5hNi5C3hQBO9vRCS905mZrJqRfqEuPaUi3FjDlYtpr

