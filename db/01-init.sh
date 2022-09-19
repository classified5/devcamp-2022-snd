#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $POSTGRES_DB;
  GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $APP_DB_USER;
  \connect $POSTGRES_DB $APP_DB_USER
  BEGIN;
    CREATE TABLE IF NOT EXISTS shipper (
	  id bigserial PRIMARY KEY,
      name varchar NOT NULL,
      image_url text NOT NULL,
      description text NOT NULL,
      max_weight int NOT NULL,
      created_at timestamp NOT NULL,
      created_by int NOT NULL,
      updated_at timestamp NOT NULL,
      updated_by int NOT NULL
	);
  COMMIT;
EOSQL