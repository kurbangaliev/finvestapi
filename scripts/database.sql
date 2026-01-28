--
-- PostgreSQL database dump
--

-- Dumped from database version 18.1 (Debian 18.1-1.pgdg13+2)
-- Dumped by pg_dump version 18.1

-- Started on 2026-01-28 20:16:47 UTC

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
-- TOC entry 220 (class 1259 OID 16386)
-- Name: news; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.news (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text,
    message_date character varying(20),
    message text,
    type bigint DEFAULT 1,
    media_link character varying(255),
    download_link character varying(255),
    status_id bigint DEFAULT 1,
    author_id bigint DEFAULT 0
);


ALTER TABLE public.news OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 16430)
-- Name: news_analytics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.news_analytics (
    news_id bigint DEFAULT 0,
    liked bigint DEFAULT 0,
    disliked bigint DEFAULT 0,
    viewed bigint DEFAULT 0
);


ALTER TABLE public.news_analytics OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16385)
-- Name: news_id_seq1; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.news_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.news_id_seq1 OWNER TO postgres;

--
-- TOC entry 3497 (class 0 OID 0)
-- Dependencies: 219
-- Name: news_id_seq1; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.news_id_seq1 OWNED BY public.news.id;


--
-- TOC entry 222 (class 1259 OID 16404)
-- Name: news_likes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.news_likes (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    news_id bigint DEFAULT 0,
    user_id bigint DEFAULT 0,
    type bigint DEFAULT 1
);


ALTER TABLE public.news_likes OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16402)
-- Name: news_likes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.news_likes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.news_likes_id_seq OWNER TO postgres;

--
-- TOC entry 3498 (class 0 OID 0)
-- Dependencies: 221
-- Name: news_likes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.news_likes_id_seq OWNED BY public.news_likes.id;


--
-- TOC entry 224 (class 1259 OID 16417)
-- Name: news_viewings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.news_viewings (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    news_id bigint DEFAULT 0,
    user_id bigint DEFAULT 0
);


ALTER TABLE public.news_viewings OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16415)
-- Name: news_viewings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.news_viewings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.news_viewings_id_seq OWNER TO postgres;

--
-- TOC entry 3499 (class 0 OID 0)
-- Dependencies: 223
-- Name: news_viewings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.news_viewings_id_seq OWNED BY public.news_viewings.id;


--
-- TOC entry 227 (class 1259 OID 16438)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    login character varying(255),
    password character varying(255),
    role character varying(255)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 16437)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3500 (class 0 OID 0)
-- Dependencies: 226
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3308 (class 2604 OID 16392)
-- Name: news id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news ALTER COLUMN id SET DEFAULT nextval('public.news_id_seq1'::regclass);


--
-- TOC entry 3312 (class 2604 OID 16407)
-- Name: news_likes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news_likes ALTER COLUMN id SET DEFAULT nextval('public.news_likes_id_seq'::regclass);


--
-- TOC entry 3316 (class 2604 OID 16420)
-- Name: news_viewings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news_viewings ALTER COLUMN id SET DEFAULT nextval('public.news_viewings_id_seq'::regclass);


