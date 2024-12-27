-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO users (
    username, email, password_hash
) VALUES (
    ?, ?, ?
);

-- name: UpdateUser :exec
UPDATE users
SET username = ?, email = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: CreateUserType :execresult
INSERT INTO user_types (name, description) VALUES (?, ?);

-- name: GetUserType :one
SELECT * FROM user_types WHERE id = ?;

-- name: ListUserTypes :many
SELECT * FROM user_types;

-- name: UpdateUserType :exec
UPDATE user_types SET name = ?, description = ? WHERE id = ?;

-- name: DeleteUserType :exec
DELETE FROM user_types WHERE id = ?;

-- name: AssignUserType :exec
INSERT INTO user_type_assignments (user_id, user_type_id) VALUES (?, ?);

-- name: RemoveUserType :exec
DELETE FROM user_type_assignments WHERE user_id = ? AND user_type_id = ?;

-- name: GetUserTypes :many
SELECT ut.* 
FROM user_types ut
JOIN user_type_assignments uta ON ut.id = uta.user_type_id
WHERE uta.user_id = ?; 