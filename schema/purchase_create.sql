-- Table: public.purchase

-- DROP TABLE public.purchase;

CREATE TABLE public.purchase
(
    id bigint NOT NULL DEFAULT nextval('purchase_id_seq'::regclass),
    "cardNo" character(16)[] COLLATE pg_catalog."default" NOT NULL,
    "expDate" date NOT NULL,
    amt money NOT NULL,
    "isoCurrency" character(3)[] COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT purchase_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.purchase
    OWNER to postgres;
	