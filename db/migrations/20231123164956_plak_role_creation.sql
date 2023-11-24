-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE plak_roles (
    id VARCHAR(60) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT
);

INSERT INTO plak_roles (id, name, description) VALUES
(replace(gen_random_uuid()::text, '-', ''), 'Administrator', 'Administrative role with full access'),
(replace(gen_random_uuid()::text, '-', ''), 'Editor', 'Editor role with editing privileges'),
(replace(gen_random_uuid()::text, '-', ''), 'Create', 'Editor role with create privileges'),
(replace(gen_random_uuid()::text, '-', ''), 'Delete', 'Editor role with delete privileges'),
(replace(gen_random_uuid()::text, '-', ''), 'Read', 'Viewer role with read-only access');

CREATE TABLE plak_user_roles (
    user_id BIGINT NOT NULL,
    role_id VARCHAR(60) NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES plak_users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES plak_roles(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS plak_roles;

-- +goose StatementEnd
