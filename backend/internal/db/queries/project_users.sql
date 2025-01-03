-- name: AddUserToProject :exec
INSERT INTO user_projects (
    user_id, project_id, role_id
) VALUES (
    ?, ?, ?
);

-- name: RemoveUserFromProject :exec
DELETE FROM user_projects
WHERE user_id = ? AND project_id = ?;

-- name: GetProjectUsers :many
SELECT u.*, pr.name as role_name
FROM users u
JOIN user_projects up ON u.id = up.user_id
JOIN project_roles pr ON up.role_id = pr.id
WHERE up.project_id = ?;

-- name: GetUserProjects :many
SELECT p.*, pr.name as role_name
FROM projects p
JOIN user_projects up ON p.id = up.project_id
JOIN project_roles pr ON up.role_id = pr.id
WHERE up.user_id = ?;

-- name: UpdateUserProjectRole :exec
UPDATE user_projects
SET role_id = ?
WHERE user_id = ? AND project_id = ?; 