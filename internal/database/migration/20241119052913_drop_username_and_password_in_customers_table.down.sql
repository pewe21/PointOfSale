ALTER TABLE customers
    ADD COLUMN username  character varying(12) NOT NULL AFTER id,
    ADD COLUMN password character varying(255) NOT NULL AFTER username,
    ADD COLUMN email character varying(255) NOT NULL AFTER password,
    ADD COLUMN verified_at timestamp without time zone BEFORE created_at

