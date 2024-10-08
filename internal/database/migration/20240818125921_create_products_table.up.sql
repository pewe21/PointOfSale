CREATE TABLE products
(
    id            character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    brand_id      character varying(36)                           NOT NULL,
    supplier_id   character varying(36)                           NOT NULL,
    sku           character varying(10)                           NOT NULL,
    name          character varying(50)                           NOT NULL,
    stock         integer               DEFAULT 0                 NOT NULL,
    created_at    timestamp without time zone DEFAULT current_timestamp,
    updated_at    timestamp without time zone,
    deleted_at    timestamp without time zone,
    PRIMARY KEY (id),
    CONSTRAINT fk_type
        FOREIGN KEY (brand_id)
            REFERENCES brands (id),
    CONSTRAINT fk_supplier
        FOREIGN KEY (supplier_id)
            REFERENCES suppliers (id)
)