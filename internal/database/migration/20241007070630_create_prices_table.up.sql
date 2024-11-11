CREATE TABLE prices
(
    id               character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    customer_role_id character varying(36)                           NOT NULL,
    product_id       character varying(36)                           NOT NULL,
    price            bigint                DEFAULT 0                 NOT NULL,
    created_at       timestamp without time zone DEFAULT current_timestamp,
    updated_at       timestamp without time zone,
    deleted_at       timestamp without time zone,
    PRIMARY KEY (id),
    CONSTRAINT fk_customer_role
        FOREIGN KEY (customer_role_id)
            REFERENCES customer_roles (id),
    CONSTRAINT fk_customer
        FOREIGN KEY (product_id)
            REFERENCES products (id)
)