-- +goose Up
-- +goose StatementBegin


CREATE SEQUENCE plak_users_email_sequence
    AS BIGINT
    START WITH 10000000000
    INCREMENT BY 1;

CREATE TABLE plak_user_email (
    id BIGINT PRIMARY KEY DEFAULT nextval('plak_users_email_sequence'),
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


