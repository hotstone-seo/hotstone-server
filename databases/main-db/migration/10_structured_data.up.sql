CREATE TABLE structured_datas (
    id serial PRIMARY KEY,
    rule_id INTEGER NOT NULL,
    "type" VARCHAR (255) NOT NULL,
    "data" JSONB NOT NULL DEFAULT '{}'::JSONB,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX structured_datas_rule_id_idx ON structured_datas(rule_id);