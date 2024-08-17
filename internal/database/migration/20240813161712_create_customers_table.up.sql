CREATE TABLE customers
(
    id          character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    email       character varying(255)                          NOT NULL,
    username    character varying(12)                           NOT NULL,
    name        character varying(255)                          NOT NULL,
    password    character varying(255)                          NOT NULL,
    phone       character varying(15),
    address     text,
    verified_at timestamp without time zone,
    created_at  timestamp without time zone default current_timestamp,
    updated_at  timestamp without time zone,
    deleted_at  timestamp without time zone,
    PRIMARY KEY (id)
)