CREATE TABLE role_type (
    id serial PRIMARY KEY,
    "name" VARCHAR (35) NOT NULL UNIQUE,
    modules JSONB NOT NULL DEFAULT '{}'::JSONB,
    menus JSONB NOT NULL DEFAULT '{}'::JSONB,
    paths JSONB NOT NULL DEFAULT '{}'::JSONB,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
)