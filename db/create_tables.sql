CREATE TABLE IF NOT EXISTS shipper (
    id bigserial PRIMARY KEY,
    name varchar NOT NULL,
    image_url text NOT NULL,
    description text NOT NULL,
    max_weight int NOT NULL,
    seller_id varchar NOT NULL,
    created_at timestamp NOT NULL,
    created_by int NOT NULL,
    updated_at timestamp NOT NULL,
    updated_by int NOT NULL,
    FOREIGN KEY (seller_id) REFERENCES seller(id)
);

CREATE TABLE IF NOT EXISTS seller (
    id bigserial PRIMARY KEY,
    name varchar NOT NULL,
    password varchar NOT NULL,
    created_at timestamp NOT NULL,
    created_by int NOT NULL,
    updated_at timestamp,
    updated_by int
);