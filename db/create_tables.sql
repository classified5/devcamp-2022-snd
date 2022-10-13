CREATE TABLE IF NOT EXISTS products (
    id bigserial PRIMARY KEY,
    name varchar NOT NULL,
    price int NOT NULL,
    stock int NOT NULL,
    discount int NOT NULL,
    created_at timestamp NOT NULL,
    created_by int NOT NULL,
    updated_at timestamp NOT NULL,
    updated_by int NOT NULL
);

CREATE TABLE IF NOT EXISTS product_variants (
    id bigserial PRIMARY KEY,
    product_id bigserial REFERENCES products
    name varchar NOT NULL,
    price int NOT NULL,
    stock int NOT NULL,
    discount int NOT NULL,
    created_at timestamp NOT NULL,
    created_by int NOT NULL,
    updated_at timestamp NOT NULL,
    updated_by int NOT NULL
);