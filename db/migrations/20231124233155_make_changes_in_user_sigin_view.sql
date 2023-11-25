-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE VIEW view_plak_user_signed_in AS
SELECT
    u.id AS user_id,
    CONCAT(u.first_name, ' ', u.last_name) AS name,
	su.user_name as username,
	su.password_hash as passwordhash,
	su.email_confirmed,
	su.phone_number_confirmed,
	su.two_factor_enabled as mfa_enabled,
	su.lockout_enabled as lockedout,	
	
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
INNER JOIN
    plak_user_sign_in su ON u.id=su.user_id
LEFT JOIN
    plak_user_roles ur ON u.id = ur.user_id
LEFT JOIN
    plak_roles r ON ur.role_id = r.id
GROUP BY
    u.id,su.user_name,su.password_hash,su.email_confirmed,su.phone_number_confirmed,su.two_factor_enabled,su.lockout_enabled;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS view_plak_user_signed_in;
-- +goose StatementEnd
