-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE VIEW view_plak_user_roles AS
SELECT 
    u.id AS user_id,  
    r.id AS role_id, 
    r.name AS role_name
FROM 
    plak_user_roles ur
INNER JOIN 
    plak_users u ON ur.user_id = u.id
INNER JOIN 
    plak_roles r ON ur.role_id = r.id;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
