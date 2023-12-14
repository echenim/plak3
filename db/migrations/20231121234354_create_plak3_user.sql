-- +goose Up
-- +goose StatementBegin



CREATE TABLE plak_user_email (
    id TEXT PRIMARY KEY DEFAULT REPLACE(uuid_generate_v4()::text, '-', ''),
    email VARCHAR(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    lastUpdated TIMESTAMP WITH TIME ZONE
);


CREATE SEQUENCE plak_users_sequence
    AS BIGINT
    START WITH 100000000000
    INCREMENT BY 1;

CREATE TABLE plak_users (
    id BIGINT PRIMARY KEY DEFAULT nextval('plak_users_sequence'),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    lastUpdated TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS plak_users;

DROP SEQUENCE IF EXISTS plak_users_sequence;
-- +goose StatementEnd


