CREATE TABLE modules (
    id serial PRIMARY KEY,
    "name" VARCHAR (60) NOT NULL UNIQUE,
    pattern VARCHAR (60) NOT NULL,
    "path" VARCHAR (60) NOT NULL,
    label VARCHAR (60) NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)