-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE VIEW view_plak_user_roles AS
SELECT 
    u.id AS user_id,  
    r.id AS role_id, 
    r.name AS role_name,
    r.description AS role_description
FROM 
    plak_user_roles ur
INNER JOIN 
    plak_users u ON ur.user_id = u.id
INNER JOIN 
    plak_roles r ON ur.role_id = r.id;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS view_plak_user_roles;
-- +goose StatementEnd
