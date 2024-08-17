CREATE TABLE types
(
    id          character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    name        character varying(50)                           NOT NULL,
    description text,
    created_at  timestamp without time zone DEFAULT current_timestamp,
    updated_at  timestamp without time zone,
    deleted_at  timestamp without time zone,
    PRIMARY KEY (id)
)