CREATE TABLE IF NOT EXISTS public.records
(
    records_id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    long_url text COLLATE pg_catalog."default",
    short_url character varying(10) COLLATE pg_catalog."default",
    CONSTRAINT records_pkey PRIMARY KEY (records_id)
)