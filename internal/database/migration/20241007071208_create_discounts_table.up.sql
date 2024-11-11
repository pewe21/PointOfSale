CREATE TABLE discounts
(
    id         character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    product_id character varying(36)                           NOT NULL,
    min_buy    integer               DEFAULT 1                 NOT NULL,
    disc_for   text[]                                            NOT NULL,
    price      bigint                DEFAULT 0                 NOT NULL,
    date_start date                                            NOT NULL,
    date_end   date                                            NOT NULL,
    created_at timestamp without time zone DEFAULT current_timestamp,
    updated_at timestamp without time zone,
    PRIMARY KEY (id),
    CONSTRAINT fk_customer
        FOREIGN KEY (product_id)
            REFERENCES products (id)
)