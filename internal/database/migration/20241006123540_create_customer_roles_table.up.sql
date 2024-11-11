CREATE TABLE customer_roles(
    id          character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    customer_id character varying(36)                           NOT NULL,
    role_id     character varying(36)                           NOT NULL,
    created_at  timestamp without time zone DEFAULT current_timestamp,
    updated_at  timestamp without time zone,
    deleted_at  timestamp without time zone,
    PRIMARY KEY (id),
    CONSTRAINT fk_role
        FOREIGN KEY (role_id)
            REFERENCES roles (id),
    CONSTRAINT fk_customer
        FOREIGN KEY (customer_id)
            REFERENCES customers (id)
)