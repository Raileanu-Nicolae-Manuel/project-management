-- name: GetProjectRole :one
SELECT * FROM project_roles
WHERE id = ? LIMIT 1;

-- name: GetProjectRoleByName :one
SELECT * FROM project_roles
WHERE name = ? LIMIT 1;

-- name: ListProjectRoles :many
SELECT * FROM project_roles
ORDER BY id; 