CREATE TABLE data_sources (
    id          INTEGER AUTO_INCREMENT PRIMARY KEY,
    `name`      VARCHAR (255) NOT NULL,
    url         VARCHAR (255) NOT NULL,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
