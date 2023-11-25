-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE VIEW view_plak_user_details AS
SELECT
    u.id AS user_id,
    u.first_name,
    u.last_name,
    u.email,
    COALESCE(jsonb_agg(
        jsonb_build_object(
            'id', r.id,
            'name', r.name,
            'description', r.description
        ) ORDER BY r.id
    ) FILTER (WHERE r.id IS NOT NULL), '[]') AS roles
FROM
    plak_users u
LEFT JOIN
    plak_user_roles ur ON u.id = ur.user_id
LEFT JOIN
    plak_roles r ON ur.role_id = r.id
GROUP BY
    u.id;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
