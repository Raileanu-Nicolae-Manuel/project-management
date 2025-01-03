-- name: CreateProject :execresult
INSERT INTO projects (
    name, description
) VALUES (
    ?, ?
);

-- name: GetProject :one
SELECT * FROM projects
WHERE id = ? LIMIT 1;

-- name: ListProjects :many
SELECT * FROM projects
ORDER BY id;

-- name: UpdateProject :exec
UPDATE projects
SET name = ?, description = ?
WHERE id = ?;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = ?; 