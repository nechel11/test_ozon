CREATE TABLE IF NOT EXISTS records
(
	records_id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
	long_url text,
	short_url character varying(10),
	CONSTRAINT records_pkey PRIMARY KEY (records_id)
);