--
-- TOC entry 3323 (class 2604 OID 16441)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3484 (class 0 OID 16386)
-- Dependencies: 220
-- Data for Name: news; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.news (id, created_at, updated_at, deleted_at, title, message_date, message, type, media_link, download_link, status_id, author_id) FROM stdin;
1	2026-01-28 19:25:39.541523+00	2026-01-28 19:25:39.541523+00	\N	Видео-презентации	2013-03-14	&lt;p&gt;Представляем Вашему вниманию видео-презентацию о компании &amp;laquo;Fortune Invest&amp;raquo;.&lt;/p&gt;	1	https://www.youtube.com/embed/ntN8oEwL35s	https://www.youtube.com/embed/ntN8oEwL35s	1	0
2	2026-01-28 19:41:20.39676+00	2026-01-28 19:41:20.39676+00	\N	Видео-презентации	2013-08-26	&lt;p&gt;erfqwer&lt;/p&gt;	1	https://www.youtube.com/embed/G7HyNt79tQ4	https://www.youtube.com/embed/G7HyNt79tQ4	1	0
3	2026-01-28 19:43:31.062945+00	2026-01-28 19:43:31.062945+00	\N	Брайан Трейси	2025-11-20	&lt;p&gt;Брайан Трейси: 5 способов инвестировать в себя. C чего начать саморазвитие.&lt;/p&gt;	1	https://www.youtube.com/embed/srQNyCdwNjQ	https://www.youtube.com/embed/srQNyCdwNjQ	1	0
4	2026-01-28 19:44:55.363986+00	2026-01-28 19:44:55.363986+00	\N	Брайан Трейси: Искусство заключения сделок	2025-12-12	&lt;p&gt;Предлагаем вашему вниманию вдохновляющую аудиокнигу от легендарного эксперта по продажам — Брайана Трейси. В ней он простым языком объясняет, как мыслит успешный продавец, как выстраивать доверие с клиентом и что действительно помогает довести сделку до победного финала. Это не просто советы — это инструменты, которые меняют подход к работе и дают сильный профессиональный заряд. Рекомендуем к прослушиванию всем, кто хочет уверенно развивать свои навыки, увеличивать результаты и двигаться вперёд вместе с Fortune Invest.&lt;/p&gt;	1	https://www.youtube.com/embed/ef2FsZJ4rcU	https://www.youtube.com/embed/ef2FsZJ4rcU	1	0
5	2026-01-28 19:46:32.21251+00	2026-01-28 19:46:32.21251+00	\N	Роберт Кийосаки — «Как стать богатым за 30 минут»: ключевые идеи финансового мышления	2025-12-08	&lt;p style="text-align: justify;" data-mce-style="text-align: justify;"&gt;Видео раскрывает ключевые принципы финансового мышления: умение создавать активы, управлять денежным потоком и принимать стратегические решения. Кийосаки подчёркивает важность дисциплины и ответственности в формировании устойчивого финансового будущего.&lt;/p&gt;&lt;p style="text-align: justify;" data-mce-style="text-align: justify;"&gt;Материал будет полезен всем сотрудникам, стремящимся укреплять финансовую грамотность и повышать качество решений в работе и жизни. Рекомендуем к просмотру всем сотрудникам Fortune Invest как источник вдохновения и полезных идей для развития финансового мышления.&lt;/p&gt;	1	https://www.youtube.com/embed/cTx--vmzJ5w	https://www.youtube.com/embed/cTx--vmzJ5w	1	0
6	2026-01-28 19:49:46.485849+00	2026-01-28 19:49:46.485849+00	\N	Как при любой зарплате стать богатым	2025-12-15	&lt;p style="text-align: justify;" data-mce-style="text-align: justify;"&gt;Видео раскрывает ключевые принципы финансового мышления: умение создавать активы, управлять денежным потоком и принимать стратегические решения. Кийосаки подчёркивает важность дисциплины и ответственности в формировании устойчивого финансового будущего.&lt;/p&gt;&lt;p style="text-align: justify;" data-mce-style="text-align: justify;"&gt;Материал будет полезен всем сотрудникам, стремящимся укреплять финансовую грамотность и повышать качество решений в работе и жизни. Рекомендуем к просмотру всем сотрудникам Fortune Invest как источник вдохновения и полезных идей для развития финансового мышления.&lt;/p&gt;	1	https://www.youtube.com/embed/4uwu8I-oDcM	https://www.youtube.com/embed/4uwu8I-oDcM	1	0
7	2026-01-28 19:50:49.716276+00	2026-01-28 19:50:49.716276+00	\N	Закрытие сделки: как уверенно подвести клиента к решению	2025-12-22	&lt;p style="text-align: justify;" data-mce-style="text-align: justify;"&gt;В новом видео поднимается один из самых важных вопросов финансового мышления: почему уровень дохода не всегда определяет уровень благосостояния.&amp;nbsp;Материал показывает, что финансовый результат формируется не размером зарплаты, а подходом к деньгам, привычками и качеством принимаемых решений.&lt;/p&gt;&lt;p style="text-align: justify;" data-mce-style="text-align: justify;"&gt;Видео акцентирует внимание на дисциплине, умении управлять расходами и выстраивать долгосрочную стратегию, которая позволяет постепенно создавать финансовую устойчивость независимо от стартовых условий.&lt;/p&gt;&lt;p style="text-align: justify;" data-mce-style="text-align: justify;"&gt;📌 Рекомендуем к просмотру всем сотрудникам Fortune Invest как источник практичных идей и мотивации для развития финансового мышления.&lt;/p&gt;	4			1	0
8	2026-01-28 19:53:42.285006+00	2026-01-28 19:53:42.285006+00	\N	Как всё успевать: тайм-менеджмент без перегрузки	2026-01-26	&lt;p class="p1"&gt;В этой подборке — два видео о том, как управлять временем, а не работать в режиме постоянной спешки.&lt;/p&gt;&lt;p class="p1"&gt;Брайан Трейси и практический фильм о тайм-менеджменте показывают, как расставлять приоритеты, фокусироваться на главном и повышать личную эффективность без выгорания.&lt;/p&gt;&lt;p class="p1"&gt;Коротко, по делу и с идеями, которые легко применить в ежедневной работе.&lt;/p&gt;&lt;p class="p1"&gt;📌 &lt;em&gt;Рекомендуем к просмотру всем, кто хочет успевать больше и управлять своим результатом осознанно.&lt;/em&gt;&lt;/p&gt;&lt;p&gt;&lt;br data-mce-bogus="1"&gt;&lt;/p&gt;&lt;p&gt;&lt;span contenteditable="false" data-mce-object="iframe" class="mce-preview-object mce-object-iframe" data-mce-p-allowfullscreen="allowfullscreen" data-mce-p-src="https://www.youtube.com/embed/r6Lvc_HyWCI"&gt;&lt;iframe width="320" height="195" src="https://www.youtube.com/embed/r6Lvc_HyWCI" allowfullscreen="allowfullscreen" frameborder="0"&gt;&lt;/iframe&gt;&lt;span class="mce-shim"&gt;&lt;/span&gt;&lt;/span&gt;&lt;span contenteditable="false" data-mce-object="iframe" class="mce-preview-object mce-object-iframe" data-mce-p-allowfullscreen="allowfullscreen" data-mce-p-src="https://www.youtube.com/embed/TlxOSSsTP90"&gt;&lt;iframe width="320" height="195" src="https://www.youtube.com/embed/TlxOSSsTP90" allowfullscreen="allowfullscreen" frameborder="0"&gt;&lt;/iframe&gt;&lt;span class="mce-shim"&gt;&lt;/span&gt;&lt;/span&gt;&lt;/p&gt;	4			1	0
9	2026-01-28 19:54:46.007563+00	2026-01-28 19:54:46.007563+00	\N	Обновление расчетных листов	2025-12-15	&lt;p&gt;Обновилась база данных расчетных листов за ноябрь 2025.&lt;/p&gt;	4			1	0
\.


