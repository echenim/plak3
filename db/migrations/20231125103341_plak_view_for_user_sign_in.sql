-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE VIEW public.view_plak_user_signed_in_roles
 AS
 SELECT u.id AS user_id,
    concat(u.first_name, ' ', u.last_name) AS name,
    su.user_name AS username,
    su.password_hash AS passwordhash,
    su.email_confirmed,
    su.phone_number_confirmed,
    su.two_factor_enabled AS mfa_enabled,
    su.lockout_enabled AS lockedout,
    u.email,
   COALESCE(json_agg(r.name) FILTER (WHERE r.name IS NOT NULL), '[]') AS roles,
   su.security_stamp
   FROM plak_users u
     JOIN plak_user_sign_in su ON u.id = su.user_id
     LEFT JOIN plak_user_roles ur ON u.id = ur.user_id
     LEFT JOIN plak_roles r ON ur.role_id::text = r.id::text
  GROUP BY u.id, su.user_name, su.password_hash, su.email_confirmed, su.phone_number_confirmed, su.two_factor_enabled, su.lockout_enabled,su.security_stamp,u.email;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
