-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE plak_user_sign_in (
    id TEXT PRIMARY KEY DEFAULT REPLACE(uuid_generate_v4()::text, '-', ''),
    user_id BIGINT NOT NULL UNIQUE,
    user_name VARCHAR(50),
    normalized_user_name VARCHAR(50),
    email VARCHAR(50),
    normalized_email VARCHAR(256),
    email_confirmed BOOLEAN NOT NULL,
    password_hash TEXT,
    security_stamp TEXT,
    concurrency_stamp TEXT,
    phone_number VARCHAR(20),
    phone_number_confirmed BOOLEAN NOT NULL,
    two_factor_enabled BOOLEAN NULL,
    lockout_end TIMESTAMPTZ NULL,
    lockout_enabled BOOLEAN NULL,
    access_failed_count INTEGER NULL,
    last_successful_sign_in_date TIMESTAMPTZ NULL,
    last_failed_sign_in_date TIMESTAMPTZ NULL,
    FOREIGN KEY (user_id) REFERENCES plak_users(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS plak_user_sign_in;
-- +goose StatementEnd