--
-- TOC entry 3489 (class 0 OID 16430)
-- Dependencies: 225
-- Data for Name: news_analytics; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.news_analytics (news_id, liked, disliked, viewed) FROM stdin;
\.


--
-- TOC entry 3486 (class 0 OID 16404)
-- Dependencies: 222
-- Data for Name: news_likes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.news_likes (id, created_at, updated_at, deleted_at, news_id, user_id, type) FROM stdin;
\.


--
-- TOC entry 3488 (class 0 OID 16417)
-- Dependencies: 224
-- Data for Name: news_viewings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.news_viewings (id, created_at, updated_at, deleted_at, news_id, user_id) FROM stdin;
\.


--
-- TOC entry 3491 (class 0 OID 16438)
-- Dependencies: 227
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, created_at, updated_at, deleted_at, login, password, role) FROM stdin;
1	2026-01-28 18:56:27.992345+00	2026-01-28 18:56:27.992345+00	\N	admin	ISMvKXpXpadDiUoOSoAfww==	admin
\.


--
-- TOC entry 3501 (class 0 OID 0)
-- Dependencies: 219
-- Name: news_id_seq1; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.news_id_seq1', 9, true);


--
-- TOC entry 3502 (class 0 OID 0)
-- Dependencies: 221
-- Name: news_likes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.news_likes_id_seq', 1, false);


--
-- TOC entry 3503 (class 0 OID 0)
-- Dependencies: 223
-- Name: news_viewings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.news_viewings_id_seq', 1, false);


--
-- TOC entry 3504 (class 0 OID 0)
-- Dependencies: 226
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- TOC entry 3329 (class 2606 OID 16413)
-- Name: news_likes news_likes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news_likes
    ADD CONSTRAINT news_likes_pkey PRIMARY KEY (id);


--
-- TOC entry 3326 (class 2606 OID 16400)
-- Name: news news_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news
    ADD CONSTRAINT news_pkey PRIMARY KEY (id);


--
-- TOC entry 3332 (class 2606 OID 16428)
-- Name: news_viewings news_viewings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news_viewings
    ADD CONSTRAINT news_viewings_pkey PRIMARY KEY (id);


--
-- TOC entry 3335 (class 2606 OID 16446)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3324 (class 1259 OID 16401)
-- Name: idx_news_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_news_deleted_at ON public.news USING btree (deleted_at);


--
-- TOC entry 3327 (class 1259 OID 16414)
-- Name: idx_news_likes_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_news_likes_deleted_at ON public.news_likes USING btree (deleted_at);


--
-- TOC entry 3330 (class 1259 OID 16429)
-- Name: idx_news_viewings_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_news_viewings_deleted_at ON public.news_viewings USING btree (deleted_at);


--
-- TOC entry 3333 (class 1259 OID 16451)
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


-- Completed on 2026-01-28 20:16:47 UTC

--
-- PostgreSQL database dump complete
--
